package BasicRoom

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/BasicRoom"
	BasicRoomReq "github.com/flipped-aurora/gin-vue-admin/server/model/BasicRoom/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BasicRoomSApi struct {
}

var basicroomService = service.ServiceGroupApp.BasicroomServiceGroup.BasicRoomSService

// CreateBasicRoomS 创建BasicRoomS
// @Tags BasicRoomS
// @Summary 创建BasicRoomS
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body BasicRoom.BasicRoomS true "创建BasicRoomS"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicroom/createBasicRoomS [post]
func (basicroomApi *BasicRoomSApi) CreateBasicRoomS(c *gin.Context) {
	var basicroom BasicRoom.BasicRoomS
	err := c.ShouldBindJSON(&basicroom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	basicroom.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Broom":   {utils.NotEmpty()},
		"Bstatus": {utils.NotEmpty()},
	}
	if err := utils.Verify(basicroom, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := basicroomService.CreateBasicRoomS(basicroom); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBasicRoomS 删除BasicRoomS
// @Tags BasicRoomS
// @Summary 删除BasicRoomS
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body BasicRoom.BasicRoomS true "删除BasicRoomS"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /basicroom/deleteBasicRoomS [delete]
func (basicroomApi *BasicRoomSApi) DeleteBasicRoomS(c *gin.Context) {
	var basicroom BasicRoom.BasicRoomS
	err := c.ShouldBindJSON(&basicroom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	basicroom.DeletedBy = utils.GetUserID(c)
	if err := basicroomService.DeleteBasicRoomS(basicroom); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBasicRoomSByIds 批量删除BasicRoomS
// @Tags BasicRoomS
// @Summary 批量删除BasicRoomS
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BasicRoomS"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /basicroom/deleteBasicRoomSByIds [delete]
func (basicroomApi *BasicRoomSApi) DeleteBasicRoomSByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := basicroomService.DeleteBasicRoomSByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBasicRoomS 更新BasicRoomS
// @Tags BasicRoomS
// @Summary 更新BasicRoomS
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body BasicRoom.BasicRoomS true "更新BasicRoomS"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /basicroom/updateBasicRoomS [put]
func (basicroomApi *BasicRoomSApi) UpdateBasicRoomS(c *gin.Context) {
	var basicroom BasicRoom.BasicRoomS
	err := c.ShouldBindJSON(&basicroom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	basicroom.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Broom":   {utils.NotEmpty()},
		"Bstatus": {utils.NotEmpty()},
	}
	if err := utils.Verify(basicroom, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := basicroomService.UpdateBasicRoomS(basicroom); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBasicRoomS 用id查询BasicRoomS
// @Tags BasicRoomS
// @Summary 用id查询BasicRoomS
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query BasicRoom.BasicRoomS true "用id查询BasicRoomS"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /basicroom/findBasicRoomS [get]
func (basicroomApi *BasicRoomSApi) FindBasicRoomS(c *gin.Context) {
	var basicroom BasicRoom.BasicRoomS
	err := c.ShouldBindQuery(&basicroom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rebasicroom, err := basicroomService.GetBasicRoomS(basicroom.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebasicroom": rebasicroom}, c)
	}
}

// GetBasicRoomSList 分页获取BasicRoomS列表
// @Tags BasicRoomS
// @Summary 分页获取BasicRoomS列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query BasicRoomReq.BasicRoomSSearch true "分页获取BasicRoomS列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicroom/getBasicRoomSList [get]
func (basicroomApi *BasicRoomSApi) GetBasicRoomSList(c *gin.Context) {
	var pageInfo BasicRoomReq.BasicRoomSSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := basicroomService.GetBasicRoomSInfoList(pageInfo); err != nil {
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
