// 自动生成模板BasicRoomS
package BasicRoom

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// BasicRoomS 结构体
type BasicRoomS struct {
	global.GVA_MODEL
	Broom     *int   `json:"broom" form:"broom" gorm:"column:broom;comment:会议室;"`
	Bamounts  *int   `json:"bamounts" form:"bamounts" gorm:"column:bamounts;comment:可容纳人数;"`
	Bstatus   *bool  `json:"bstatus" form:"bstatus" gorm:"column:bstatus;comment:会议室状态;"`
	Bmarks    string `json:"bmarks" form:"bmarks" gorm:"column:bmarks;comment:备注;"`
	Bother    string `json:"bother" form:"bother" gorm:"column:bother;comment:其他;"`
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName BasicRoomS 表名
func (BasicRoomS) TableName() string {
	return "BasicRoomT"
}
