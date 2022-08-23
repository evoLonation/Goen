package entityManager

type entityForCast interface {
	GetRealType() GoenInheritType
}

func (p *manager[T, PT]) CastFrom(e entityForCast) (PT, error) {
	//e.GetRealType()
	ei := (any(e)).(EntityForManager)
	return p.Get(ei.GetGoenId())
}
