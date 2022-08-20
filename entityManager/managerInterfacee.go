package entityManager

type ManagerForOther[T entityForManager] interface {
	New() T
	GetBy(member string, value any) (T, error)
	FindBy(member string, value any) ([]T, error)
	AddInAllInstance(e T)
	RemoveFromAllInstance(e T)
}

type ManagerForEntity[PT entityForManager] interface {
	Get(goenId int) (PT, error)
	FindFromMultiAssTable(tableName string, ownerId int) ([]PT, error)
}
