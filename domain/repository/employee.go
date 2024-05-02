package repository

import (
	"Test_Mek/domain/model"
)

type IEmployee interface {
	CreateOne(data *model.Employee) error
	UpdateOne(data *model.Employee) error
	FindOneFirst(id string) (data *model.Employee, err error)
	FindAll(qry *model.Employee) (data []*model.Employee, err error)
	DeleteOne(id string) error
}
