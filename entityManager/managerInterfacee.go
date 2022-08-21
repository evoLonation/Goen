package entityManager

type ManagerForOther[T entityInterfaceForManager] interface {
	New() T
	GetFromAllInstanceBy(member string, value any) (T, error)
	FindFromAllInstanceBy(member string, value any) ([]T, error)
	AddInAllInstance(e T)
	RemoveFromAllInstance(e T)
}

type ManagerForEntity[PT entityInterfaceForManager] interface {
	// Get 实际上不需要检查是否在allinstance里面
	Get(goenId int) (PT, error)
	FindFromMultiAssTable(tableName string, ownerId int) ([]PT, error)
}
