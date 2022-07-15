package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// ルートコマンド
// Runが定義されていない状態だと、ヘルプのメッセージが表示される
var rootCmd = &cobra.Command{
	Use:     "nhn-toast-api",
	Version: "0.0.1",
	Short:   `アスキーアートと簡単な説明を入れる`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// それ以外のコマンドは別ファイルに切り出して、init関数内でルートコマンドに対してサブコマンドとして追加する
// Ex) go run main.go --create
