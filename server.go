package main

import (
	"github.com/joho/godotenv"
	"github.com/yongyuth-chuankhuntod/bootstrap"
)

func main() {
	_ = godotenv.Load()
	bootstrap.RootApp.Execute()
}
