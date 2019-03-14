package main

import (
	"fmt"

	"github.com/endersonmaia/sngpc-go/sngpc"
)

func main() {

	mysngpcmov := sngpc.MovimentoFromXMLPath(`mensagemSNGPC.sample.xml`)
	fmt.Printf("\n%s", &mysngpcmov)

	mysngpcinv := sngpc.InventarioFromXMLPath(`mensagemSNGPCInventario.sample.xml`)
	fmt.Printf("\n%s", &mysngpcinv)

}
