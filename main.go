package main

import (
	"fmt"
	"nhn-toast-api/api"
)

// TODO: 質問する -> main関数は以下のように関数呼び出して戻り値が2つあるものを呼ぶのは定石ではないか？ main関数以外でfunction作って、そこで呼び出すべきか？ -> なぜかというとエラーが返ってきたとき
// main() token, err とはしないと思うので。
func main() {

	token, err := api.GetToken()

	if err != nil {
		return, nil
	}

	fmt.Println(token.Access.Token.ID) // データが格納されている構造体にアクセス

	// instance, err := api.CreateInstance(token.Access.Token.ID, token.Access.Token.Tenant.ID)
	// fmt.Println(instance) // ここに500のメッセージが入ってる
	// fmt.Println(err)      // nil？x

	// if err != nil {
	// 	fmt.Printf("fail to generate server: %v", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(instance)

	// str := util.RandomGenerate(5)
	// fmt.Println(str)
}