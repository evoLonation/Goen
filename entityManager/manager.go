package entityManager

import (
	"Cocome/operation"
	"fmt"
)

type EntityType int

const (
	Created EntityType = 0
	Founded EntityType = 1
)

type entity struct {
	entityType      EntityType
	setCandidate    [10]string
	setCandidateLen int
	getTableName    func() string
	getIdName       func() string
}

func (p *entity) SetEntityType(entityType EntityType) {
	p.entityType = entityType
}
func (p *entity) GetEntityType() EntityType {
	return p.entityType
}
func (p *entity) setTableNameFunc(foo func() string) {
	p.getTableName = foo
}
func (p *entity) setIdNameFunc(foo func() string) {
	p.getIdName = foo
}

type entityType interface {
	entity | Store | Item
}

type Manager[T entityType, PT interface {
	*T
	SetEntityType(entityType EntityType)
	GetEntityType() EntityType
	setTableNameFunc(func() string)
	setIdNameFunc(func() string)
	getInsertQuery() string
	getUpdateQuery() string
	clearSetField()
}] struct {
	tableName string
	idName    string
}

func (p *entity) addSetField(str string) {
	p.setCandidate[p.setCandidateLen] = str
	p.setCandidateLen++
}
func (p *entity) clearSetField() {
	p.setCandidateLen = 0
}

func (p *entity) getInsertQuery() string {
	query := fmt.Sprintf("insert into %s(", p.getTableName())
	for i := 0; i < p.setCandidateLen-1; i++ {
		query += fmt.Sprintf("%s,", p.setCandidate[i])
	}
	query += fmt.Sprintf("%s) values(", p.setCandidate[p.setCandidateLen-1])
	for i := 0; i < p.setCandidateLen-1; i++ {
		query += fmt.Sprintf(":%s ,", p.setCandidate[i])
	}
	query += fmt.Sprintf(":%s )", p.setCandidate[p.setCandidateLen-1])
	print(query)
	return query
}

func (p *entity) getUpdateQuery() string {
	query := fmt.Sprintf("update %s set ", p.getTableName())
	for i := 0; i < p.setCandidateLen-1; i++ {
		query += p.setCandidate[i] + "= :" + p.setCandidate[i] + ", "
	}
	query += p.setCandidate[p.setCandidateLen-1] + "= :" + p.setCandidate[p.setCandidateLen-1]
	query += fmt.Sprintf(" where %s = :%s", p.getIdName(), p.getIdName())
	print(query)
	return query
}

func (p *Manager[T, PT]) Find(id any) (PT, error) {
	e := PT(new(T))
	query := fmt.Sprintf("select * from %s where %s=?", p.tableName, p.idName)
	err := operation.Db.Get(e, query, id)
	if err != nil {
		return nil, err
	}
	e.SetEntityType(Founded)
	e.setIdNameFunc(func() string { return p.idName })
	e.setTableNameFunc(func() string { return p.tableName })
	return e, nil
}

func (p *Manager[T, PT]) Create() PT {
	e := PT(new(T))
	e.SetEntityType(Created)
	e.setIdNameFunc(func() string { return p.idName })
	e.setTableNameFunc(func() string { return p.tableName })
	return e
}

func (p *Manager[T, PT]) Update(e PT) error {
	_, err := operation.Db.NamedExec(e.getUpdateQuery(), e)
	if err != nil {
		return err
	}
	e.clearSetField()
	return nil
}
func (p *Manager[T, PT]) Insert(e PT) error {
	_, err := operation.Db.NamedExec(e.getInsertQuery(), e)
	if err != nil {
		return err
	}
	e.clearSetField()
	return nil
}
func (p *Manager[T, PT]) Save(e PT) error {
	if e.GetEntityType() == Founded {
		return p.Update(e)
	} else {
		return p.Insert(e)
	}
}
