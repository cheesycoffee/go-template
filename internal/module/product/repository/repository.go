package repository

import (
	"github.com/cheesycoffee/go-template/internal/module/product/repository/internal/repomongo"
	"github.com/cheesycoffee/go-template/internal/module/product/repository/internal/reposql"
)

type repository struct {
	repoMongo repomongo.RepoMongo
	repoSQL   reposql.RepoSQL
}

// Repository product
type Repository interface {
	Mongo() repomongo.RepoMongo
	SQL() reposql.RepoSQL
}

// New repository producr
func New() Repository {
	return &repository{
		repoMongo: repomongo.New(),
		repoSQL:   reposql.New(),
	}
}

func (r *repository) Mongo() repomongo.RepoMongo {
	return r.repoMongo
}
func (r *repository) SQL() reposql.RepoSQL {
	return r.repoSQL
}
