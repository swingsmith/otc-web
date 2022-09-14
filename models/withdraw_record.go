package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type WithdrawRecord struct {
	Id	int	`json:"id" gorm:"column:id"`
	OrderId	string	`json:"order_id" gorm:"column:order_id"`
	From	string	`json:"from" gorm:"column:from"`	
	To	string	`json:"to" gorm:"column:to"`
	Gas	decimal.Decimal	`json:"gas" gorm:"column:gas"`
	Fee	decimal.Decimal	`json:"fee" gorm:"column:fee"`
	Value	decimal.Decimal	`json:"value" gorm:"column:value"`
	Tx	string	`json:"tx" gorm:"column:tx"`	
	Type	string	`json:"type" gorm:"column:type"`	
	CreateTime	time.Time	`json:"create_time" gorm:"column:create_time"`
	UpdateTime	time.Time	`json:"update_time" gorm:"column:update_time"`
	Status	int	`json:"status" gorm:"column:status"`
	Version	int64	`json:"version" gorm:"column:version"`
	Count	int	`json:"count" gorm:"column:count"`
}
func (e *WithdrawRecord) TableName() string { 
    return "withdraw_record"
}