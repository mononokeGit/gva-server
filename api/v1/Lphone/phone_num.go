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

type PhoneNumApi struct {
}

var phoneNumService = service.ServiceGroupApp.LphoneServiceGroup.PhoneNumService

// CreatePhoneNum 创建PhoneNum
// @Tags PhoneNum
// @Summary 创建PhoneNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Lphone.PhoneNum true "创建PhoneNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /phoneNum/createPhoneNum [post]
func (phoneNumApi *PhoneNumApi) CreatePhoneNum(c *gin.Context) {
	var phoneNum Lphone.PhoneNum
	err := c.ShouldBindJSON(&phoneNum)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	phoneNum.CreatedBy = utils.GetUserID(c)
	if err := phoneNumService.CreatePhoneNum(phoneNum); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePhoneNum 删除PhoneNum
// @Tags PhoneNum
// @Summary 删除PhoneNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Lphone.PhoneNum true "删除PhoneNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /phoneNum/deletePhoneNum [delete]
func (phoneNumApi *PhoneNumApi) DeletePhoneNum(c *gin.Context) {
	var phoneNum Lphone.PhoneNum
	err := c.ShouldBindJSON(&phoneNum)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	phoneNum.DeletedBy = utils.GetUserID(c)
	if err := phoneNumService.DeletePhoneNum(phoneNum); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePhoneNumByIds 批量删除PhoneNum
// @Tags PhoneNum
// @Summary 批量删除PhoneNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PhoneNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /phoneNum/deletePhoneNumByIds [delete]
func (phoneNumApi *PhoneNumApi) DeletePhoneNumByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := phoneNumService.DeletePhoneNumByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePhoneNum 更新PhoneNum
// @Tags PhoneNum
// @Summary 更新PhoneNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Lphone.PhoneNum true "更新PhoneNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /phoneNum/updatePhoneNum [put]
func (phoneNumApi *PhoneNumApi) UpdatePhoneNum(c *gin.Context) {
	var phoneNum Lphone.PhoneNum
	err := c.ShouldBindJSON(&phoneNum)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	phoneNum.UpdatedBy = utils.GetUserID(c)
	if err := phoneNumService.UpdatePhoneNum(phoneNum); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPhoneNum 用id查询PhoneNum
// @Tags PhoneNum
// @Summary 用id查询PhoneNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Lphone.PhoneNum true "用id查询PhoneNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /phoneNum/findPhoneNum [get]
func (phoneNumApi *PhoneNumApi) FindPhoneNum(c *gin.Context) {
	var phoneNum Lphone.PhoneNum
	err := c.ShouldBindQuery(&phoneNum)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rephoneNum, err := phoneNumService.GetPhoneNum(phoneNum.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rephoneNum": rephoneNum}, c)
	}
}

// GetPhoneNumList 分页获取PhoneNum列表
// @Tags PhoneNum
// @Summary 分页获取PhoneNum列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query LphoneReq.PhoneNumSearch true "分页获取PhoneNum列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /phoneNum/getPhoneNumList [get]
func (phoneNumApi *PhoneNumApi) GetPhoneNumList(c *gin.Context) {
	var pageInfo LphoneReq.PhoneNumSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := phoneNumService.GetPhoneNumInfoList(pageInfo); err != nil {
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
