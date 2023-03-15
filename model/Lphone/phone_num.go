// 自动生成模板PhoneNum
package Lphone

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PhoneNum 结构体
type PhoneNum struct {
	global.GVA_MODEL
	Uid       string `json:"uid" form:"uid" gorm:"column:uid;comment:标识符;size:255;"`
	NOther    string `json:"nOther" form:"nOther" gorm:"column:n_other;comment:备注;size:255;"`
	NBackup   string `json:"nBackup" form:"nBackup" gorm:"column:n_backup;comment:预留字段;size:255;"`
	NNumber   *int   `json:"nNumber" form:"nNumber" gorm:"column:n_number;comment:号码;size:19;"`
	NRes      string `json:"nRes" form:"nRes" gorm:"column:n_res;comment:分配的人或科室;size:255;"`
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName PhoneNum 表名
func (PhoneNum) TableName() string {
	return "phone_num"
}
