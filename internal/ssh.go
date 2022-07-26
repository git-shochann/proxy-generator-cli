package internal

import (
	"fmt"
	"log"
	"net"
	"os"

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

	// ssh接続の実行
	connection, err := ssh.Dial("tcp", net.JoinHostPort(ip, port), config)
	if err != nil {
		log.Fatalln(err)
	}

	defer connection.Close()

	fmt.Println("done!")
	fmt.Println(connection)

	// ここで処理が終わっている

	// セッションを開く
	session, err := connection.NewSession()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(session)

	defer session.Close()

	// リモートサーバーのコマンド実行結果をローカルの標準出力と標準エラーへと渡す

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	fmt.Println("start script")

	// シェルスクリプトの実行
	if err = session.Run("echo hello world"); err != nil {
		log.Fatalln(err)
	}

	// 実際の値を返す 型ではない。
	return session, nil

}
