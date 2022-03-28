package main

import (
	"fmt"
	"nhn-toast-api/api"
)

func main() {
	token := api.GetToken()
	fmt.Println(token.Access.Token.ID) // データが格納されている構造体にアクセス
}
