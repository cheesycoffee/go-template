package reposql

type repoSQL struct{}

// RepoSQL product
type RepoSQL interface{}

// New sql repository product
func New() RepoSQL {
	return &repoSQL{}
}
