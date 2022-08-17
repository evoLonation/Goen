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
	entityType EntityType
	Manager    *any
}

func (p *entity) SetEntityType(entityType EntityType) {
	p.entityType = entityType
}
func (p *entity) GetEntityType() EntityType {
	return p.entityType
}

type entityType interface {
	entity | Store | Item
}

type Manager[T entityType, PT interface {
	*T
	SetEntityType(entityType EntityType)
	GetEntityType() EntityType
}] struct {
	currentEntity   PT
	setCandidate    [10]string
	setCandidateLen int
	tableName       string
	idName          string
}

func (p *Manager[T, PT]) addSetField(str string) {
	p.setCandidate[p.setCandidateLen] = str
	p.setCandidateLen++
}
func (p *Manager[T, PT]) clearSetField() {
	p.setCandidateLen = 0
}

func (p *Manager[T, PT]) getUpdateQuery() string {
	query := fmt.Sprintf("insert into %s(", p.tableName)
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

func (p *Manager[T, PT]) getInsertQuery() string {
	query := fmt.Sprintf("update %s p.set ", p.tableName)
	for i := 0; i < p.setCandidateLen-1; i++ {
		query += p.setCandidate[i] + "= :" + p.setCandidate[i] + ", "
	}
	query += p.setCandidate[p.setCandidateLen-1] + "= :" + p.setCandidate[p.setCandidateLen-1]
	query += fmt.Sprintf(" where %s = :%s", p.idName, p.idName)
	print(query)
	return query
}

func (p *Manager[T, PT]) Find(id any) (PT, error) {
	e := PT(new(T))
	err := operation.Db.Get(e, fmt.Sprintf("select * from %s where %s = $1", p.tableName, p.idName), id)
	if err != nil {
		return nil, err
	}
	e.SetEntityType(Founded)
	p.currentEntity = e
	return e, nil
}

func (p *Manager[T, PT]) Create() PT {
	e := PT(new(T))
	e.SetEntityType(Created)
	p.currentEntity = e
	return e
}

func (p *Manager[T, PT]) Update() error {
	_, err := operation.Db.NamedExec(p.getUpdateQuery(), p.currentEntity)
	if err != nil {
		return err
	}
	return nil
}
func (p *Manager[T, PT]) Insert() error {
	_, err := operation.Db.NamedExec(p.getInsertQuery(), p.currentEntity)
	if err != nil {
		return err
	}
	return nil
}
func (p *Manager[T, PT]) Save() error {
	if p.currentEntity.GetEntityType() == Founded {
		return p.Update()
	} else {
		return p.Insert()
	}
}

func Test() {
	man := &Manager[Store, *Store]{}
	man.tableName = "1"
	//foo[entity]()
}
