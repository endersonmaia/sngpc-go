package sngpc

import (
	"encoding/xml"
	"os"

	"golang.org/x/net/html/charset"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// http://portal.anvisa.gov.br/sngpc/desenvolvedores

const SNGPCXMLVersion = "2.0"

type TipoMensagemSNGPC uint8

const (
	Inventario TipoMensagemSNGPC = iota
	Movimentacao
)

func InventarioFromXMLPath(f string) MensagemSNGPCInventario {

	xmlFile, err := os.Open(f)
	check(err)

	dec := xml.NewDecoder(xmlFile)
	dec.CharsetReader = charset.NewReaderLabel

	defer xmlFile.Close()

	mysngpc := MensagemSNGPCInventario{}

	err = dec.Decode(&mysngpc)
	check(err)

	return mysngpc
}

func MovimentoFromXMLPath(f string) MensagemSNGPC {

	xmlFile, err := os.Open(f)
	check(err)

	dec := xml.NewDecoder(xmlFile)
	dec.CharsetReader = charset.NewReaderLabel

	defer xmlFile.Close()

	mysngpc := MensagemSNGPC{}

	err = dec.Decode(&mysngpc)
	check(err)

	return mysngpc
}
