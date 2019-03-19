package main

import (
	"fmt"
	"log"

	"github.com/endersonmaia/sngpc-go/sngpc"
)

func main() {
	mysngpcmov, err := sngpc.MovimentoFromXMLPath(`mensagemSNGPC.sample.xml`)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("\n%s", &mysngpcmov)

	mysngpcinv, err := sngpc.InventarioFromXMLPath(`mensagemSNGPCInventario.sample.xml`)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("\n%s", &mysngpcinv)
}
