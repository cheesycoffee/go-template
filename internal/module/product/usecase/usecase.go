package usecase

type usecase struct{}

// UseCase product
type UseCase interface{}

// New usecase product
func New() UseCase {
	return &usecase{}
}
