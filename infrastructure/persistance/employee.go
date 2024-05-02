package persistance

import (
	"Test_Mek/domain/model"
	"Test_Mek/domain/repository"

	"gorm.io/gorm"
)

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) repository.IEmployee {
	return &employeeRepo{db: db}
}
func (r *employeeRepo) CreateOne(data *model.Employee) error {

	return r.db.Create(&data).Error
}
func (r *employeeRepo) UpdateOne(data *model.Employee) error {

	return r.db.Updates(&data).Error
}

func (r *employeeRepo) DeleteOne(id string) error {

	return r.db.Where(id).Delete(&model.Employee{}).Error
}

func (r *employeeRepo) FindOneFirst(id string) (*model.Employee, error) {
	data := new(model.Employee)

	err := r.db.Table("public.employee em").Select("em.*").Where("em.id = ?", id).Scan(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (r *employeeRepo) FindAll(qry *model.Employee) (data []*model.Employee, err error) {
	if err = r.db.Where(qry).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
