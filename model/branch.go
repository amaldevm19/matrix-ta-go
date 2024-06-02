package model

import "gorm.io/gorm"

type Px_BranchMst struct {
	gorm.Model
	BranchId    uint   `gorm:"primaryKey autoIncrement"`
	BranchName  string `gorm:"not null" json:"branch_name"`
	BranchCode  string `gorm:"not null" json:"branch_code"`
	TnaBranchId string
}
