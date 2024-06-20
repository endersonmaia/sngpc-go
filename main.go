package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/endersonmaia/sngpc-go/sngpc"
	flag "github.com/spf13/pflag"
)

var (
	isInventario bool
	arquivoFlag  string
	toCSV        bool
	anonymize    bool
)

func init() {
	flag.BoolVarP(&isInventario, "inventario", "i", false, "arquivo SNGPC de inventário")
	flag.StringVar(&arquivoFlag, "arquivo", "", "caminho para um arquivo XML (Ex.: --from sngpc.xml)")
	flag.BoolVarP(&toCSV, "csv", "", false, "formato de exportaçao")
	flag.BoolVarP(&anonymize, "anonymise", "", false, "anonimizar os dados exportados")
}

func main() {
	flag.Parse()

	if arquivoFlag == "" {
		log.Fatalln("É preciso informar o arquivo XML (ex.: --arquivo sngpc.xml)")
	}

	if isInventario {
		m := sngpc.MensagemSNGPCInventario{}
		m, err := sngpc.InventarioFromXMLPath(arquivoFlag)
		if err != nil {
			log.Panic(err)
		}
		fmt.Print(&m)
	} else {
		m := sngpc.MensagemSNGPC{}
		m, err := sngpc.MovimentoFromXMLPath(arquivoFlag)
		if err != nil {
			log.Panic(err)
		}

		if toCSV {
			writer := csv.NewWriter(os.Stdout)
			if anonymize {
				m.VendasToCSVAnonymized(writer)
			} else {
				m.VendasToCSV(writer)
			}
		} else {
			fmt.Print(&m)
		}
	}
}
