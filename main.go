package main

import (
	"fmt"
	"nhn-toast-api/api"
	"os"
)

func main() {
	token := api.GetToken()
	fmt.Println(token.Access.Token.ID) // データが格納されている構造体にアクセス
	instance, err := api.CreateInstance(token.Access.Token.ID, token.Access.Token.Tenant.ID)
	fmt.Println(instance) // ここに500のメッセージが入ってる
	fmt.Println(err)      // nil？x

	if err != nil {
		fmt.Printf("fail to generate server: %v", err)
		os.Exit(1)
	}

	fmt.Println(instance)

	// str := util.RandomGenerate(5)
	// fmt.Println(str)
}
