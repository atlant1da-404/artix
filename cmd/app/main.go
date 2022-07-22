package main

import (
	"fmt"
	"github.com/atlant1da-404/artik_db/pkg/authenticate"
	"github.com/atlant1da-404/artik_db/pkg/bucket"
	"github.com/atlant1da-404/artik_db/pkg/config"
	"log"
)

func main() {

	settings := config.GetConfig("config.yaml")
	session := authenticate.New(settings)

	if err := session.Required("admin", "password"); err != nil {
		log.Fatalln(err.Error())
	}

	head := bucket.New()
	if err := head.CreateBucket("mongo"); err != nil {
		log.Fatalln(err.Error())
	}

	if err := head.InsertIntoBucket("mongo", "hello world"); err != nil {
		log.Fatalln(err.Error())
	}

	data, _ := head.GetAll("mongo")

	fmt.Println(data)
}
