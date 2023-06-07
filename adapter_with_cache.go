package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	name string
	age  int
}

type User interface {
	GetName() string
	GetAge() int
}

func userFactory() func(name string, age int) User {
	return func(name string, age int) User {
		u := user{}
		u.name = name
		u.age = age
		return u
	}
}

func (u user) GetName() string {
	return u.name
}

func (u user) GetAge() int {
	return u.age
}

type encripted_user struct {
	b [16]byte
}

var userCache = make(map[string][16]byte)

func showAnonimazedUser(user encripted_user) {
	fmt.Println("User data anonimized is:", user)
}

func userAdapter(u User) encripted_user {
	var userAnonimized [16]byte

	if v, ok := userCache[u.GetName()]; ok {
		userAnonimized = v
		return encripted_user{userAnonimized}
	}

	userMarsheled, err := json.Marshal(u)
	if err != nil {
		log.Fatal("Not able to marshal user data")
	}

	userAnonimized = md5.Sum(userMarsheled)
	fmt.Println("Storing data in cache")
	userCache[u.GetName()] = userAnonimized

	return encripted_user{userAnonimized}
}

func adpter() {
	fmt.Println("Adapter with cache")
	user := userFactory()
	u := user("username", 21)
	fmt.Println("User data", u.GetName(), u.GetAge())

	userAnonimized := userAdapter(u)
	userAnonimizedCacheTest := userAdapter(u)

	showAnonimazedUser(userAnonimized)
	showAnonimazedUser(userAnonimizedCacheTest)

}
