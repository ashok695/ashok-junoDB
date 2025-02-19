package main

import (
	"fmt"

	"github.com/paypal/junodb/pkg/client"
	"github.com/paypal/junodb/pkg/io"
)

func main() {
	fmt.Println("hello")
	ClientConfig := client.Config{
		Appname:   "ashok",
		Namespace: "ashok-junodb",
		Server: io.ServiceEndpoint{
			Addr: "192.168.1.6:8087",
		},
	}
	cl, err := client.New(ClientConfig)
	if err != nil {
		fmt.Print("error in connecting the JUNO DB")
	}
	cl.Get("key_name")
	cl.Set("key-name", "value")
	fmt.Print(cl)
	fmt.Println("JUNO DB SUCCESS")
}
