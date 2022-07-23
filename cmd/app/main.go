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

	user := &User{
		Username: "andrii",
	}

	b, _ := json.Marshal(user)

	var res interface{}

	head.Insert("v1", b)
	head.Get("v1", &res)

	fmt.Println(res)

	user2 := &User{
		Username: "olek",
	}

	b2, _ := json.Marshal(user2)

	head.Update("v1", b2)

	var res2 interface{}
	head.Get("v1", &res2)

	head.Delete("v1")
	fmt.Println(res2)

	var res3 interface{}
	head.Get("v1", &res3)
	fmt.Println(res3)
}
