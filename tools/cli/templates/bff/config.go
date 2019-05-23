package bff

import "html/template"

var Config = template.Must(
	template.New("").Parse(`Address: ":2243"
Servers:
	UserServer: "127.0.0.1:50054"

Metrics:
	Address: "0.0.0.0:9090"

Name: "user_bff"`))
