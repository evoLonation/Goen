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
	EntityStatus

	basicFieldChange []string
	assFieldChange   []string
	multiAssChange   []multiAssInfo

	GoenId            int  `db:"goen_id"`
	GoenInAllInstance bool `db:"goen_in_all_instance"`
}

func (p *Entity) getEntityStatus() EntityStatus {
	return p.EntityStatus
}

func (p *Entity) getBasicFieldChange() []string {
	return p.basicFieldChange
}

func (p *Entity) getAssFieldChange() []string {
	return p.assFieldChange
}

func (p *Entity) getMultiAssChange() []multiAssInfo {
	return p.multiAssChange
}

func (p *Entity) getGoenId() int {
	return p.GoenId
}

func (p *Entity) afterNew(goenId int) {
	p.EntityStatus = Created
	p.GoenId = goenId
	p.AddBasicFieldChange("goen_id")
}
func (p *Entity) afterFind() {
	p.EntityStatus = Existent
}
func (p *Entity) afterBasicSave() {
	p.EntityStatus = Existent
	p.basicFieldChange = nil
}
func (p *Entity) afterAssUpdate() {
	p.EntityStatus = Existent
	p.assFieldChange = nil
	p.multiAssChange = nil
}

func (p *Entity) setGoenInAllInstance(goenInAllInstance bool) {
	p.GoenInAllInstance = goenInAllInstance
	p.AddBasicFieldChange("goen_in_all_instance")
}

func (p *Entity) AddBasicFieldChange(field string) {
	p.basicFieldChange = append(p.basicFieldChange, field)
}

func (p *Entity) AddAssFieldChange(field string) {
	p.assFieldChange = append(p.assFieldChange, field)
}

func (p *Entity) AddMultiAssChange(typ MultiAssChangeType, tableName string, targetId int) {
	p.multiAssChange = append(p.multiAssChange, multiAssInfo{
		typ:       typ,
		targetId:  targetId,
		tableName: tableName,
	})
}
