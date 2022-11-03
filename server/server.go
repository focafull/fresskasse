package server

import (
	"akaflieg/fresskasse/v2/config"
	"fmt"
)

func Init() {
	config := config.GetConfig()
	router := NewRouter()
	router.Run(fmt.Sprintf("%s:%s", config.ADDRESS, config.PORT))
}
