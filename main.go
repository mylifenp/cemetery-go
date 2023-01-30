package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.Run()
}

type User struct {
	firstName string `json:"name"`
	lastName  string `json:"last_name"`
	email     string `json:"email"`
	password  string `json:"password"`
}
