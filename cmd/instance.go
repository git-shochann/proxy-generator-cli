package cmd

import (
	"log"
	"nhn-toast-api/internal"

	"github.com/spf13/cobra"
)

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
		newToken := token.Access.Token.ID
		// インスタンスの作成
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd) // インスタンスを作成するコマンド(createCmd)をルートコマンドに追加する
}
