package main

import "fmt"

func main()  {
	// 初期化、このあたりをコンストラクタって言ってる気がする
	repo := NewUserRepository()
	service := NewUserService(repo)
	fmt.Println(service)
}

