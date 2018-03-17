package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/astaxie/beego"
)

type BaseDao struct {
	collection *mgo.Collection
	db         *mgo.Database
}

type FindInput struct {
	FilteringQuery interface{}
	Offset 	int
	Limit int
	SortFields []string
}

func NewDao(db *mgo.Database, collectionName string) BaseDao {
	dao := BaseDao{db:db}
	dao.collection = db.C(collectionName)
	return dao
}

func (m BaseDao) Insert(obj interface{}) ( error) {
	err := m.collection.Insert(obj)
	return err
}

func (m BaseDao) Save(obj interface{}) error {
	err := m.collection.Insert(obj)
	return err
}

func (m BaseDao) Update(selector interface{}, obj interface{}) error {
	err := m.collection.Update(selector, obj)
	return err
}

func (m BaseDao) Patch(selector interface{}, changes interface{}) error {
	_, err := m.collection.UpdateAll(selector, changes)
	return err
}

func (m BaseDao) FindAll(input FindInput, receiver interface{}) (err error) {
	query := m.collection.Find(input.FilteringQuery)
	if len(input.SortFields) > 0 {
		query = query.Sort(input.SortFields...)
	}
	if input.Limit > 0 {
		query = query.Limit(input.Limit)
	}
	if input.Offset != 0 {
		query = query.Skip(input.Offset)
	}
	err = query.All(receiver)
	return
}

func (m BaseDao) FindOne(filteringQuery interface{}, receiver interface{}) (err error) {
	query := m.collection.Find(filteringQuery)
	err = query.One(receiver)
	return
}

func (m BaseDao) FindById(id bson.ObjectId, receiver interface{}) (err error) {
	err = m.collection.FindId(id).One(receiver)
	return
}

func (m BaseDao) Exists(filteringQuery interface{}) bool {
	count, err := m.collection.Find(filteringQuery).Count()
	if err != nil{
		beego.Debug(err)
		return false
	}
	return count > 0
}

func (m BaseDao) Count(filteringQuery interface{}) (int, error) {
	return m.collection.Find(filteringQuery).Count()
}