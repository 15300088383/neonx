package bff

import (
	_ "github.com/geekymedic/neonx/plugin/metrics"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	MetricsHttpRequest = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "geekymedic",
			Subsystem: "bff",
			Name:      "HttpRequest",
			Help:      "Request for context",
		},
		[]string{
			"uri",
		},
	)
)

func metricsMiddleWare(ctx *gin.Context) {

	MetricsHttpRequest.With(
		prometheus.Labels{
			"uri": ctx.Request.URL.Path,
		},
	).Inc()

	ctx.Next()
}
