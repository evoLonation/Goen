package entityManager

import "Cocome/operation"

type EntityType int

const (
	Created EntityType = 0
	Founded EntityType = 1
)

type Store struct {
	entityType EntityType
	Id         int     `db:"id"`
	Name       *string `db:"name"`
	Address    *string `db:"address"`
	IsOpened   *bool   `db:"is_opened"`
}

const (
	Id       = "id"
	Name     = "name"
	Address  = "address"
	IsOpened = "is_opened"
)

var setCandidate [4]string
var setCandidateLen int

func addSetField(str string) {
	setCandidate[setCandidateLen] = str
	setCandidateLen++
}
func clearSetField() {
	setCandidateLen = 0
}

func getUpdateQuery() string {
	query := "insert into store("
	for i := 0; i < setCandidateLen-1; i++ {
		query += setCandidate[i] + ","
	}
	query += setCandidate[setCandidateLen-1] + ") values(:"
	for i := 0; i < setCandidateLen-1; i++ {
		query += setCandidate[i] + ", :"
	}
	query += setCandidate[setCandidateLen-1] + ")"
	print(query)
	return query
}
func getInsertQuery() string {
	query := "update store set "
	for i := 0; i < setCandidateLen-1; i++ {
		query += setCandidate[i] + "= :" + setCandidate[i] + ", "
	}
	query += setCandidate[setCandidateLen-1] + "= :" + setCandidate[setCandidateLen-1]
	query += " where " + " id = :id"
	print(query)
	return query
}

func (p *Store) SetName(str string) {
	p.Name = &str
	addSetField(Name)
}
func (p *Store) SetAddress(str string) {
	p.Address = &str
	addSetField(Address)
}
func (p *Store) SetIsOpened(bool bool) {
	p.IsOpened = &bool
	addSetField(IsOpened)
}
func (p *Store) Update() error {
	_, err := operation.Db.NamedExec(getUpdateQuery(), p)
	if err != nil {
		return err
	}
	return nil
}
func (p *Store) Insert() error {
	_, err := operation.Db.NamedExec(getInsertQuery(), p)
	if err != nil {
		return err
	}
	return nil
}
func (p *Store) Save() error {
	if p.entityType == Founded {
		return p.Update()
	} else {
		return p.Insert()
	}
}

func FindStore(id int) (*Store, error) {
	store := &Store{}
	err := operation.Db.Get(store, "select * from store where id = $1", id)
	if err != nil {
		return nil, err
	}
	store.entityType = Founded
	return store, nil
}

func CreateStore() (store *Store) {
	store = new(Store)
	store.entityType = Created
	return
}