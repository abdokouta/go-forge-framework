package repository

import (
	"time"
	"crud/model"
	"crud/querybuilder"
)

type RepositoryInterface interface {
	Create(model interface{}) error
	Find(id interface{}) (interface{}, error)
	Update(model interface{}) error
	Delete(id interface{}) error
	All() ([]interface{}, error)
	GetQueryBuilder() querybuilder.QueryBuilderInterface
	GetModel() model.ModelInterface
}

type BaseRepository struct {
	Model model.ModelInterface
	QueryBuilder querybuilder.QueryBuilderInterface
}

type Model struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}