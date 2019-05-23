package bff

import (
	"github.com/gin-gonic/gin"
)

var (
	_engine = gin.Default()
	_group  = _engine.Group("/apis")
)

func init() {
	_engine.Use(
		metricsMiddleWare)
}

func Engine() *gin.RouterGroup {
	return _group
}
