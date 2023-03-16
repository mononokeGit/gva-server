package applyroom

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApproveRouter struct {
}

// InitApproveRouter 初始化 Approve 路由信息
func (s *ApproveRouter) InitApproveRouter(Router *gin.RouterGroup) {
	approveRouter := Router.Group("approve").Use(middleware.OperationRecord())
	approveRouterWithoutRecord := Router.Group("approve")
	var approveApi = v1.ApiGroupApp.ApplyroomApiGroup.ApproveApi
	{
		approveRouter.POST("createApprove", approveApi.CreateApprove)   // 新建Approve
		approveRouter.DELETE("deleteApprove", approveApi.DeleteApprove) // 删除Approve
		approveRouter.DELETE("deleteApproveByIds", approveApi.DeleteApproveByIds) // 批量删除Approve
		approveRouter.PUT("updateApprove", approveApi.UpdateApprove)    // 更新Approve
	}
	{
		approveRouterWithoutRecord.GET("findApprove", approveApi.FindApprove)        // 根据ID获取Approve
		approveRouterWithoutRecord.GET("getApproveList", approveApi.GetApproveList)  // 获取Approve列表
	}
}
