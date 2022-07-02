package main

import (
	"fmt"
	"log"
	config "nhn-toast-api/configs"
	"nhn-toast-api/internal"
	"os"
)

func main() {

	// トークンの取得
	fmt.Println("Getting Token...")
	t, err := internal.GetToken()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("done!")

	token := t.Access.Token.ID

	// イメージリストの取得
	if _, err := os.Stat("../image-list.txt"); err != nil {
		fmt.Println("Getting Image List...")
		_, err = internal.GetImageList(token)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("done!")
	}

	// インスタンスの作成
	fmt.Println("Generating Instance...")
	_, err = internal.Createinstance(token, config.Config.TenantID)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("done!")

	// floatingIP作成
	fmt.Println("Generating GlobalIP...")
	_, err = internal.CreateIP(token, config.Config.TenantID)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("done!")

}
