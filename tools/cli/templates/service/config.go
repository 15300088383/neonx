package service

import "html/template"

var Config = template.Must(
	template.New("").Parse(`Address: "127.0.0.1:50054"
DB:
  User:
    DSN: "root:Geeky.mysql.@admin1@tcp(192.168.0.202:3306)/test?charset=utf8"
    MaxIdle: 5
    MaxOpen: 5

REDIS:
  User:
    DSN: "192.168.0.202:6379"

Name: "user_service"
`))
