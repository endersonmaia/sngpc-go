package main

import (
	"fmt"
	"log"

	"github.com/endersonmaia/sngpc-go/sngpc"
)

func main() {

	mysngpcmov, err := sngpc.MovimentoFromXMLPath(`mensagemSNGPC.sample.xml`)
	if err == nil {
		fmt.Printf("\n%s", &mysngpcmov)
	} else {
		log.Panic(err)
	}

	mysngpcinv, err := sngpc.InventarioFromXMLPath(`mensagemSNGPCInventario.sample.xml`)
	if err == nil {
		fmt.Printf("\n%s", &mysngpcinv)
	} else {
		log.Panic(err)
	}
}
