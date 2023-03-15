package Lphone

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Lphone"
	LphoneReq "github.com/flipped-aurora/gin-vue-admin/server/model/Lphone/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LPhoneApi struct {
}

var lPhoneService = service.ServiceGroupApp.LphoneServiceGroup.LPhoneService

// CreateLPhone 创建LPhone
// @Tags LPhone
// @Summary 创建LPhone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Lphone.LPhone true "创建LPhone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lPhone/createLPhone [post]
func (lPhoneApi *LPhoneApi) CreateLPhone(c *gin.Context) {
	var lPhone Lphone.LPhone
	err := c.ShouldBindJSON(&lPhone)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	lPhone.CreatedBy = utils.GetUserID(c)
	if err := lPhoneService.CreateLPhone(lPhone); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLPhone 删除LPhone
// @Tags LPhone
// @Summary 删除LPhone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Lphone.LPhone true "删除LPhone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lPhone/deleteLPhone [delete]
func (lPhoneApi *LPhoneApi) DeleteLPhone(c *gin.Context) {
	var lPhone Lphone.LPhone
	err := c.ShouldBindJSON(&lPhone)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	lPhone.DeletedBy = utils.GetUserID(c)
	if err := lPhoneService.DeleteLPhone(lPhone); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLPhoneByIds 批量删除LPhone
// @Tags LPhone
// @Summary 批量删除LPhone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LPhone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /lPhone/deleteLPhoneByIds [delete]
func (lPhoneApi *LPhoneApi) DeleteLPhoneByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := lPhoneService.DeleteLPhoneByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLPhone 更新LPhone
// @Tags LPhone
// @Summary 更新LPhone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Lphone.LPhone true "更新LPhone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lPhone/updateLPhone [put]
func (lPhoneApi *LPhoneApi) UpdateLPhone(c *gin.Context) {
	var lPhone Lphone.LPhone
	err := c.ShouldBindJSON(&lPhone)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	lPhone.UpdatedBy = utils.GetUserID(c)
	if err := lPhoneService.UpdateLPhone(lPhone); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLPhone 用id查询LPhone
// @Tags LPhone
// @Summary 用id查询LPhone
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Lphone.LPhone true "用id查询LPhone"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lPhone/findLPhone [get]
func (lPhoneApi *LPhoneApi) FindLPhone(c *gin.Context) {
	var lPhone Lphone.LPhone
	err := c.ShouldBindQuery(&lPhone)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relPhone, err := lPhoneService.GetLPhone(lPhone.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relPhone": relPhone}, c)
	}
}

// GetLPhoneList 分页获取LPhone列表
// @Tags LPhone
// @Summary 分页获取LPhone列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query LphoneReq.LPhoneSearch true "分页获取LPhone列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lPhone/getLPhoneList [get]
func (lPhoneApi *LPhoneApi) GetLPhoneList(c *gin.Context) {
	var pageInfo LphoneReq.LPhoneSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lPhoneService.GetLPhoneInfoList(pageInfo); err != nil {
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
