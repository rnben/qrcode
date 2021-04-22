package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/rnben/qrcode/handler"
	"github.com/rnben/qrcode/server"
)

var (
	User string
	Pwd  string
	Port string
)

func init() {
	flag.StringVar(&User, "wifi_user", "", "wifi username")
	flag.StringVar(&Pwd, "wifi_pwd", "", "wifi password")
	flag.StringVar(&Port, "port", "", "server port")
	flag.Parse()

	fmt.Println(User, Pwd)

	if strings.TrimSpace(User) == "" {
		flag.Usage()
		os.Exit(1)
	}

	if strings.TrimSpace(Pwd) == "" {
		flag.Usage()
		os.Exit(1)
	}

	handler.WIFI = fmt.Sprintf("WIFI:T:WPA;S:%s;P:%s;;", User, Pwd+"%s")
}

func main() {
	srv := server.NewServer(server.WithTimeout(5))
	srv.WithAddr(Port)

	http.HandleFunc("/", handler.GenerateQRHandler)

	log.Fatal(srv.ListenAndServe())
}
