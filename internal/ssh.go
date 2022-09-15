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
		log.Println("Unable to read script: ", err)
	}
	defer shellscript.Close()

	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}

	defer connection.Close()

	fmt.Println("done!")

	// セッションを開く
	session, err := connection.NewSession()
	if err != nil {
		log.Println(err)
	}

	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	session.Stdin = shellscript
	if err := session.Run("/usr/bin/sh"); err != nil {
		fmt.Printf("Failed to run:%v", err)
	}

	fmt.Printf("session.Stdout: %v\n", session.Stdout)
	fmt.Printf("session.Stdin: %v\n", session.Stdin)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("session.Stderr: %v\n", session.Stderr)

	// 実際の値を返す 型ではない。
	return session, nil

}
