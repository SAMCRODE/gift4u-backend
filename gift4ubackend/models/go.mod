module gift4u/models

go 1.16

require (
	gift4u/config v0.0.0-00010101000000-000000000000
	gift4u/db v0.0.0-00010101000000-000000000000
	github.com/go-pg/pg/v10 v10.8.0
)

replace gift4u/db => ../db

replace gift4u/config => ../config
