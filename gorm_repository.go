package gorepo

import "github.com/jinzhu/gorm"

type GormRepository struct {
	BaseRepository
	Db *gorm.DB
}

func NewGormRepository(db *gorm.DB) GormRepository {
	return GormRepository{Db:db}
}

func (gr *GormRepository) Initialize(args ...interface{})  {
	if len(args) == 0 {
		panic("*gorm.DB must be supplied for initialization")
	}
	if _, ok := args[0].(*gorm.DB); !ok {
		panic("The first arg must be *gorm.DB")
	}

	gr.InitDb(args[0].(*gorm.DB))
}

func (gr *GormRepository) InitDb(db *gorm.DB)  {
	gr.Db = db
}

func (gr GormRepository) Insert(model IModel) (uint, error){
	if err := model.Validate(); err != nil{
		return 0, err
	}
	if err := gr.Db.Create(model).Error; err != nil{
		return 0, err
	}
	return model.GetId(), nil
}

func (gr GormRepository) Update(model IModel) (error){
	if err := model.Validate(); err != nil{
		return err
	}
	return gr.Db.Save(model).Error
}

func (gr GormRepository) Save(model IModel) (uint, error){
	if err := model.Validate(); err != nil{
		return 0, err
	}
	if err := gr.Db.Save(model).Error; err != nil{
		return 0, err
	}
	return model.GetId(), nil
}

func (gr GormRepository) FindById(receiver IModel, id uint) (error){
	return gr.Db.First(receiver, id).Error
}

func (gr GormRepository) FindFirst(receiver IModel, where string, args ...interface{}) (error){
	return gr.Db.Where(where, args...).Limit(1).Find(receiver).Error
}

func (gr GormRepository) FindAll(models interface{}, where string, args ...interface{}) (err error){
	err = gr.Db.Where(where, args...).Find(models).Error
	return
}

func (gr GormRepository) Delete(model IModel, where string, args ...interface{}) error {
	return gr.Db.Where(where, args...).Delete(&model).Error
}

func (gr GormRepository) NewRecord(model IModel) bool {
	return gr.Db.NewRecord(&model)
}


