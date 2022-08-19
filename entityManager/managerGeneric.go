package entityManager

import (
	"errors"
	"fmt"
)

// entity methods for managerGeneric
type entityMethods interface {
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

type entityInterface[T any] interface {
	*T
	GetEntityMethods() entityMethods
}

type ManagerGeneric[T any, PT entityInterface[T]] struct {
	tableName string
	maxGoenId int
}

func NewManager[T any, PT entityInterface[T]](tableName string) *ManagerGeneric[T, PT] {
	manager := &ManagerGeneric[T, PT]{}
	manager.tableName = tableName
	query := fmt.Sprintf("select goen_id from %s order by goen_id DESC limit 1", manager.tableName)
	Db.Get(&manager.maxGoenId, query)
	return manager
}

func (p *ManagerGeneric[T, PT]) generateGoenId() int {
	p.maxGoenId = p.maxGoenId + 1
	return p.maxGoenId
}
func (p *ManagerGeneric[T, PT]) addInQueue(e PT) {
	Manager.addInBasicSaveQueue(func() error {
		return p.saveBasic(e)
	})
	Manager.addInAssSaveQueue(func() error {
		return p.updateAss(e)
	})
}

func (p *ManagerGeneric[T, PT]) New() PT {
	e := PT(new(T))
	e.GetEntityMethods().afterNew(p.generateGoenId())
	p.addInQueue(e)
	return e
}

func (p *ManagerGeneric[T, PT]) Get(goenId int) (PT, error) {
	e := PT(new(T))
	query := fmt.Sprintf("select * from %s where goen_id=? and goen_in_all_instance = true", p.tableName)
	err := Db.Get(e, query, goenId)
	if err != nil {
		return nil, err
	}
	e.GetEntityMethods().afterFind()
	p.addInQueue(e)
	return e, nil
}

func (p *ManagerGeneric[T, PT]) GetBy(member string, value any) (PT, error) {
	e := PT(new(T))
	query := fmt.Sprintf("select * from %s where %s=? and goen_in_all_instance = true", p.tableName, member)
	err := Db.Get(e, query, value)
	if err != nil {
		return nil, err
	}
	e.GetEntityMethods().afterFind()
	p.addInQueue(e)
	return e, nil
}

func (p *ManagerGeneric[T, PT]) FindBy(member string, value any) ([]PT, error) {
	var entityArr []PT
	query := fmt.Sprintf("select * from %s where %s=? and goen_in_all_instance = true", p.tableName, member)
	err := Db.Select(&entityArr, query, value)
	if err != nil {
		return nil, err
	}
	for _, e := range entityArr {
		e.GetEntityMethods().afterFind()
		p.addInQueue(e)
	}
	return entityArr, nil
}
func (p *ManagerGeneric[T, PT]) FindFromMultiAssTable(tableName string, ownerId int) ([]PT, error) {
	var entityArr []PT
	query := fmt.Sprintf("select tmp.* from %s as ass, %s as tmp where ass.owner_goen_id = ? and ass.possession_goen_id = tmp.goen_id ",
		tableName, p.tableName)
	if err := Db.Select(&entityArr, query, ownerId); err != nil {
		return nil, err
	}
	for _, e := range entityArr {
		e.GetEntityMethods().afterFind()
		p.addInQueue(e)
	}
	return entityArr, nil
}

// the length of changedField must > 0
func (p *ManagerGeneric[T, PT]) getUpdateQuery(changedField []string) string {
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
func (p *ManagerGeneric[T, PT]) getInsertQuery(changedField []string) string {
	lastField := changedField[len(changedField)-1]
	query := fmt.Sprintf("insert into %s( ", p.tableName)
	for _, field := range changedField[0 : len(changedField)-1] {
		query += fmt.Sprintf("%s, ", field)
	}
	query += fmt.Sprintf("%s) values( ", lastField)
	for _, field := range changedField[0 : len(changedField)-1] {
		query += fmt.Sprintf(":%s ,", field)
	}
	query += fmt.Sprintf(":%s )", lastField)
	print(query)
	return query
}

func (p *ManagerGeneric[T, PT]) getMultiAssInsertQuery(tableName string) string {
	query := fmt.Sprintf("insert into %s (owner_goen_id, possession_goen_id) values (?, ?)", tableName)
	return query
}
func (p *ManagerGeneric[T, PT]) getMultiAssDeleteQuery(tableName string) string {
	query := fmt.Sprintf("delete from %s where owner_goen_id=? and possession_goen_id=?", tableName)
	return query
}

func (p *ManagerGeneric[T, PT]) saveBasic(e PT) error {
	methods := e.GetEntityMethods()
	if len(methods.getBasicFieldChange()) != 0 {
		if methods.getEntityStatus() == Created {
			if _, err := Db.NamedExec(p.getInsertQuery(methods.getBasicFieldChange()), e); err != nil {
				return err
			}
		} else {
			if _, err := Db.NamedExec(p.getUpdateQuery(methods.getBasicFieldChange()), e); err != nil {
				return err
			}
		}
	}
	methods.afterBasicSave()
	return nil
}

// updateAss e 's entityStatus must be Existence
func (p *ManagerGeneric[T, PT]) updateAss(e PT) error {
	methods := e.GetEntityMethods()

	if methods.getEntityStatus() == Created {
		return errors.New("entityStatus must be Existence")
	}
	if len(methods.getAssFieldChange()) != 0 {
		if _, err := Db.NamedExec(p.getUpdateQuery(methods.getAssFieldChange()), e); err != nil {
			return err
		}
	}
	for _, info := range methods.getMultiAssChange() {
		var query string
		if info.typ == Include {
			query = p.getMultiAssInsertQuery(info.tableName)
		} else {
			query = p.getMultiAssDeleteQuery(info.tableName)
		}
		if _, err := Db.Exec(query, methods.getGoenId(), info.targetId); err != nil {
			return err
		}
	}
	methods.afterAssUpdate()
	return nil
}

func (p *ManagerGeneric[T, PT]) AddInAllInstance(e PT) {
	e.GetEntityMethods().setGoenInAllInstance(true)
}

func (p *ManagerGeneric[T, PT]) RemoveFromAllInstance(e PT) {
	e.GetEntityMethods().setGoenInAllInstance(false)
}
