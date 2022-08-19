package entityManager

//import "Cocome/operation"
//
//type StoreManager struct {
//}
//
//type Store struct {
//	Entity
//	Id       int     `db:"id"`
//	Name     *string `db:"name"`
//	Address  *string `db:"address"`
//	IsOpened *bool   `db:"is_opened"`
//}
//
//var basicCandidate [4]string
//var setCandidateLen int
//
//func addBasicField(str string) {
//	basicCandidate[setCandidateLen] = str
//	setCandidateLen++
//}
//func clearSetField() {
//	setCandidateLen = 0
//}
//
//const (
//	Id       = "id"
//	Name     = "name"
//	Address  = "address"
//	IsOpened = "is_opened"
//)
//
//func getUpdateQuery() string {
//	query := "insert into store("
//	for i := 0; i < setCandidateLen-1; i++ {
//		query += basicCandidate[i] + ","
//	}
//	query += basicCandidate[setCandidateLen-1] + ") values(:"
//	for i := 0; i < setCandidateLen-1; i++ {
//		query += basicCandidate[i] + ", :"
//	}
//	query += basicCandidate[setCandidateLen-1] + ")"
//	print(query)
//	return query
//}
//func getInsertQuery() string {
//	query := "update store set "
//	for i := 0; i < setCandidateLen-1; i++ {
//		query += basicCandidate[i] + "= :" + basicCandidate[i] + ", "
//	}
//	query += basicCandidate[setCandidateLen-1] + "= :" + basicCandidate[setCandidateLen-1]
//	query += " where " + " id = :id"
//	print(query)
//	return query
//}
//
//func (p *Store) SetName(str string) {
//	p.Name = &str
//	addBasicField(Name)
//}
//func (p *Store) SetAddress(str string) {
//	p.Address = &str
//	addBasicField(Address)
//}
//func (p *Store) SetIsOpened(bool bool) {
//	p.IsOpened = &bool
//	addBasicField(IsOpened)
//}
//func (p *Store) Update() error {
//	_, err := operation.Db.NamedExec(getUpdateQuery(), p)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func (p *Store) Insert() error {
//	_, err := operation.Db.NamedExec(getInsertQuery(), p)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func (p *Store) Save() error {
//	if p.entityType == Founded {
//		return p.Update()
//	} else {
//		return p.Insert()
//	}
//}
//
//func FindStore(id int) (*Store, error) {
//	store := &Store{}
//	err := operation.Db.Get(store, "select * from store where id = $1", id)
//	if err != nil {
//		return nil, err
//	}
//	store.entityType = Founded
//	return store, nil
//}
//
//func CreateStore() (store *Store) {
//	store = new(Store)
//	store.entityType = Created
//	return
//}
