module github.com/gift4u/main.go

go 1.16

require (
	gift4u/config v0.0.0-00010101000000-000000000000
	gift4u/db v0.0.0-00010101000000-000000000000
	gift4u/models v0.0.0-00010101000000-000000000000
	gift4u/server v0.0.0-00010101000000-000000000000
)

replace gift4u/db => ./db

replace gift4u/config => ./config

replace gift4u/server => ./server

replace gift4u/models => ./models

replace gift4u/controllers => ./controllers
