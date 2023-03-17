package applyroom

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/applyroom"
	applyroomReq "github.com/flipped-aurora/gin-vue-admin/server/model/applyroom/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApproveApi struct {
}

var approveService = service.ServiceGroupApp.ApplyroomServiceGroup.ApproveService

// CreateApprove 创建Approve
// @Tags Approve
// @Summary 创建Approve
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body applyroom.Approve true "创建Approve"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /approve/createApprove [post]
func (approveApi *ApproveApi) CreateApprove(c *gin.Context) {
	var approve applyroom.Approve
	err := c.ShouldBindJSON(&approve)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	approve.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Appunit": {utils.NotEmpty()},
		"Approom": {utils.NotEmpty()},
	}
	if err := utils.Verify(approve, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := approveService.CreateApprove(approve); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteApprove 删除Approve
// @Tags Approve
// @Summary 删除Approve
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body applyroom.Approve true "删除Approve"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /approve/deleteApprove [delete]
func (approveApi *ApproveApi) DeleteApprove(c *gin.Context) {
	var approve applyroom.Approve
	err := c.ShouldBindJSON(&approve)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	approve.DeletedBy = utils.GetUserID(c)
	if err := approveService.DeleteApprove(approve); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteApproveByIds 批量删除Approve
// @Tags Approve
// @Summary 批量删除Approve
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Approve"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /approve/deleteApproveByIds [delete]
func (approveApi *ApproveApi) DeleteApproveByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := approveService.DeleteApproveByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateApprove 更新Approve
// @Tags Approve
// @Summary 更新Approve
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body applyroom.Approve true "更新Approve"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /approve/updateApprove [put]
func (approveApi *ApproveApi) UpdateApprove(c *gin.Context) {
	var approve applyroom.Approve
	err := c.ShouldBindJSON(&approve)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	approve.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Appunit": {utils.NotEmpty()},
		"Approom": {utils.NotEmpty()},
	}
	if err := utils.Verify(approve, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := approveService.UpdateApprove(approve); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindApprove 用id查询Approve
// @Tags Approve
// @Summary 用id查询Approve
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyroom.Approve true "用id查询Approve"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /approve/findApprove [get]
func (approveApi *ApproveApi) FindApprove(c *gin.Context) {
	var approve applyroom.Approve
	err := c.ShouldBindQuery(&approve)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reapprove, err := approveService.GetApprove(approve.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapprove": reapprove}, c)
	}
}

// GetApproveList 分页获取Approve列表
// @Tags Approve
// @Summary 分页获取Approve列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyroomReq.ApproveSearch true "分页获取Approve列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /approve/getApproveList [get]
func (approveApi *ApproveApi) GetApproveList(c *gin.Context) {
	var pageInfo applyroomReq.ApproveSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := approveService.GetApproveInfoList(pageInfo); err != nil {
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
