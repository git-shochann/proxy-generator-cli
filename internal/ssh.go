package internal

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func SSHwithPublicKeyAuthentication(ip, port, user string, privateKey []byte) (*ssh.Session, error) {

	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		log.Fatalln(err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// インスタンスに接続するまで時間がかかりそうなので、リトライ処理を入れる
	fmt.Println("Wait 60 seconds...")
	time.Sleep(60 * time.Second)

	// ssh接続の実行
	connection, err := ssh.Dial("tcp", net.JoinHostPort(ip, port), config)
	if err != nil {
		log.Fatalln(err)
	}

	defer connection.Close()

	fmt.Println("done!")
	fmt.Println(connection)

	// セッションを開く
	session, err := connection.NewSession()
	if err != nil {
		log.Fatalln(err)
	}

	defer session.Close()

	// リモートサーバーのコマンド実行結果をローカルの標準出力と標準エラーへと渡す

	session.Stdout = os.Stdout // 出力
	session.Stderr = os.Stderr // エラー

	// standard in(標準入力)
	session.Run("/bin.sh")
	session.Run("echo hello world")

	fmt.Println("---")
	fmt.Println(session.Stdout)
	fmt.Println("---")

	// 実際の値を返す 型ではない。
	return session, nil

}
