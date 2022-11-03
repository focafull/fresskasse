package entities

import "github.com/uptrace/bun"

type Meal struct {
	bun.BaseModel `bun:"table:meals,alias:m"`
	ID            string `bun:"id,pk,default:gen_random_uuid()"`
	PRICE         float64
	USERS         []string `bun:"users,array"`
}
