module gift4u/server

go 1.16

require (
	gift4u/controllers v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
)

replace gift4u/controllers => ../controllers

replace gift4u/models => ../models
