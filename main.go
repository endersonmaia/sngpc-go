package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

//ClasseTerapeutica - st_classeTerapeutica
type ClasseTerapeutica uint8

//ClasseTerapeutica - st_classeTerapeutica enum
const (
	Antimicrobiano           ClasseTerapeutica = 1
	SujeitoAControleEspecial ClasseTerapeutica = 2
)

type UnidadeMedidaMedicamento uint8

const (
	Caixas  UnidadeMedidaMedicamento = 1
	Frascos UnidadeMedidaMedicamento = 2
)

type TipoUnidadeInsumo uint8

const (
	Grama     TipoUnidadeInsumo = 1
	Mililitro TipoUnidadeInsumo = 2
	Unidade   TipoUnidadeInsumo = 3
)

//Cabecalho XML
//    <cabecalho>
//         <cnpjEmissor>05059874000138</cnpjEmissor>
//         <cpfTransmissor>72586648153</cpfTransmissor>
//         <data>2006-09-30</data>
//     </cabecalho>
type Cabecalho struct {
	XMLName        xml.Name `xml:"cabecalho"`
	CnpjEmissor    string   `xml:"cnpjEmissor"`
	CpfTransmissor string   `xml:"cpfTransmissor"`
	Data           string   `xml:"data"`
}

//Insumos XML
type Insumos struct {
	XMLName xml.Name `xml:"insumos"`
}

//EntradaInsumos
type EntradaInsumos struct {
	XMLName       xml.Name        `xml:"entradaInsumos"`
	InsumoEntrada []InsumoEntrada `xml:"insumoEntrada"`
}

//MedicamentoEntrada
// <insumoEntrada>
// 		<classeTerapeutica>1</classeTerapeutica>
// 		<codigoInsumo>00092</codigoInsumo>
// 		<numeroLoteInsumo>A315</numeroLoteInsumo>
// 		<insumoCNPJFornecedor>99900099900000</insumoCNPJFornecedor>
// 		<quantidadeInsumo>300000.0</quantidadeInsumo>
// 		<tipoUnidade>1</tipoUnidade>
// </insumoEntrada>
type InsumoEntrada struct {
	XMLName              xml.Name          `xml:"InsumoEntrada"`
	ClasseTerapeutica    ClasseTerapeutica `xml:"classeTerapeutica"`
	CodigoInsumo         string            `xml:"codigoInsumo"`
	NumeroLoteInsumo     string            `xml:"numeroLoteInsumo"`
	InsumoCNPJFornecedor string            `xml:"insumoCNPJFornecedor"`
	QuantidadeInsumo     int               `xml:"quantidadeInsumo"`
	UnidadeMedidaInsumo  TipoUnidadeInsumo `xml:"unidadeMedidaInsumo"`
}

//Medicamentos XML
type Medicamentos struct {
	XMLName xml.Name `xml:"medicamentos"`
}

//EntradaMedicamentos
type EntradaMedicamentos struct {
	XMLName            xml.Name             `xml:"entradamedicamentos"`
	MedicamentoEntrada []MedicamentoEntrada `xml:"medicamentoEntrada"`
}

//MedicamentoEntrada
// <medicamentoEntrada>
// <classeTerapeutica>1</classeTerapeutica>
// <registroMSMedicamento>1888888888888</registroMSMedicamento>
// <numeroLoteMedicamento>200678</numeroLoteMedicamento>
// <quantidadeMedicamento>1234</quantidadeMedicamento>
// <unidadeMedidaMedicamento>1</unidadeMedidaMedicamento>
// </medicamentoEntrada>
type MedicamentoEntrada struct {
	XMLName                  xml.Name                 `xml:"medicamentoEntrada"`
	ClasseTerapeutica        ClasseTerapeutica        `xml:"classeTerapeutica"`
	RegistroMSMedicamento    string                   `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string                   `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    int                      `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento UnidadeMedidaMedicamento `xml:"unidadeMedidaMedicamento"`
}

//Corpo XML
type Corpo struct {
	XMLName      xml.Name       `xmml:"corpo"`
	Medicamentos []Medicamentos `xml:"medicamentos"`
	Insumos      []Insumos      `xml:"insumos"`
}

//MensagemSNGPCInventario root XML
type MensagemSNGPCInventario struct {
	XMLName   xml.Name  `xml:"mensagemSNGPCInventario"`
	Cabecalho Cabecalho `xml:"cabecalho"`
	Corpo     Corpo     `xml:"corpo"`
}

func (s MensagemSNGPCInventario) String() string {
	return fmt.Sprintf("%v\n%v", s.Cabecalho, s.Corpo)
}

func (s Cabecalho) String() string {
	return fmt.Sprintf("cnpj : %v ; cpf : %v ; data : %v\n", s.CnpjEmissor, s.CpfTransmissor, s.Data)
}

func (s Corpo) String() string {
	return fmt.Sprintf("medicamentos : %v ; insumos : %v\n", s.Medicamentos, s.Insumos)
}

func main() {
	var sngpc MensagemSNGPCInventario

	xmlFile, err := os.Open("inventario.xml")

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	XML, _ := ioutil.ReadAll(xmlFile)

	var re = regexp.MustCompile(`^.*(ISO-8859-1).*$`)
	s := re.ReplaceAllString(string(XML), `asdasdasd`)
	fmt.Println(s)
	xml.Unmarshal([]byte(s), &sngpc)
	fmt.Println(sngpc)
}
