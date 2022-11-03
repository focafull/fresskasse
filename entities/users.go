package entities

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	Id            string `bun:"id,pk,default:gen_random_uuid()"`
	Name          string `bun:"name,notnull"`
	Balance       float64
}
