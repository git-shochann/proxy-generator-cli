package temp

func temp() {

	// イメージリストの取得
	// if _, err := os.Stat("../image-list.txt"); err != nil {
	// 	fmt.Println("Getting Image List...")
	// 	_, err = internal.GetImageList(token)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Println("done!")
	// }

	// インスタンス詳細情報の取得

	// var GetInstanceInfoRes *internal.GetInstanceInfoRes

	// count := 0
	// // Retry up to 5 times
	// for count < 5 {
	// 	time.Sleep(time.Second * 10)
	// 	times := "Getting Server Detail" + "(" + strconv.Itoa(count) + ")"
	// 	fmt.Println(times)
	// 	instanceInfo, err := instance.GetInstanceInfo(token, tenantid)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	if instanceInfo.Server.Status == "BUILD" {
	// 		count += 1
	// 		continue
	// 	}
	// 	GetInstanceInfoRes = instanceInfo
	// 	break
	// }

	// fmt.Println(GetInstanceInfoRes)
	// fmt.Println("done!")
}
