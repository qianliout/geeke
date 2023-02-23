package main

import (
	"outback/stack/pipline"
	"outback/stack/spiders"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/stack?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Error().Err(err)
		return
	}
	pip := pipline.NewCreate(db)

	sp := spiders.NewStarkSpider(pip)
	// sp := spiders.NewNameCode(pip)
	//
	// sp.ListSh()
	sp.Start()
}
