package Lphone

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LPhoneRouter struct {
}

// InitLPhoneRouter 初始化 LPhone 路由信息
func (s *LPhoneRouter) InitLPhoneRouter(Router *gin.RouterGroup) {
	lPhoneRouter := Router.Group("lPhone").Use(middleware.OperationRecord())
	lPhoneRouterWithoutRecord := Router.Group("lPhone")
	var lPhoneApi = v1.ApiGroupApp.LphoneApiGroup.LPhoneApi
	{
		lPhoneRouter.POST("createLPhone", lPhoneApi.CreateLPhone)             // 新建LPhone
		lPhoneRouter.DELETE("deleteLPhone", lPhoneApi.DeleteLPhone)           // 删除LPhone
		lPhoneRouter.DELETE("deleteLPhoneByIds", lPhoneApi.DeleteLPhoneByIds) // 批量删除LPhone
		lPhoneRouter.PUT("updateLPhone", lPhoneApi.UpdateLPhone)              // 更新LPhone
	}
	{
		lPhoneRouterWithoutRecord.GET("findLPhone", lPhoneApi.FindLPhone)       // 根据ID获取LPhone
		lPhoneRouterWithoutRecord.GET("getLPhoneList", lPhoneApi.GetLPhoneList) // 获取LPhone列表
	}
}
