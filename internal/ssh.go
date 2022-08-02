package internal

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func SSHwithPublicKeyAuthentication(ip, port, user string, privateKey []byte) (*ssh.Session, error) {

	shellscript, err := os.Open("proxy.sh")
	if err != nil {
		log.Fatalln("Unable to read script: ", err)
	}
	defer shellscript.Close()

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
	fmt.Println("Wait 80 seconds...")
	time.Sleep(80 * time.Second)

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

	var b bytes.Buffer
	session.Stdout = &b
	session.Stdin = shellscript
	if err := session.Run("/usr/bin/sh"); err != nil {
		log.Fatalf("Failed to run:%v", err)
	}
	fmt.Println("---")
	fmt.Println(b.String())
	fmt.Println(session.Stderr)
	fmt.Println("---")

	// 実際の値を返す 型ではない。
	return session, nil

}
