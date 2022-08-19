package entityManager

import (
	"fmt"
)

type entityType interface {
	Entity | Item
}

type entityInterface[T any] interface {
	*T
	initEntity(EntityType, string, int)
}

type ManagerGeneric[T entityType, PT entityInterface[T]] struct {
	tableName string
	maxGoenId int
}

func NewManager[T entityType, PT entityInterface[T]](tableName string) (*ManagerGeneric[T, PT], error) {
	manager := &ManagerGeneric[T, PT]{}
	manager.tableName = tableName
	query := fmt.Sprintf("select goen_id from %s order by goen_id DESC limit 1", manager.tableName)
	Db.Get(&manager.maxGoenId, query)
	return manager, nil
}

func (p *ManagerGeneric[T, PT]) generateGoenId() int {
	p.maxGoenId = p.maxGoenId + 1
	return p.maxGoenId
}

func (p *ManagerGeneric[T, PT]) Find(goenId int) (PT, error) {
	e := PT(new(T))
	query := fmt.Sprintf("select * from %s where goen_id=?", p.tableName)
	err := Db.Get(e, query, goenId)
	if err != nil {
		return nil, err
	}
	e.initEntity(Founded, p.tableName, 0)

	return e, nil
}

func (p *ManagerGeneric[T, PT]) GetBy(member string, value any) (PT, error) {
	e := PT(new(T))
	query := fmt.Sprintf("select * from %s where %s=?", p.tableName, member)
	err := Db.Get(e, query, value)
	if err != nil {
		return nil, err
	}
	e.initEntity(Founded, p.tableName, 0)

	return e, nil
}

func (p *ManagerGeneric[T, PT]) FindBy(member string, value any) ([]PT, error) {
	var earr []PT
	query := fmt.Sprintf("select * from %s where %s=?", p.tableName, member)
	err := Db.Select(&earr, query, value)
	if err != nil {
		return nil, err
	}
	for _, e := range earr {
		e.initEntity(Founded, p.tableName, 0)
	}
	return earr, nil
}

func (p *ManagerGeneric[T, PT]) New() PT {
	e := PT(new(T))
	e.initEntity(Created, p.tableName, p.generateGoenId())
	return e
}

//func (p *ManagerGeneric[T, PT]) Update(e PT) error {
//	_, err := operation.Db.NamedExec(e.getUpdateQuery(), e)
//	if err != nil {
//		return err
//}
//e.clearSetField()
//return nil
//}
//func (p *ManagerGeneric[T, PT]) Insert(e PT) error {
//	_, err := operation.Db.NamedExec(e.getInsertQuery(), e)
//	if err != nil {
//		return err
//	}
//	e.clearSetField()
//	e.setEntityType(Founded)
//	return nil
//}
