package rpc

import (
	"strings"

	"github.com/geekymedic/neonx/errors"
	"github.com/geekymedic/neonx"
	
	"google.golang.org/grpc"
	"github.com/spf13/viper"
)

var connections = map[string]*grpc.ClientConn{}

func init() {

	neonx.AddPlugin("rpc_server", func(status neonx.PluginStatus, viper *viper.Viper) error {
		switch status {
		case neonx.PluginLoad:
			servers := viper.GetStringMapString("servers")
			for name, address := range servers {
				conn, err := grpc.Dial(address, grpc.WithInsecure())
				if err != nil {
					return errors.By(err)
				}
				connections[name] = conn
			}
		}
		return nil

	})
}

func GetConnection(name string) *grpc.ClientConn {
	return connections[strings.ToLower(name)]
}
