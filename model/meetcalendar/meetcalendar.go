// 自动生成模板Meetcalendar
package meetcalendar

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Meetcalendar 结构体
type Meetcalendar struct {
	global.GVA_MODEL
	MDate     *time.Time `json:"mDate" form:"mDate" gorm:"column:m_date;comment:会议日期;"`
	MRoom     *int       `json:"mRoom" form:"mRoom" gorm:"column:m_room;comment:会议室;size:19;"`
	MType     string     `json:"mType" form:"mType" gorm:"column:m_type;comment:会议类型;size:191;"`
	Title     string     `json:"title" form:"title" gorm:"column:title;comment:会议名称;size:191;"`
	Stime     *time.Time `json:"stime" form:"stime" gorm:"column:stime;comment:开始时间;"`
	Ptime     *int       `json:"ptime" form:"ptime" gorm:"column:ptime;comment:会议时长;size:19;"`
	Leader    string     `json:"leader" form:"leader" gorm:"column:leader;comment:上级领导;size:191;"`
	Mleader   string     `json:"mleader" form:"mleader" gorm:"column:mleader;comment:我区参会领导;size:191;"`
	Amount    *int       `json:"amount" form:"amount" gorm:"column:amount;comment:参会人数;size:19;"`
	Level     string     `json:"level" form:"level" gorm:"column:level;comment:会议等级;size:191;"`
	Resposen  string     `json:"resposen" form:"resposen" gorm:"column:resposen;comment:责任单位;size:191;"`
	Route     string     `json:"route" form:"route" gorm:"column:route;comment:会议线路;size:191;"`
	Remark    string     `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:191;"`
	CreatedBy uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Meetcalendar 表名
func (Meetcalendar) TableName() string {
	return "meetcalendar"
}
