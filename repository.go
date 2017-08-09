package gorepo


type IRepository interface {

	// Insert puts a new instance of the give IModel in the database
	Insert(model IModel) (uint, error)

	Update(model IModel) (error)

	Save(model IModel) (uint, error)

	FindById(receiver IModel, uint uint) (error)

	FindFirst(receiver IModel, where string, args ...interface{}) (error)

	FindAll(models interface{}, where string, args ...interface{}) (err error)

	Delete(model IModel, where string, args ...interface{}) error

	// NewRecord check if the model exist in the store
	NewRecord(model IModel) bool
}
