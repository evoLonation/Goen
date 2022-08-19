package entityManager

type EntityManager struct {
	waitBasicSave       []func() error
	waitAssociationSave []func() error
}

var Manager EntityManager

func (p *EntityManager) Save() error {
	for _, foo := range p.waitBasicSave {
		if err := foo(); err != nil {
			return err
		}
	}
	p.waitBasicSave = nil
	for _, foo := range p.waitAssociationSave {
		if err := foo(); err != nil {
			return err
		}
	}
	p.waitAssociationSave = nil
	return nil
}
func (p *EntityManager) addInBasicSaveQueue(foo func() error) {
	p.waitBasicSave = append(p.waitBasicSave, foo)
}
func (p *EntityManager) addInAssSaveQueue(foo func() error) {
	p.waitAssociationSave = append(p.waitAssociationSave, foo)
}
