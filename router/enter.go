package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Lphone"
	"github.com/flipped-aurora/gin-vue-admin/server/router/applyroom"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/meetcalendar"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System       system.RouterGroup
	Example      example.RouterGroup
	Lphone       Lphone.RouterGroup
	Meetcalendar meetcalendar.RouterGroup
	Applyroom    applyroom.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
