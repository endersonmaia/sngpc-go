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

func InventarioFromXMLPath(f string) (MensagemSNGPCInventario, error) {

	xmlFile, err := os.Open(f)
	if err != nil {
		return MensagemSNGPCInventario{}, err
	}

	dec := xml.NewDecoder(xmlFile)
	dec.CharsetReader = charset.NewReaderLabel
	defer xmlFile.Close()

	msg := MensagemSNGPCInventario{}

	if err = dec.Decode(&msg); err != nil {
		return MensagemSNGPCInventario{}, err
	}

	return msg, err
}

func MovimentoFromXMLPath(f string) (MensagemSNGPC, error) {

	xmlFile, err := os.Open(f)
	if err != nil {
		return MensagemSNGPC{}, err
	}

	dec := xml.NewDecoder(xmlFile)
	dec.CharsetReader = charset.NewReaderLabel
	defer xmlFile.Close()

	msg := MensagemSNGPC{}

	if err = dec.Decode(&msg); err != nil {
		return MensagemSNGPC{}, err
	}

	return msg, err
}
