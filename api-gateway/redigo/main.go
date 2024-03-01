package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type User struct {
	Id       string
	Name     string
	LastName string
	Age      int64
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	user := User{
		Id:       "1",
		Name:     "Test Name",
		LastName: "Test Last Name",
		Age:      20,
	}

	byteData, err := json.Marshal(&user)
	if err != nil {
		log.Fatal(err)
		return
	}

	// set - create
	err = rdb.Set(context.Background(), user.Id, byteData, 0).Err()
	if err != nil {
		log.Fatal(err)
		return
	}

	// get - get
	respUser, err := rdb.Get(context.Background(), user.Id).Result()
	if err != nil {
		log.Fatal(err)
		return
	}
	var newUser User
	if err := json.Unmarshal([]byte(respUser), &newUser); err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println(newUser)
	}

	// del - delete
	res := rdb.Del(context.Background(), user.Id)
	
	fmt.Println(res)
	fmt.Println(rdb.Get(context.Background(), user.Id))

}
