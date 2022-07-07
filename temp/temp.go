package temp

// import (
// 	"fmt"
// 	"log"
// 	config "nhn-toast-api/configs"
// 	"nhn-toast-api/internal"
// 	"strconv"
// 	"time"
// )

// func main() {

// 	// トークンの取得
// 	fmt.Println("Getting token...")
// 	t, err := internal.GetToken()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println("done!")

// 	token := t.Access.Token.ID

// 	tenantid := config.Config.TenantID

// 	// イメージリストの取得
// 	// if _, err := os.Stat("../image-list.txt"); err != nil {
// 	// 	fmt.Println("Getting Image List...")
// 	// 	_, err = internal.GetImageList(token)
// 	// 	if err != nil {
// 	// 		log.Fatalln(err)
// 	// 	}
// 	// 	fmt.Println("done!")
// 	// }

// 	// // インスタンスの作成
// 	fmt.Println("Generating Instance...")
// 	instance, err := internal.Createinstance(token, tenantid)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(instance)
// 	fmt.Println("done!")

// 	// インスタンス詳細情報の取得

// 	// var GetInstanceInfoRes *internal.GetInstanceInfoRes

// 	// count := 0
// 	// // Retry up to 5 times
// 	// for count < 5 {
// 	// 	time.Sleep(time.Second * 10)
// 	// 	times := "Getting Server Detail" + "(" + strconv.Itoa(count) + ")"
// 	// 	fmt.Println(times)
// 	// 	instanceInfo, err := instance.GetInstanceInfo(token, tenantid)
// 	// 	if err != nil {
// 	// 		log.Fatalln(err)
// 	// 	}
// 	// 	if instanceInfo.Server.Status == "BUILD" {
// 	// 		count += 1
// 	// 		continue
// 	// 	}
// 	// 	GetInstanceInfoRes = instanceInfo
// 	// 	break
// 	// }

// 	// fmt.Println(GetInstanceInfoRes)
// 	// fmt.Println("done!")

// 	// floatingIP作成
// 	fmt.Println("Generating floatingIP...")
// 	floatingip, err := internal.CreateFloatingIP(token, tenantid)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(floatingip)
// 	fmt.Println("done!")

// 	// ポートID取得
// 	var GetPortListRes *internal.GetPortListRes

// 	count := 0
// 	// Retry up to 5 times
// 	for count < 5 {
// 		time.Sleep(time.Second * 10)
// 		times := "Getting Port List..." + "(" + strconv.Itoa(count) + ")"
// 		fmt.Println(times)
// 		portinfo, err := instance.GetPortList(token)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		if len(portinfo.Ports) == 0 {
// 			count += 1
// 			continue
// 		}
// 		GetPortListRes = portinfo
// 		break
// 	}
// 	fmt.Println("done!")

// 	portid := GetPortListRes.Ports[0].ID

// 	//floatingIP接続 + ポートID取得も含む
// 	fmt.Println("Connecting to instance...")
// 	connectingIPRes, err := floatingip.ConnectingIP(token, portid)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(connectingIPRes)
// 	fmt.Println("done!")

// }
