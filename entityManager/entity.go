package entityManager

type EntityStatus int

const (
	Created EntityStatus = iota
	Existent
)

type MultiAssChangeType int

const (
	Include MultiAssChangeType = iota
	Exclude
)

type multiAssInfo struct {
	typ       MultiAssChangeType
	tableName string
	targetId  int
}

type Entity struct {
	FieldChange

	GoenId            int  `db:"goen_id"`
	GoenInAllInstance bool `db:"goen_in_all_instance"`
}

type FieldChange struct {
	EntityStatus
	basicFieldChange []string
	assFieldChange   []string
	multiAssChange   []multiAssInfo
}

func (p *FieldChange) getEntityStatus() EntityStatus {
	return p.EntityStatus
}

func (p *FieldChange) getBasicFieldChange() []string {
	return p.basicFieldChange
}

func (p *FieldChange) getAssFieldChange() []string {
	return p.assFieldChange
}

func (p *FieldChange) getMultiAssChange() []multiAssInfo {
	return p.multiAssChange
}

func (p *FieldChange) afterNew(goenId int) {
	p.EntityStatus = Created
}
func (p *FieldChange) afterFind() {
	p.EntityStatus = Existent
}
func (p *FieldChange) afterBasicSave() {
	p.EntityStatus = Existent
	p.basicFieldChange = nil
}
func (p *FieldChange) afterAssUpdate() {
	p.EntityStatus = Existent
	p.assFieldChange = nil
	p.multiAssChange = nil
}

//剩下的after都继承了FieldChange
func (p *Entity) afterNew(goenId int) {
	p.EntityStatus = Created
	p.GoenId = goenId
}

func (p *Entity) getGoenId() int {
	return p.GoenId
}

func (p *Entity) setGoenInAllInstance(goenInAllInstance bool) {
	p.GoenInAllInstance = goenInAllInstance
	p.AddBasicFieldChange("goen_in_all_instance")
}

func (p *FieldChange) AddBasicFieldChange(field string) {
	p.basicFieldChange = append(p.basicFieldChange, field)
}

func (p *FieldChange) AddAssFieldChange(field string) {
	p.assFieldChange = append(p.assFieldChange, field)
}

func (p *FieldChange) AddMultiAssChange(typ MultiAssChangeType, tableName string, targetId int) {
	p.multiAssChange = append(p.multiAssChange, multiAssInfo{
		typ:       typ,
		targetId:  targetId,
		tableName: tableName,
	})
}
