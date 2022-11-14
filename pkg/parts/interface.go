package parts

type PartsService interface {
	GetAllParts() ([]Part, error)
}

type PartsRepository interface {
	GetParts() ([]Part, error)
}
