package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/BasicRoom"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Lphone"
	"github.com/flipped-aurora/gin-vue-admin/server/service/applyroom"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/meetcalendar"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	ExampleServiceGroup      example.ServiceGroup
	LphoneServiceGroup       Lphone.ServiceGroup
	MeetcalendarServiceGroup meetcalendar.ServiceGroup
	ApplyroomServiceGroup    applyroom.ServiceGroup
	BasicroomServiceGroup    BasicRoom.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
