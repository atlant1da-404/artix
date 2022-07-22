package main

import (
	"encoding/json"
	"fmt"
	"github.com/atlant1da-404/artik_db/pkg/authenticate"
	"github.com/atlant1da-404/artik_db/pkg/bucket"
	"github.com/atlant1da-404/artik_db/pkg/config"
	"log"
)

type User struct {
	Username string `json:"username"`
}

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

	user := &User{
		Username: "rocket",
	}

	b, _ := json.Marshal(user)

	if err := head.InsertIntoBucket("mongo", b); err != nil {
		log.Fatalln(err.Error())
	}

	var getUser User

	if err := head.GetAllFromBucket("mongo", &getUser); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(getUser)
}
