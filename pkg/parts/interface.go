package parts

type PartsService interface {
}

type PartsRepository interface {
	GetParts() ([]Part, error)
}
