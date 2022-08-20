package entityManager

import (
	"errors"
	"fmt"
)

type entityForManager interface {
	afterNew(int)
	afterFind()
	afterBasicSave()
	afterAssUpdate()
	getEntityStatus() EntityStatus
	setGoenInAllInstance(bool)
	getBasicFieldChange() []string
	getAssFieldChange() []string
	getMultiAssChange() []multiAssInfo
	getGoenId() int
}

type managerTypeParam[T any] interface {
	*T
	entityForManager
}

type Manager[T any, PT managerTypeParam[T]] struct {
	tableName string
	maxGoenId int
}

func NewManager[T any, PT managerTypeParam[T]](tableName string) *Manager[T, PT] {
	manager := &Manager[T, PT]{}
	manager.tableName = tableName
	query := fmt.Sprintf("select goen_id from %s order by goen_id DESC limit 1", manager.tableName)
	Db.Get(&manager.maxGoenId, query)
	return manager
}

func (p *Manager[T, PT]) generateGoenId() int {
	p.maxGoenId = p.maxGoenId + 1
	return p.maxGoenId
}
func (p *Manager[T, PT]) addInQueue(e PT) {
	Saver.addInBasicSaveQueue(func() error {
		return p.saveBasic(e)
	})
	Saver.addInAssSaveQueue(func() error {
		return p.updateAss(e)
	})
}

func (p *Manager[T, PT]) New() PT {
	e := PT(new(T))
	e.afterNew(p.generateGoenId())
	p.addInQueue(e)
	return e
}

func (p *Manager[T, PT]) Get(goenId int) (PT, error) {
	e := PT(new(T))
	//query := fmt.Sprintf("select * from %s where goen_id=? and goen_in_all_instance = true", p.tableName)
	query := fmt.Sprintf("select * from %s where goen_id=?", p.tableName)
	err := Db.Get(e, query, goenId)
	if err != nil {
		return nil, err
	}
	e.afterFind()
	p.addInQueue(e)
	return e, nil
}

func (p *Manager[T, PT]) GetFromAllInstanceBy(member string, value any) (PT, error) {
	e := PT(new(T))
	query := fmt.Sprintf("select * from %s where %s=? and goen_in_all_instance = true", p.tableName, member)
	err := Db.Get(e, query, value)
	if err != nil {
		return nil, err
	}
	e.afterFind()
	p.addInQueue(e)
	return e, nil
}

func (p *Manager[T, PT]) FindFromAllInstanceBy(member string, value any) ([]PT, error) {
	var entityArr []PT
	query := fmt.Sprintf("select * from %s where %s=? and goen_in_all_instance = true", p.tableName, member)
	err := Db.Select(&entityArr, query, value)
	if err != nil {
		return nil, err
	}
	for _, e := range entityArr {
		e.afterFind()
		p.addInQueue(e)
	}
	return entityArr, nil
}
func (p *Manager[T, PT]) FindFromMultiAssTable(tableName string, ownerId int) ([]PT, error) {
	var entityArr []PT
	query := fmt.Sprintf("select tmp.* from %s as ass, %s as tmp where ass.owner_goen_id = ? and ass.possession_goen_id = tmp.goen_id ",
		tableName, p.tableName)
	if err := Db.Select(&entityArr, query, ownerId); err != nil {
		return nil, err
	}
	for _, e := range entityArr {
		e.afterFind()
		p.addInQueue(e)
	}
	return entityArr, nil
}

// the length of changedField must > 0
func (p *Manager[T, PT]) getUpdateQuery(changedField []string) string {
	query := fmt.Sprintf("update %s set ", p.tableName)
	for _, field := range changedField[0 : len(changedField)-1] {
		query += fmt.Sprintf("%s= :%s,", field, field)
	}
	field := changedField[len(changedField)-1]
	query += fmt.Sprintf("%s = :%s", field, field)
	query += fmt.Sprintf(" where goen_id = :goen_id")
	print(query)
	return query
}

// the length of changedField must > 0
func (p *Manager[T, PT]) getInsertQuery(changedField []string) string {
	lastField := changedField[len(changedField)-1]
	query := fmt.Sprintf("insert into %s(goen_id, ", p.tableName)
	for _, field := range changedField[0 : len(changedField)-1] {
		query += fmt.Sprintf("%s, ", field)
	}
	query += fmt.Sprintf("%s) values(:goen_id, ", lastField)
	for _, field := range changedField[0 : len(changedField)-1] {
		query += fmt.Sprintf(":%s ,", field)
	}
	query += fmt.Sprintf(":%s )", lastField)
	print(query)
	return query
}

func (p *Manager[T, PT]) getMultiAssInsertQuery(tableName string) string {
	query := fmt.Sprintf("insert into %s (owner_goen_id, possession_goen_id) values (?, ?)", tableName)
	return query
}
func (p *Manager[T, PT]) getMultiAssDeleteQuery(tableName string) string {
	query := fmt.Sprintf("delete from %s where owner_goen_id=? and possession_goen_id=?", tableName)
	return query
}

func (p *Manager[T, PT]) saveBasic(e PT) error {
	if len(e.getBasicFieldChange()) != 0 {
		if e.getEntityStatus() == Created {
			if _, err := Db.NamedExec(p.getInsertQuery(e.getBasicFieldChange()), e); err != nil {
				return err
			}
		} else {
			if _, err := Db.NamedExec(p.getUpdateQuery(e.getBasicFieldChange()), e); err != nil {
				return err
			}
		}
	}
	e.afterBasicSave()
	return nil
}

// updateAss e 's entityStatus must be Existence
func (p *Manager[T, PT]) updateAss(e PT) error {

	if e.getEntityStatus() == Created {
		return errors.New("entityStatus must be Existence")
	}
	if len(e.getAssFieldChange()) != 0 {
		if _, err := Db.NamedExec(p.getUpdateQuery(e.getAssFieldChange()), e); err != nil {
			return err
		}
	}
	for _, info := range e.getMultiAssChange() {
		var query string
		if info.typ == Include {
			query = p.getMultiAssInsertQuery(info.tableName)
		} else {
			query = p.getMultiAssDeleteQuery(info.tableName)
		}
		if _, err := Db.Exec(query, e.getGoenId(), info.targetId); err != nil {
			return err
		}
	}
	e.afterAssUpdate()
	return nil
}

func (p *Manager[T, PT]) AddInAllInstance(e PT) {
	e.setGoenInAllInstance(true)
}

func (p *Manager[T, PT]) RemoveFromAllInstance(e PT) {
	e.setGoenInAllInstance(false)
}
