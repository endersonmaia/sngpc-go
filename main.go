package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/endersonmaia/sngpc/parser"
	"golang.org/x/net/html/charset"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	xmlFile, err := os.Open("mensagemSNGPCInventario.sample.xml")
	check(err)

	dec := xml.NewDecoder(xmlFile)
	dec.CharsetReader = charset.NewReaderLabel

	defer xmlFile.Close()

	_, err = xmlFile.Seek(40, 0)
	check(err)

	reader := bufio.NewReader(xmlFile)
	buffer, err := reader.Peek(100)
	check(err)

	if strings.Contains(string(buffer), `<mensagemSNGPCInventario xmlns="urn:sngpc-schema">`) {

		_, err = xmlFile.Seek(0, 0)
		check(err)

		mysngpc := parser.MensagemSNGPCInventario{}

		err = dec.Decode(&mysngpc)
		check(err)

		fmt.Printf("\n%s", mysngpc)
	} else if strings.Contains(string(buffer), `<mensagemSNGPC xmlns="urn:sngpc-schema">`) {
		_, err = xmlFile.Seek(0, 0)
		check(err)

		mysngpc := parser.MensagemSNGPC{}

		err = dec.Decode(&mysngpc)
		check(err)

		fmt.Printf("\n%s", mysngpc)
	}

}
