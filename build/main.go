package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/tyfura/gwconn"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type connInfo struct {
	host  string
	login string
	key   string
	token string
}

func callLogin(cI *connInfo) {
	//	fmt.Println(fmt.Sprintf("%s:%v", cI.host, gwconn.GW_RpcPort))
	cc, err := grpc.NewClient(
		fmt.Sprintf("%s:%v", cI.host, gwconn.GW_RpcPort),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
	)

	// Link gw (domain:port)

	if err != nil {
		log.Fatalf("error dialing gw: %v", err)
	}
	defer cc.Close()

	cli := gwconn.NewGWClient(cc)
	resp, err := cli.Login(context.TODO(), &gwconn.LoginReq{Login: cI.login, Key: cI.key})
	if err != nil {
		log.Fatalf("error response when dialing gw login: %v", err)
	}

	cI.token = resp.Token
}

func loginFn(cI *connInfo) {
	fmt.Print("Login: ")
	fmt.Scanln(&cI.login)

	fmt.Print("Key: \033[8m") // Hide input
	fmt.Scan(&cI.key)
	fmt.Println("\033[28m") // Show input

	callLogin(cI)
}

const evansHeaderOption = "--header"
const tgwAuthHeader = "tgwauth="
const tgwProtoFile = "./proto/gw.proto"

func runEvans(cI *connInfo) {
	var tgwOpts []string = []string{"-t", "-p", gwconn.GW_RpcPort, "--host"}
	tgwOpts = append(tgwOpts, cI.host)
	tgwOpts = append(tgwOpts, evansHeaderOption)
	tgwOpts = append(tgwOpts, tgwAuthHeader+cI.token)
	tgwOpts = append(tgwOpts, tgwProtoFile)

	cmd := exec.Command("evans", tgwOpts...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run() // add error checking
	//fmt.Println(cmd)

}

func main() {
	cI := &connInfo{}
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("missing gw address")
	}
	cI.host = args[0]
	loginFn(cI)
	runEvans(cI)
}
