package entityManager

type GoenInheritType int

type InheritEntity interface {
	GetParentEntity() InheritEntity
	inheritAfterNew(goenId int, inheritType GoenInheritType)
	entityForManager
}

type BasicEntity struct {
	Entity
	GoenInheritType GoenInheritType `db:"goen_inherit_type"`
}

func (p *BasicEntity) inheritAfterNew(goenId int, inheritType GoenInheritType) {
	p.afterNew(goenId)
	p.GoenInheritType = inheritType
	p.AddBasicFieldChange("goen_inherit_type")
}

func (p *BasicEntity) GetParentEntity() InheritEntity {
	return nil
}

func (p *BasicEntity) GetRealType() GoenInheritType {
	return p.GoenInheritType
}
