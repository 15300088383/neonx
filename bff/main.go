package bff

import (
	"net"
	"net/http"

	"github.com/geekymedic/neonx"
	"github.com/geekymedic/neonx/config"
	"github.com/geekymedic/neonx/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Main() error {

	var (
		logger = logrus.New()
	)

	err := config.Load()

	if err != nil {
		return err
	}

	var (
		address = viper.GetString("Address")
	)

	err = neonx.LoadPlugins(viper.GetViper())

	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", address)

	if err != nil {
		logger.Errorf("listen %s fail, %v\n", address, err)
		return errors.By(err)
	} else {
		logger.Infof("listen %s", l.Addr())
	}

	return http.Serve(l, _engine)
}
