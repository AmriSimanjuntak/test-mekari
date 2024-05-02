package usecase

import (
	"Test_Mek/domain/model"
	"Test_Mek/domain/repository"
	"errors"

	"gorm.io/gorm"
)

type employeeUsecase struct {
	db           *gorm.DB
	employeeRepo repository.IEmployee
}

type IEmployeeUsecase interface {
	ListEmployeeUseCase(*model.Employee) *model.BaseResp
	AddEmployeeUseCase(*model.Employee) *model.BaseResp
	EditEmployeeUseCase(*model.Employee) *model.BaseResp
	ViewEmployeeUseCase(id string) *model.BaseResp
	DeleteEmployeeUseCase(id string) *model.BaseResp
}

func NewEmployeeUsecase(db *gorm.DB, employeeRepo repository.IEmployee) IEmployeeUsecase {
	return &employeeUsecase{
		db:           db,
		employeeRepo: employeeRepo,
	}
}

func (u *employeeUsecase) ListEmployeeUseCase(req *model.Employee) *model.BaseResp {

	dataEmployees, err := u.employeeRepo.FindAll(req)
	if err != nil {

		return new(model.BaseResp).Err(errors.New("Failed Get Data"))
	}
	return new(model.BaseResp).OK(dataEmployees)

}
func (u *employeeUsecase) AddEmployeeUseCase(req *model.Employee) *model.BaseResp {
	var err error
	employeeData := model.Employee{}

	err = u.employeeRepo.CreateOne(&employeeData)
	if err != nil {
		return new(model.BaseResp).Err(errors.New("Failed Save Data"))
	}
	return new(model.BaseResp).OK("")
}

func (u *employeeUsecase) EditEmployeeUseCase(req *model.Employee) *model.BaseResp {
	var err error
	employeeData := model.Employee{}

	err = u.employeeRepo.UpdateOne(&employeeData)
	if err != nil {
		return new(model.BaseResp).Err(errors.New("Failed Update Data"))
	}
	return new(model.BaseResp).OK("")
}

func (u *employeeUsecase) DeleteEmployeeUseCase(id string) *model.BaseResp {
	dataEmployee, err := u.employeeRepo.FindOneFirst(id)
	if err != nil {

		return new(model.BaseResp).Err(errors.New("Failed Get Data"))
	}
	return new(model.BaseResp).OK(dataEmployee)
}

func (u *employeeUsecase) ViewEmployeeUseCase(id string) *model.BaseResp {
	err := u.employeeRepo.DeleteOne(id)
	if err != nil {
		return new(model.BaseResp).Err(errors.New("Failed Delete Data"))
	}
	return new(model.BaseResp).OK("")
}
