package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Lphone"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PhoneNumSearch struct {
	Lphone.PhoneNum
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartNNumber   *int       `json:"startNNumber" form:"startNNumber"`
	EndNNumber     *int       `json:"endNNumber" form:"endNNumber"`
	request.PageInfo
}
