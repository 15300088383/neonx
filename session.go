package neonx

import (
	"context"

	"google.golang.org/grpc/metadata"
	"github.com/spf13/viper"
)

/*
_uid	用户ID	N	标示当前用户
_token	用户Token	N	当前用户令牌
_platform	接入平台	Y	IOS,Android,Web,WechatApp,Windows，Linux，...
_version	客户端版本	Y	1.0.0
_net	当前网络环境	N	2G,3G,4G,5G,wifi,unknown
_mobile	用户手机号	N	15644441111
_os	操作系统	N	系统获取的操作系统名称
_device	设备ID	N	客户端计算唯一的设备ID
_describe	设备描述	N	OPPO R33
_trace	调用链跟踪ID	Y	客户端生成唯一UUID
_sequence	调用序列	Y	客户端每次调用时加一
_time	调用时间戳	Y	调用时候客户端当前时间戳
_stack 服务各系统间的调用栈
_chain 服务调用链条
*/

type Session struct {
	Uid      string
	Token    string
	Platform string
	Version  string
	Net      string
	Mobile   string
	OS       string
	Device   string
	Describe string
	Trace    string
	Sequence string
	Time     string
}

func (m *Session) Keys() map[string]*string {
	return map[string]*string{
		"_uid":      &m.Uid,
		"_token":    &m.Token,
		"_platform": &m.Platform,
		"_version":  &m.Version,
		"_net":      &m.Net,
		"_mobile":   &m.Mobile,
		"_os":       &m.OS,
		"_device":   &m.Device,
		"_describe": &m.Describe,
		"_trace":    &m.Trace,
		"_sequence": &m.Sequence,
		"_time":     &m.Time,
	}
}


func (m *Session) Encode() map[string]string {

	return map[string]string{
		"_uid":      m.Uid,
		"_token":    m.Token,
		"_platform": m.Platform,
		"_version":  m.Version,
		"_net":      m.Net,
		"_mobile":   m.Mobile,
		"_os":       m.OS,
		"_device":   m.Device,
		"_describe": m.Describe,
		"_trace":    m.Trace,
		"_sequence": m.Sequence,
		"_time":     m.Time,
	}
}

func CreateSessionFromContext(ctx context.Context) *Session {
	session := &Session{}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return session
	}
	session.Uid = md.Get("_uid")[0]
	session.Trace = md.Get("_trace")[0]
	session.Token = md.Get("_token")[0]
	session.Describe = md.Get("_describe")[0]
	session.Device = md.Get("_device")[0]
	session.Mobile = md.Get("_mobile")[0]
	session.Net = md.Get("_net")[0]
	session.OS = md.Get("_os")[0]
	session.Platform = md.Get("_platform")[0]
	session.Sequence = md.Get("_sequence")[0]
	session.Time = md.Get("_time")[0]
	session.Version = md.Get("_version")[0]

	return session
}

func getSystemName() string {
	return viper.GetString("Name")
}
