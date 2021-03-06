package main

import (
	"github.com/mattn/go-uv"
	"log"
)

func main() {
	tcp, _ := uv.TcpInit(nil)
	addr, err := uv.Ip4Addr("0.0.0.0", 8888)
	if err != nil {
		log.Fatal(err)
	}
	err = tcp.Bind(addr)
	if err != nil {
		log.Fatal(err)
	}
	tcp.Listen(10, func(h *uv.Handle, status int) {
		client, _ := tcp.Accept()
		log.Println("server: accept")
		line := ""
		client.ReadStart(func(h *uv.Handle, data []byte) {
			if data == nil {
				log.Println("client: closed")
				client.Close(nil)
				return
			}
			s := string(data)
			print(s)
			line += s
			if s[len(s)-1] == '\n' {
				client.Write([]byte(line), func(r *uv.Request, status int) {
					log.Println("client: written")
				})
			}
		})
	})

	uv.DefaultLoop().Run()
}
