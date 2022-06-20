package main

import (
	"fmt"
	"log"
	"nhn-toast-api/internal"
	"os"
)

func main() {

	// トークンの取得
	fmt.Println("Getting Token")
	t, err := internal.GetToken()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("done!")

	// イメージリストの取得
	if _, err := os.Stat("../image-list.txt"); err != nil {
		fmt.Println("Getting Image List")
		_, err = internal.GetImageList(t.Access.Token.ID)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("done!")
	}
}
