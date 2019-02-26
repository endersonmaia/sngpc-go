package main

import (
	"encoding/xml"
	"fmt"
)

var XML = []byte(`
<mensagemSNGPCInventario xmlns="urn:sngpc-schema">
    <cabecalho>
        <cnpjEmissor>05059874000138</cnpjEmissor>
        <cpfTransmissor>72586648153</cpfTransmissor>
        <data>2006-09-30</data>
    </cabecalho>
    <corpo>
        <medicamentos>
            <entradaMedicamentos>
                <medicamentoEntrada>
                    <classeTerapeutica>1</classeTerapeutica>
                    <registroMSMedicamento>1888888888888</registroMSMedicamento>
                    <numeroLoteMedicamento>200678</numeroLoteMedicamento>
                    <quantidadeMedicamento>1234</quantidadeMedicamento>
                    <unidadeMedidaMedicamento>1</unidadeMedidaMedicamento>
                </medicamentoEntrada>
            </entradaMedicamentos>
            <entradaMedicamentos>
                <medicamentoEntrada>
                    <classeTerapeutica>1</classeTerapeutica>
                    <registroMSMedicamento>1888888888888</registroMSMedicamento>
                    <numeroLoteMedicamento>200678</numeroLoteMedicamento>
                    <quantidadeMedicamento>1234</quantidadeMedicamento>
                    <unidadeMedidaMedicamento>1</unidadeMedidaMedicamento>
                </medicamentoEntrada>
            </entradaMedicamentos>
        </medicamentos>
        <insumos>
            <entradaInsumos>
                <insumoEntrada>
                    <classeTerapeutica>1</classeTerapeutica>
                    <codigoInsumo>00092</codigoInsumo>
                    <numeroLoteInsumo>A315</numeroLoteInsumo>
                    <insumoCNPJFornecedor>99900099900000</insumoCNPJFornecedor>
                    <quantidadeInsumo>300000.0</quantidadeInsumo>
                    <tipoUnidade>1</tipoUnidade>
                </insumoEntrada>
            </entradaInsumos>
            <entradaInsumos>
                <insumoEntrada>
                    <classeTerapeutica>1</classeTerapeutica>
                    <codigoInsumo>00092</codigoInsumo>
                    <numeroLoteInsumo>A315</numeroLoteInsumo>
                    <insumoCNPJFornecedor>99900099900000</insumoCNPJFornecedor>
                    <quantidadeInsumo>300000.0</quantidadeInsumo>
                    <tipoUnidade>1</tipoUnidade>
                </insumoEntrada>
            </entradaInsumos>
        </insumos>
    </corpo>
</mensagemSNGPCInventario>
`)

//ClasseTerapeutica - st_classeTerapeutica
type SNGPCClasseTerapeutica uint8

//ClasseTerapeutica - st_classeTerapeutica enum
const (
	Antimicrobiano           SNGPCClasseTerapeutica = 1
	SujeitoAControleEspecial SNGPCClasseTerapeutica = 2
)

type SNGPCUnidadeMedidaMedicamento uint8

const (
	Caixas  SNGPCUnidadeMedidaMedicamento = 1
	Frascos SNGPCUnidadeMedidaMedicamento = 2
)

type SNGPCTipoUnidadeInsumo uint8

const (
	Grama     SNGPCTipoUnidadeInsumo = 1
	Mililitro SNGPCTipoUnidadeInsumo = 2
	Unidade   SNGPCTipoUnidadeInsumo = 3
)

//SNGPCCabecalho XML
//    <cabecalho>
//         <cnpjEmissor>05059874000138</cnpjEmissor>
//         <cpfTransmissor>72586648153</cpfTransmissor>
//         <data>2006-09-30</data>
//     </cabecalho>
type SNGPCCabecalho struct {
	XMLName        xml.Name `xml:"cabecalho"`
	CnpjEmissor    string   `xml:"cnpjEmissor"`
	CpfTransmissor string   `xml:"cpfTransmissor"`
	Data           string   `xml:"data"`
}

//SNGPCInsumos XML
type SNGPCInsumos struct {
	XMLName xml.Name `xml:"insumos"`
}

//SNGPCEntradaInsumos
type SNGPCEntradaInsumos struct {
	XMLName       xml.Name              `xml:"entradaInsumos"`
	InsumoEntrada []SNGPCInsumoEnrtrada `xml:"insumoEntrada"`
}

//SNGPCMedicamentoEntrada
// <insumoEntrada>
// 		<classeTerapeutica>1</classeTerapeutica>
// 		<codigoInsumo>00092</codigoInsumo>
// 		<numeroLoteInsumo>A315</numeroLoteInsumo>
// 		<insumoCNPJFornecedor>99900099900000</insumoCNPJFornecedor>
// 		<quantidadeInsumo>300000.0</quantidadeInsumo>
// 		<tipoUnidade>1</tipoUnidade>
// </insumoEntrada>
type SNGPCInsumoEnrtrada struct {
	XMLName              xml.Name               `xml:"InsumoEntrada"`
	ClasseTerapeutica    SNGPCClasseTerapeutica `xml:"classeTerapeutica"`
	CodigoInsumo         string                 `xml:"codigoInsumo"`
	NumeroLoteInsumo     string                 `xml:"numeroLoteInsumo"`
	InsumoCNPJFornecedor string                 `xml:"insumoCNPJFornecedor"`
	QuantidadeInsumo     int                    `xml:"quantidadeInsumo"`
	UnidadeMedidaInsumo  SNGPCTipoUnidadeInsumo `xml:"unidadeMedidaInsumo"`
}

//SNGPCMedicamentos XML
type SNGPCMedicamentos struct {
	XMLName xml.Name `xml:"medicamentos"`
}

//SNGPCEntradaMedicamentos
type SNGPCEntradaMedicamentos struct {
	XMLName            xml.Name                   `xml:"entradamedicamentos"`
	MedicamentoEntrada []SNGPCMedicamentoEnrtrada `xml:"medicamentoEntrada"`
}

//SNGPCMedicamentoEntrada
// <medicamentoEntrada>
// <classeTerapeutica>1</classeTerapeutica>
// <registroMSMedicamento>1888888888888</registroMSMedicamento>
// <numeroLoteMedicamento>200678</numeroLoteMedicamento>
// <quantidadeMedicamento>1234</quantidadeMedicamento>
// <unidadeMedidaMedicamento>1</unidadeMedidaMedicamento>
// </medicamentoEntrada>
type SNGPCMedicamentoEnrtrada struct {
	XMLName                  xml.Name                      `xml:"medicamentoEntrada"`
	ClasseTerapeutica        SNGPCClasseTerapeutica        `xml:"classeTerapeutica"`
	RegistroMSMedicamento    string                        `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string                        `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    int                           `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento SNGPCUnidadeMedidaMedicamento `xml:"unidadeMedidaMedicamento"`
}

//SNGPCCorpo XML
type SNGPCCorpo struct {
	XMLName      xml.Name            `xmml:"corpo"`
	Medicamentos []SNGPCMedicamentos `xml:"medicamentos"`
	Insumos      []SNGPCInsumos      `xml:"insumos"`
}

//SNGPCMensagemSNGPCInventario root XML
type SNGPCMensagemSNGPCInventario struct {
	XMLName   xml.Name       `xml:"mensagemSNGPCInventario"`
	Cabecalho SNGPCCabecalho `xml:"cabecalho"`
	Corpo     SNGPCCorpo     `xml:"corpo"`
}

func (s SNGPCMensagemSNGPCInventario) String() string {
	return fmt.Sprintf("%v\n%v", s.Cabecalho, s.Corpo)
}

func (s SNGPCCabecalho) String() string {
	return fmt.Sprintf("cnpj : %v ; cpf : %v ; data : %v\n", s.CnpjEmissor, s.CpfTransmissor, s.Data)
}

func (s SNGPCCorpo) String() string {
	return fmt.Sprintf("medicamentos : %v ; insumos : %v\n", s.Medicamentos, s.Insumos)
}

func main() {
	var sngpc SNGPCMensagemSNGPCInventario
	xml.Unmarshal(XML, &sngpc)
	fmt.Println(sngpc)
}
