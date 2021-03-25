module gift4u/server

go 1.16

require (
	gift4u/controllers v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.1.13 // indirect
)

replace gift4u/controllers => ../controllers

replace gift4u/models => ../models

replace gift4u/config => ../config

replace gift4u/db => ../db
