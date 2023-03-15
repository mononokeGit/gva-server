package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Lphone"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/applyroom"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/meetcalendar"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup       system.ApiGroup
	ExampleApiGroup      example.ApiGroup
	LphoneApiGroup       Lphone.ApiGroup
	MeetcalendarApiGroup meetcalendar.ApiGroup
	ApplyroomApiGroup    applyroom.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
