package model

import "time"

type Model struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type ModelInterface interface {
	GetID() interface{}
	TableName() string
	SetID(id interface{})
}

func NewModel() *Model {
	return &Model{}
}