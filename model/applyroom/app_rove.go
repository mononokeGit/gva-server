// 自动生成模板Approve
package applyroom

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Approve 结构体
type Approve struct {
	global.GVA_MODEL
	Appunit   string     `json:"appunit" form:"appunit" gorm:"column:appunit;comment:审批部门;"`
	Appleader string     `json:"appleader" form:"appleader" gorm:"column:appleader;comment:参会领导;"`
	Approom   *int       `json:"approom" form:"approom" gorm:"column:approom;comment:会议室;"`
	Appamount *int       `json:"appamount" form:"appamount" gorm:"column:appamount;comment:参会人数;"`
	Apptype   string     `json:"apptype" form:"apptype" gorm:"column:apptype;comment:会议类型;"`
	Apptime   *time.Time `json:"apptime" form:"apptime" gorm:"column:apptime;comment:使用时间;"`
	Appdevice string     `json:"appdevice" form:"appdevice" gorm:"column:appdevice;comment:是否需要设备;"`
	Appres1   string     `json:"appres1" form:"appres1" gorm:"column:appres1;comment:;"`
	Appres2   string     `json:"appres2" form:"appres2" gorm:"column:appres2;comment:;"`
	CreatedBy uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Approve 表名
func (Approve) TableName() string {
	return "approve"
}
