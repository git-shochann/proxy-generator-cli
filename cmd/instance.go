package cmd

import (
	"log"
	"nhn-toast-api/internal"

	"github.com/spf13/cobra"
)

var Token string
var TenantID string

// 構造体の初期化
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an instance and attach a global IP.",
	RunE: func(cmd *cobra.Command, args []string) error { // TODO: check
		// トークンの取得
		token, err := internal.GetToken()
		if err != nil {
			log.Fatalln(err)
		}
		Token = token.Access.Token.ID
		// TenantID =
		// instance, err = internal.Createinstance(Token)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
