package cmd

import (
	"fmt"
	"log"
	config "nhn-toast-api/configs"
	"nhn-toast-api/internal"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

var tenantid = config.Config.TenantID

// インスタンスを作成
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an instance, attach an IP, and build a Proxy server",
	RunE: func(cmd *cobra.Command, args []string) error {
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

		// ポートID取得
		var portinfo *internal.GetPortListRes
		count := 1
		for count < 5 {
			time.Sleep(time.Second * 10)
			times := "Getting Port List..." + "(" + strconv.Itoa(count) + ")"
			fmt.Println(times)
			port, err := internal.GetPortList(token, instance)
			if err != nil {
				log.Fatalln(err)
			}
			if len(port.Ports) == 0 {
				count += 1
				continue
			}
			portinfo = port
			break
		}

		// IPをインスタンスに接続
		fmt.Println("Connecting to instance...")
		connectedInstance, err := internal.ConnectingIP(token, floatingip, portinfo)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Done!")
		fmt.Println(connectedInstance)

		// ssh接続を行いシェルスクリプトの実行
		ip := connectedInstance.Floatingip.FloatingIPAddress
		port := "22"
		user := "centos"

		// buf, err := ioutil.ReadFile("")

		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh
			}
		}

		return nil

	},
}

func init() {
	rootCmd.AddCommand(createCmd) // インスタンスを作成するコマンド(createCmd)をルートコマンドに追加する
}
