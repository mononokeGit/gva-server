package meetcalendar

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MeetcalendarRouter struct {
}

// InitMeetcalendarRouter 初始化 Meetcalendar 路由信息
func (s *MeetcalendarRouter) InitMeetcalendarRouter(Router *gin.RouterGroup) {
	meetcalendarRouter := Router.Group("meetcalendar").Use(middleware.OperationRecord())
	meetcalendarRouterWithoutRecord := Router.Group("meetcalendar")
	var meetcalendarApi = v1.ApiGroupApp.MeetcalendarApiGroup.MeetcalendarApi
	{
		meetcalendarRouter.POST("createMeetcalendar", meetcalendarApi.CreateMeetcalendar)             // 新建Meetcalendar
		meetcalendarRouter.DELETE("deleteMeetcalendar", meetcalendarApi.DeleteMeetcalendar)           // 删除Meetcalendar
		meetcalendarRouter.DELETE("deleteMeetcalendarByIds", meetcalendarApi.DeleteMeetcalendarByIds) // 批量删除Meetcalendar
		meetcalendarRouter.PUT("updateMeetcalendar", meetcalendarApi.UpdateMeetcalendar)              // 更新Meetcalendar
	}
	{
		meetcalendarRouterWithoutRecord.GET("findMeetcalendar", meetcalendarApi.FindMeetcalendar)       // 根据ID获取Meetcalendar
		meetcalendarRouterWithoutRecord.GET("getMeetcalendarList", meetcalendarApi.GetMeetcalendarList) // 获取Meetcalendar列表
	}
}
