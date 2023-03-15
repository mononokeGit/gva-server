package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/meetcalendar"
	"time"
)

type MeetcalendarSearch struct {
	meetcalendar.Meetcalendar
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartMDate     *time.Time `json:"startMDate" form:"startMDate"`
	EndMDate       *time.Time `json:"endMDate" form:"endMDate"`
	StartStime     *time.Time `json:"startStime" form:"startStime"`
	EndStime       *time.Time `json:"endStime" form:"endStime"`
	StartPtime     *int       `json:"startPtime" form:"startPtime"`
	EndPtime       *int       `json:"endPtime" form:"endPtime"`
	request.PageInfo
}
