package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/applyroom"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ApproveSearch struct{
    applyroom.Approve
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartApptime  *time.Time  `json:"startApptime" form:"startApptime"`
    EndApptime  *time.Time  `json:"endApptime" form:"endApptime"`
    request.PageInfo
}
