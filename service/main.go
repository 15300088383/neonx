package service

import (
	"net"

	"github.com/geekymedic/neonx/logger"

	"github.com/geekymedic/neonx"
	"github.com/geekymedic/neonx/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	_rpcServer = grpc.NewServer()
)

func Main() error {

	cmd := cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {

			var (
				logger = logger.NewLogger(nil)
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
				logger.LogError("listen %s fail,%s", address, err)
				return err
			} else {
				logger.LogInfo("listen %s", l.Addr())
			}

			err = _rpcServer.Serve(l)

			if err != nil {
				logger.LogError("%s", err)
				return err
			}

			return nil
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&_flags.Debug, "debug", false, "use debug mode")

	return cmd.Execute()

}

func Server() *grpc.Server {
	return _rpcServer
}
