package meetcalendar

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/meetcalendar"
	meetcalendarReq "github.com/flipped-aurora/gin-vue-admin/server/model/meetcalendar/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MeetcalendarApi struct {
}

var meetcalendarService = service.ServiceGroupApp.MeetcalendarServiceGroup.MeetcalendarService

// CreateMeetcalendar 创建Meetcalendar
// @Tags Meetcalendar
// @Summary 创建Meetcalendar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body meetcalendar.Meetcalendar true "创建Meetcalendar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /meetcalendar/createMeetcalendar [post]
func (meetcalendarApi *MeetcalendarApi) CreateMeetcalendar(c *gin.Context) {
	var meetcalendar meetcalendar.Meetcalendar
	err := c.ShouldBindJSON(&meetcalendar)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	meetcalendar.CreatedBy = utils.GetUserID(c)
	if err := meetcalendarService.CreateMeetcalendar(meetcalendar); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMeetcalendar 删除Meetcalendar
// @Tags Meetcalendar
// @Summary 删除Meetcalendar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body meetcalendar.Meetcalendar true "删除Meetcalendar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /meetcalendar/deleteMeetcalendar [delete]
func (meetcalendarApi *MeetcalendarApi) DeleteMeetcalendar(c *gin.Context) {
	var meetcalendar meetcalendar.Meetcalendar
	err := c.ShouldBindJSON(&meetcalendar)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	meetcalendar.DeletedBy = utils.GetUserID(c)
	if err := meetcalendarService.DeleteMeetcalendar(meetcalendar); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMeetcalendarByIds 批量删除Meetcalendar
// @Tags Meetcalendar
// @Summary 批量删除Meetcalendar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Meetcalendar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /meetcalendar/deleteMeetcalendarByIds [delete]
func (meetcalendarApi *MeetcalendarApi) DeleteMeetcalendarByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := meetcalendarService.DeleteMeetcalendarByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMeetcalendar 更新Meetcalendar
// @Tags Meetcalendar
// @Summary 更新Meetcalendar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body meetcalendar.Meetcalendar true "更新Meetcalendar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /meetcalendar/updateMeetcalendar [put]
func (meetcalendarApi *MeetcalendarApi) UpdateMeetcalendar(c *gin.Context) {
	var meetcalendar meetcalendar.Meetcalendar
	err := c.ShouldBindJSON(&meetcalendar)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	meetcalendar.UpdatedBy = utils.GetUserID(c)
	if err := meetcalendarService.UpdateMeetcalendar(meetcalendar); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMeetcalendar 用id查询Meetcalendar
// @Tags Meetcalendar
// @Summary 用id查询Meetcalendar
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query meetcalendar.Meetcalendar true "用id查询Meetcalendar"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /meetcalendar/findMeetcalendar [get]
func (meetcalendarApi *MeetcalendarApi) FindMeetcalendar(c *gin.Context) {
	var meetcalendar meetcalendar.Meetcalendar
	err := c.ShouldBindQuery(&meetcalendar)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remeetcalendar, err := meetcalendarService.GetMeetcalendar(meetcalendar.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remeetcalendar": remeetcalendar}, c)
	}
}

// GetMeetcalendarList 分页获取Meetcalendar列表
// @Tags Meetcalendar
// @Summary 分页获取Meetcalendar列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query meetcalendarReq.MeetcalendarSearch true "分页获取Meetcalendar列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /meetcalendar/getMeetcalendarList [get]
func (meetcalendarApi *MeetcalendarApi) GetMeetcalendarList(c *gin.Context) {
	var pageInfo meetcalendarReq.MeetcalendarSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := meetcalendarService.GetMeetcalendarInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
