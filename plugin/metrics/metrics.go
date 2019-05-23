package metrics

import (
	"net"
	"net/http"

	"github.com/geekymedic/neonx"
	"github.com/geekymedic/neonx/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func metricsHandler(l net.Listener) {
	err := http.Serve(l, promhttp.Handler())

	if err != nil {
		logrus.Errorf("new service for metrics fail: %s", err)
	}
}

func init() {

	neonx.AddPlugin("metrics", func(status neonx.PluginStatus, viper *viper.Viper) error {
		switch status {
		case neonx.PluginLoad:

			addr := viper.GetString("Metrics.Address")

			if len(addr) == 0 {
				return errors.Format("load Metrics.Address fail, empty address.")
			}

			l, err := net.Listen("tcp", addr)

			if err != nil {
				return errors.By(err)
			}

			go metricsHandler(l)

		}

		return nil
	})

}
