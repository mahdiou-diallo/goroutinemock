package main

import (
	"log"

	"github.com/mahdiou-diallo/goroutinemock/core"
)

func main() {
	adder := core.NewRunner(&core.MyEchoer{})
	log.Default().Printf("sync : %s", adder.EchoSync("hello"))
	log.Default().Printf("async : %s", adder.EchoAsync("hello async"))
}
