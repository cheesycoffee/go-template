package repomongo

type repoMongo struct{}

// RepoMongo product
type RepoMongo interface{}

// New mongo repository product
func New() RepoMongo {
	return &repoMongo{}
}
