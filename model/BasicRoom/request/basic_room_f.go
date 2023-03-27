package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/BasicRoom"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BasicRoomSSearch struct {
	BasicRoom.BasicRoomS
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartBamounts  *int       `json:"startBamounts" form:"startBamounts"`
	EndBamounts    *int       `json:"endBamounts" form:"endBamounts"`
	request.PageInfo
}
