package BasicRoom

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BasicRoomSRouter struct {
}

// InitBasicRoomSRouter 初始化 BasicRoomS 路由信息
func (s *BasicRoomSRouter) InitBasicRoomSRouter(Router *gin.RouterGroup) {
	basicroomRouter := Router.Group("basicroom").Use(middleware.OperationRecord())
	basicroomRouterWithoutRecord := Router.Group("basicroom")
	var basicroomApi = v1.ApiGroupApp.BasicroomApiGroup.BasicRoomSApi
	{
		basicroomRouter.POST("createBasicRoomS", basicroomApi.CreateBasicRoomS)             // 新建BasicRoomS
		basicroomRouter.DELETE("deleteBasicRoomS", basicroomApi.DeleteBasicRoomS)           // 删除BasicRoomS
		basicroomRouter.DELETE("deleteBasicRoomSByIds", basicroomApi.DeleteBasicRoomSByIds) // 批量删除BasicRoomS
		basicroomRouter.PUT("updateBasicRoomS", basicroomApi.UpdateBasicRoomS)              // 更新BasicRoomS
	}
	{
		basicroomRouterWithoutRecord.GET("findBasicRoomS", basicroomApi.FindBasicRoomS)       // 根据ID获取BasicRoomS
		basicroomRouterWithoutRecord.GET("getBasicRoomSList", basicroomApi.GetBasicRoomSList) // 获取BasicRoomS列表
	}
}
