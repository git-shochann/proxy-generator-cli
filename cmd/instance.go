package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	config "nhn-toast-api/configs"
	"nhn-toast-api/internal"
	"time"

	"github.com/spf13/cobra"
)

var tenantid = config.Config.TenantID

// インスタンスを作成
var createCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "create",
	Short: "Create an instance, attach an IP, and build a Proxy server",
	RunE: func(cmd *cobra.Command, args []string) error {

		var pemFile []byte

		for _, v := range args {
			file, err := ioutil.ReadFile(v)
			if err != nil {
				log.Fatalln(err)
			}

			pemFile = file
		}

		// トークンの取得
		token, err := internal.GetToken()
		if err != nil {
			log.Fatalln(err)
		}

		// インスタンスの作成
		instance, err := internal.CreateInstance(token, tenantid)
		if err != nil {
			log.Fatalln(err)
		}

		// floatingIP作成
		fmt.Println("Generating FloatingIP...")
		floatingip, err := internal.CreateFloatingIP(token, tenantid)
		if err != nil {
			log.Fatalln(err)
		}

		time.Sleep(10 * time.Second)

		// ポートID取得
		// var portinfo *internal.GetPortListRes
		// count := 1
		// for count < 5 {
		// 	time.Sleep(time.Second * 10)
		// 	times := "Getting Port List..." + "(" + strconv.Itoa(count) + ")"
		fmt.Println("Getting Port List...")
		// 	fmt.Println(times)
		portlist, err := internal.GetPortList(token, instance)
		if err != nil {
			log.Fatalln(err)
		}
		// if len(port.Ports) == 0 {
		// 	count += 1
		// 	continue
		// }
		// 	portinfo = port
		// 	break
		// }

		// IPをインスタンスに接続
		fmt.Println("Connecting to instance...")
		connectedInstance, err := internal.ConnectingIP(token, floatingip, portlist)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Done!")
		fmt.Printf("%+v", connectedInstance)

		// ここでStatusのチェック処理を入れる -> ACTIVEであれば処理を続ける 10回まで繰り返す + 待機処理を入れる
		for i := 0; i < 10; i++ {
			status, err := internal.CheckIPStatus(token, floatingip)
			if err != nil {
				log.Fatalln(err)
			}
			if status != internal.Active {
				fmt.Println(status)
				fmt.Println("リトライします")
				time.Sleep(20 * time.Second)
				continue
			}
			break
		}

		// ssh接続を行いシェルスクリプトの実行
		ip := connectedInstance.Floatingip.FloatingIPAddress
		port := "22"
		user := "centos"

		session, err := internal.SSHwithPublicKeyAuthentication(ip, port, user, pemFile)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(session)

		return nil

	},
}

func init() {
	rootCmd.AddCommand(createCmd) // インスタンスを作成するコマンド(createCmd)をルートコマンドに追加する
}
