package model

import (
	"time"
)

type BaseResp struct {
	ResponseCode string      `json:"responseCode"`
	ResponseDesc string      `json:"responseDesc"`
	Code         int         `json:"code"` // http status at header response api-gateway
	Data         interface{} `json:"data"`
}
type Employee struct {
	ID        int       `gorm:"primaryKey"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email"`
	HireDate  time.Time `gorm:"column:hire_date"`
}

// func (br *BaseResp) OK(data interface{}) *BaseResp {
// 	return &BaseResp{ResponseCode: "00", ResponseDesc: "success", Data: data}
// }
