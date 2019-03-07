package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html/charset"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//CabecalhoInventario
type CabecalhoInventario struct {
	XMLName        xml.Name `xml:"cabecalho"`
	CnpjEmissor    string   `xml:"cnpjEmissor"`
	CpfTransmissor string   `xml:"cpfTransmissor"`
	Data           string   `xml:"data"`
}

//CabecalhoMovimentacao
type CabecalhoMovimentacao struct {
	XMLName        xml.Name `xml:"cabecalho"`
	CnpjEmissor    string   `xml:"cnpjEmissor"`
	CpfTransmissor string   `xml:"cpfTransmissor"`
	DataInicio     string   `xml:"dataInicio"`
	DataFim        string   `xml:"dataFim"`
}

//Insumos
type Insumos struct {
	XMLName        xml.Name         `xml:"insumos"`
	EntradaInsumos []EntradaInsumos `xml:"entradaInsumos"`
}

//EntradaInsumos
type EntradaInsumos struct {
	XMLName       xml.Name      `xml:"entradaInsumos"`
	InsumoEntrada InsumoEntrada `xml:"insumoEntrada"`
}

//InsumoEntrada
type InsumoEntrada struct {
	XMLName              xml.Name `xml:"insumoEntrada"`
	ClasseTerapeutica    uint     `xml:"classeTerapeutica"`
	CodigoInsumo         string   `xml:"codigoInsumo"`
	NumeroLoteInsumo     string   `xml:"numeroLoteInsumo"`
	InsumoCNPJFornecedor string   `xml:"insumoCNPJFornecedor"`
	QuantidadeInsumo     float32  `xml:"quantidadeInsumo"`
	UnidadeMedidaInsumo  uint     `xml:"unidadeMedidaInsumo"`
}

//Medicamentos XML
type Medicamentos struct {
	XMLName                           xml.Name                            `xml:"medicamentos"`
	EntradaMedicamentos               []EntradaMedicamentos               `xml:"entradaMedicamentos"`
	SaidaMedicamentoVendaAoConsumidor []SaidaMedicamentoVendaAoConsumidor `xml:"daidaMedicamentoVendaAoConsumidor"`
	// name="saidaMedicamentoTransferencia
	// name="saidaMedicamentoPerda
	// name="entradaMedicamentoTransformacao
	// name="saidaMedicamentoTransformacaoVendaAoConsumidor
	// name="saidaMedicamentoTransformacaoPerda
}

//EntradaMedicamentos
type EntradaMedicamentos struct {
	XMLName            xml.Name           `xml:"entradaMedicamentos"`
	MedicamentoEntrada MedicamentoEntrada `xml:"medicamentoEntrada"`
}

type SaidaMedicamentoVendaAoConsumidor struct {
	XMLName                      xml.Name              `xml:"saidaMedicamentoVendaAoConsumidos"`
	TipoReceituarioMedicamento   string                `xml:"tipoReceituarioMedicamento"`
	NumeroNotificacaoMedicamento string                `xml:"numeroNotificacaoMedicamento"`
	DataPrescricaoMedicamento    string                `xml:"dataPrescricaoMedicamento"`
	PrescritorMedicamento        PrescritorMedicamento `xml:"prescritorMedicamento"`
	UsoMedicamento               string                `xml:"usoMedicamento"`
	CompradorMedicamento         CompradorMedicamento  `xml:"compradorMedicamento"`
	PacienteMedicamento          PacienteMedicamento   `xml:"pacienteMedicamento"`
	MedicamentoVenda             []MedicamentoVenda    `xml:"medicamentoVenda"`
	DataVendaMedicamento         string                `xml:"dataVendaMedicamento"`
}

type PrescritorMedicamento struct {
	XMLName                    xml.Name `xml:"prescritorMedicamento"`
	NomePrescritor             string   `xml:"nomePrescritor"`
	NumeroRegistroProfissional string   `xml:"numeroRegistroProfissional"`
	ConselhoProfissional       string   `xml:"conselhoProfissional"`
	UFConselho                 string   `xml:"UFConselho"`
}

type CompradorMedicamento struct {
	XMLName            xml.Name `xml:"compradorMedicamento"`
	NomeComprador      string   `xml:"nomeComprador"`
	TipoDocumento      string   `xml:"tipoDocumento"`
	NumeroDocumento    string   `xml:"numeroDocumento"`
	OrgaoExpedidor     string   `xml:"orgaoExpedidor"`
	UFEmissaoDocumento string   `xml:"UFEmissaoDocumento"`
}

type MedicamentoVenda struct {
	XMLName                  xml.Name `xml:"medicamentoVenda"`
	UsoProlongado            string   `xml:"usoProlongado"`
	RegistroMSMedicamento    string   `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string   `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    string   `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento string   `xml:"unidadeMedidaMedicamento"`
}

type PacienteMedicamento struct {
	XMLName      xml.Name `xml:"pacienteMedicamento"`
	Nome         string   `xml:"nome"`
	Idade        string   `xml:"idade"`
	UnidadeIdade string   `xml:"unidadeIdade"`
	Sexo         string   `xml:"sexo"`
	Cid          string   `xml:"cid"`
}

//MedicamentoEntrada
type MedicamentoEntrada struct {
	XMLName                  xml.Name `xml:"medicamentoEntrada"`
	ClasseTerapeutica        uint     `xml:"classeTerapeutica"`
	RegistroMSMedicamento    string   `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string   `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    uint     `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento uint     `xml:"unidadeMedidaMedicamento"`
}

//Corpo
type CorpoInventario struct {
	XMLName      xml.Name     `xmml:"corpo"`
	Medicamentos Medicamentos `xml:"medicamentos"`
	Insumos      Insumos      `xml:"insumos"`
}

//Corpo
type CorpoMovimentacao struct {
	XMLName      xml.Name     `xmml:"corpo"`
	Medicamentos Medicamentos `xml:"medicamentos"`
	Insumos      Insumos      `xml:"insumos"`
}

//MensagemSNGPCInventario
type MensagemSNGPCInventario struct {
	XMLName   xml.Name            `xml:"mensagemSNGPCInventario"`
	Cabecalho CabecalhoInventario `xml:"cabecalho"`
	Corpo     CorpoInventario     `xml:"corpo"`
}

//MensagemSNGPC
type MensagemSNGPC struct {
	XMLName   xml.Name              `xml:"mensagemSNGPC"`
	Cabecalho CabecalhoMovimentacao `xml:"cabecalho"`
	Corpo     CorpoMovimentacao     `xml:"corpo"`
}

func (s MensagemSNGPC) String() string {
	return fmt.Sprintf("%v\n%v", s.Cabecalho, s.Corpo)
}

func (s CabecalhoMovimentacao) String() string {
	return fmt.Sprintf("cnpj : %v ; cpf : %v ; inicio : %v, fim : %v\n", s.CnpjEmissor, s.CpfTransmissor, s.DataInicio, s.DataFim)
}

func (s CorpoMovimentacao) String() string {
	out := "corpo\n"
	out += string(len(s.Medicamentos.SaidaMedicamentoVendaAoConsumidor))

	for _, s := range s.Medicamentos.SaidaMedicamentoVendaAoConsumidor {
		for _, m := range s.MedicamentoVenda {
			out += fmt.Sprintf("medicamento : %v, lote : %v, quantidade : %v\n", m.RegistroMSMedicamento, m.NumeroLoteMedicamento, m.QuantidadeMedicamento)
		}
	}
	return out
}

//Inventario
func (s MensagemSNGPCInventario) String() string {
	return fmt.Sprintf("%v\n%v", s.Cabecalho, s.Corpo)
}

func (s CabecalhoInventario) String() string {
	return fmt.Sprintf("Inventário :\nCNPJ : %v ; CPF : %v ; Data : %v\n", s.CnpjEmissor, s.CpfTransmissor, s.Data)
}

func (s CorpoInventario) String() string {
	return fmt.Sprintf("Corpo : \nMedicamentos : \n%v\nInsumos : \n%v\n", s.Medicamentos, s.Insumos)
}

func (s Medicamentos) String() string {
	out := ""

	for _, e := range s.EntradaMedicamentos {
		out += fmt.Sprint(e.MedicamentoEntrada)
	}

	return out
}

func (s Insumos) String() string {
	out := ""

	for _, e := range s.EntradaInsumos {
		out += fmt.Sprint(e.InsumoEntrada)
	}

	return out
}

func (s MedicamentoEntrada) String() string {
	return fmt.Sprintf("RegistroMS : %v, Lote : %v, Quantidade : %v\n", s.RegistroMSMedicamento, s.NumeroLoteMedicamento, s.QuantidadeMedicamento)
}

func (s InsumoEntrada) String() string {
	return fmt.Sprintf("RegistroMS : %v, Lote : %v, Quantidade : %v\n", s.CodigoInsumo, s.NumeroLoteInsumo, s.QuantidadeInsumo)
}

func main() {
	xmlFile, err := os.Open("inventario.xml")
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

		sngpc := MensagemSNGPCInventario{}

		err = dec.Decode(&sngpc)
		check(err)

		fmt.Printf("\n%s", sngpc)
	} else if strings.Contains(string(buffer), `<mensagemSNGPC xmlns="urn:sngpc-schema">`) {
		_, err = xmlFile.Seek(0, 0)
		check(err)

		sngpc := MensagemSNGPC{}

		err = dec.Decode(&sngpc)
		check(err)

		fmt.Printf("\n%s", sngpc)
	}

}
