package gorepo

type IModel interface {
	GetId() uint

	Validate() error
}

