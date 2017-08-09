package gorepo

import "github.com/jinzhu/gorm"

type GormRepository struct {
	BaseRepository
	Db *gorm.DB
}

func (b *GormRepository) InitDb(db *gorm.DB)  {
	b.Db = db
}

func (r *GormRepository) Initialize(db * gorm.DB) {
	r.InitDb(db)
}

func (r GormRepository) Insert(model IModel) (uint, error){
	if err := model.Validate(); err != nil{
		return 0, err
	}
	if err := r.Db.Create(model).Error; err != nil{
		return 0, err
	}
	return model.GetId(), nil
}

func (r GormRepository) Update(model IModel) (error){
	if err := model.Validate(); err != nil{
		return err
	}
	return r.Db.Save(model).Error
}

func (r GormRepository) Save(model IModel) (uint, error){
	if err := model.Validate(); err != nil{
		return 0, err
	}
	if err := r.Db.Save(model).Error; err != nil{
		return 0, err
	}
	return model.GetId(), nil
}

func (r GormRepository) FindById(receiver IModel, uint uint) (error){
	return r.Db.First(receiver).Error
}

func (r GormRepository) FindFirst(receiver IModel, where string, args ...interface{}) (error){
	return r.Db.Where(where, args).Limit(1).Find(receiver).Error
}

func (r GormRepository) FindAll(models interface{}, where string, args ...interface{}) (err error){
	err = r.Db.Where(where, args).Find(models).Error
	return
}

func (r GormRepository) Delete(model IModel, where string, args ...interface{}) error {
	return r.Db.Where(where, args).Delete(&model).Error
}

func (r GormRepository) NewRecord(model IModel) bool {
	return r.Db.NewRecord(&model)
}


