package sngpc

import (
	"encoding/xml"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

// SNGPCXMLVersion armazena a versão do XML do SNGPC
// Ver: http://portal.anvisa.gov.br/sngpc/desenvolvedores
const SNGPCXMLVersion = "2.0"

type TipoMensagemSNGPC uint

const (
	Movimentacao TipoMensagemSNGPC = iota + 1
	Inventario
)

// InventarioFromXMLPath lê o conteúdo de um XML SNGPC do tipo `MensagemSNGPCInventario`
func InventarioFromXMLPath(path string) (msg MensagemSNGPCInventario, err error) {
	msg = MensagemSNGPCInventario{}
	err = loadFromXMLPath(path, &msg)

	return
}

// MovimentoFromXMLPath lê o conteúdo de um XML SNGPC do tipo `MensagemSNGPC`
func MovimentoFromXMLPath(path string) (msg MensagemSNGPC, err error) {
	msg = MensagemSNGPC{}
	err = loadFromXMLPath(path, &msg)

	return
}

func loadFromXMLPath(path string, o interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return fromXML(f, &o)
}

func fromXML(r io.Reader, o interface{}) error {
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReaderLabel
	return dec.Decode(o)
}
