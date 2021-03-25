module github.com/gift4u/main.go

go 1.16

require (
	docs v0.0.0-00010101000000-000000000000 // indirect
	gift4u/config v0.0.0-00010101000000-000000000000
	gift4u/db v0.0.0-00010101000000-000000000000
	gift4u/models v0.0.0-00010101000000-000000000000
	gift4u/server v0.0.0-00010101000000-000000000000
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14 // indirect
	github.com/swaggo/gin-swagger v1.3.0 // indirect
	github.com/swaggo/swag v1.7.0
)

replace gift4u/db => ./db

replace gift4u/config => ./config

replace gift4u/server => ./server

replace gift4u/models => ./models

replace gift4u/controllers => ./controllers

replace docs => ./docs
