module gift4u/controllers

go 1.16

replace gift4u/models => ../models

require (
	gift4u/models v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
)

replace gift4u/config => ../config

replace gift4u/db => ../db
