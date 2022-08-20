package entityManager

import (
	"fmt"
)

type InheritManagerMethod interface {
	getParentManager() InheritManagerMethod
	getAllTables() []string
	recurAddInQueue(layer InheritEntity)
	recurAfterFind(layer InheritEntity)
	recurInheritAfterNew(goenId int, inheritType GoenInheritType, layer InheritEntity)
	generateGoenId() int
}

// 以下是Manager扩展的方法用于继承实体中的基实体
func (p *Manager[T, PT]) getParentManager() InheritManagerMethod {
	return nil
}
func (p *Manager[T, PT]) getAllTables() []string {
	return []string{p.tableName}
}
func (p *Manager[T, PT]) recurAddInQueue(layer InheritEntity) {
	p.addInQueue(layer.(PT))
}
func (p *Manager[T, PT]) recurAfterFind(layer InheritEntity) {
	layer.afterFind()
}
func (p *Manager[T, PT]) recurInheritAfterNew(goenId int, inheritType GoenInheritType, layer InheritEntity) {
	layer.inheritAfterNew(goenId, inheritType)
}

// 以下是InheritManager的实现
type inheritManagerTypeParam[T any] interface {
	*T
	InheritEntity
}

type InheritManager[T any, PT inheritManagerTypeParam[T]] struct {

	// 如果该类是基类，可以借用之前实现的管理器的方法
	*Manager[T, PT]
	// 如果该类不是基类，可以借用父辈管理器的方法
	parentManager InheritManagerMethod

	GoenInheritType
}

func (p *InheritManager[T, PT]) New() PT {
	e := PT(new(T))
	p.recurInheritAfterNew(p.generateGoenId(), p.GoenInheritType, e)
	//e.inheritAfterNew(p.generateGoenId(), p.GoenInheritType)
	p.recurAddInQueue(e)
	return e
}
func (p *InheritManager[T, PT]) generateGoenId() int {
	return p.getParentManager().generateGoenId()
}

func NewInheritManager[T any, PT inheritManagerTypeParam[T]](tableName string, parentManager InheritManagerMethod, inheritType GoenInheritType) *InheritManager[T, PT] {
	manager := &Manager[T, PT]{}
	manager.tableName = tableName
	return &InheritManager[T, PT]{
		GoenInheritType: inheritType,
		parentManager:   parentManager,
		Manager:         manager,
	}
}

func (p *InheritManager[T, PT]) recurAddInQueue(layer InheritEntity) {
	p.getParentManager().recurAddInQueue(layer.GetParentEntity())
	p.addInQueue(layer.(PT))
}
func (p *InheritManager[T, PT]) recurAfterFind(layer InheritEntity) {
	p.getParentManager().recurAfterFind(layer.GetParentEntity())
	layer.afterFind()
}
func (p *InheritManager[T, PT]) recurInheritAfterNew(goenId int, inheritType GoenInheritType, layer InheritEntity) {
	p.getParentManager().recurInheritAfterNew(goenId, inheritType, layer.GetParentEntity())
	layer.afterNew(goenId)
}

func (p *InheritManager[T, PT]) getAllTables() []string {
	return append(p.getParentManager().getAllTables(), p.tableName)
}

func (p *InheritManager[T, PT]) getParentManager() InheritManagerMethod {
	return p.parentManager
}

func (p *InheritManager[T, PT]) Get(goenId int) (PT, error) {
	tables := p.getAllTables()
	e := PT(new(T))
	//query := fmt.Sprintf("select * from %s  where goen_id=? and goen_in_all_instance = true", tables[0])
	for _, table := range tables {
		query := fmt.Sprintf("select * from %s  where goen_id=?", table)
		err := Db.Get(e, query, goenId)
		if err != nil {
			return nil, err
		}
	}

	//e.afterFind()
	p.recurAfterFind(e)
	p.recurAddInQueue(e)

	return e, nil
}

func (p *InheritManager[T, PT]) GetFromAllInstanceBy(field string, value any) (PT, error) {
	e := PT(new(T))
	query := fmt.Sprintf("select * from %s where %s=? and goen_in_all_instance = true %s", p.getTablesQuery(), field, p.getJoinQuery())
	err := Db.Get(e, query, value)
	if err != nil {
		return nil, err
	}
	//e.afterFind()
	p.recurAfterFind(e)
	p.recurAddInQueue(e)
	return e, nil
}

func (p *InheritManager[T, PT]) getTablesQuery() string {
	tables := p.getAllTables()
	tablesQuery := tables[0]
	for _, table := range tables[1:] {
		tablesQuery += fmt.Sprintf(", %s", table)
	}
	return tablesQuery
}

// form: table1.goen_id = table2.goen_id and table2.goen_id = table3.goen_id
func (p *InheritManager[T, PT]) getJoinQuery() string {
	tables := p.getAllTables()
	joinQuery := fmt.Sprintf("and %s.goen_id = ", tables[0])
	for _, table := range tables[1 : len(tables)-1] {
		joinQuery += fmt.Sprintf("%s.goen_id and %s.goen_id = ", table, table)
	}
	joinQuery += fmt.Sprintf("%s.goen_id", tables[len(tables)-1])
	return joinQuery
}

func (p *InheritManager[T, PT]) FindFromAllInstanceBy(field string, value any) ([]PT, error) {
	var entityArr []PT
	query := fmt.Sprintf("select * from %s where %s=? and goen_in_all_instance = true %s", p.getTablesQuery(), field, p.getJoinQuery())
	err := Db.Get(entityArr, query, value)
	if err != nil {
		return nil, err
	}
	for _, e := range entityArr {
		//e.afterFind()
		p.recurAfterFind(e)
		p.recurAddInQueue(e)
	}
	return entityArr, nil
}

func (p *InheritManager[T, PT]) FindFromMultiAssTable(assTableName string, ownerId int) ([]PT, error) {
	var entityArr []PT
	tables := p.getAllTables()
	query := fmt.Sprintf("select tmp.* from %s as ass, %s where ass.owner_goen_id = ? and ass.possession_goen_id = %s.goen_id %s ",
		assTableName, p.getTablesQuery(), tables[0], p.getJoinQuery())
	if err := Db.Select(&entityArr, query, ownerId); err != nil {
		return nil, err
	}
	for _, e := range entityArr {
		//e.afterFind()
		p.recurAfterFind(e)
		p.addInQueue(e)
	}
	return entityArr, nil
}
