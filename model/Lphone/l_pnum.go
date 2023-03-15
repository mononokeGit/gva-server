// 自动生成模板LPhone
package Lphone

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LPhone 结构体
type LPhone struct {
	global.GVA_MODEL
	IPhone    *int   `json:"IPhone" form:"IPhone" gorm:"column:I_phone;comment:;size:19;"`
	LName     string `json:"lName" form:"lName" gorm:"column:l_name;comment:领导姓名;size:191;"`
	LNum      *int   `json:"lNum" form:"lNum" gorm:"column:l_num;comment:领导序号;"`
	LOther    string `json:"lOther" form:"lOther" gorm:"column:l_other;comment:备注;size:191;"`
	LTitle    string `json:"lTitle" form:"lTitle" gorm:"column:l_title;comment:领导职务;size:191;"`
	LUnit     string `json:"lUnit" form:"lUnit" gorm:"column:l_unit;comment:所属单位;size:191;"`
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName LPhone 表名
func (LPhone) TableName() string {
	return "l_phone"
}
