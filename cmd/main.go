package main

import (
	"fmt"
	"log"
	config "nhn-toast-api/configs"
	"nhn-toast-api/internal"
)

func main() {

	// トークンの取得
	fmt.Println("Getting token...")
	t, err := internal.GetToken()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("done!")

	token := t.Access.Token.ID

	// イメージリストの取得
	// if _, err := os.Stat("../image-list.txt"); err != nil {
	// 	fmt.Println("Getting Image List...")
	// 	_, err = internal.GetImageList(token)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Println("done!")
	// }

	// // インスタンスの作成
	fmt.Println("Generating Instance...")
	instance, err := internal.Createinstance(token, config.Config.TenantID)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(instance)
	fmt.Println("done!")

	// インスタンス詳細情報の取得
	fmt.Println("Getting Server Detail...")
	instanceInfo, err := instance.GetInstanceInfo(token, config.Config.TenantID)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(instanceInfo)
	fmt.Println("done!")

	// floatingIP作成
	// fmt.Println("Generating floatingIP...")
	// floatingip, err := internal.CreateFloatingIP(token, config.Config.TenantID)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(floatingip)
	// fmt.Println("done!")

	// // floatingIP接続
	// fmt.Println("Connecting to instance...")
	// floatingip.ConnectingIP(token, instance.Server.ID) // fixed_ip 192~

}
