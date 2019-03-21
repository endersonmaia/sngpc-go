package sngpc

import (
	"encoding/xml"
	"fmt"
)

// MensagemSNGPCInventario armazena o conteúdo de um arquivo de inventário do SNGPC
// O arquivo de inventário contém uma data de referência e os registros de inventário
// para medicamentos e insumos
type MensagemSNGPCInventario struct {
	XMLName   xml.Name            `xml:"mensagemSNGPCInventario"`
	Cabecalho CabecalhoInventario `xml:"cabecalho"`
	Corpo     CorpoInventario     `xml:"corpo"`
}

func (s MensagemSNGPCInventario) String() string {
	return fmt.Sprintf("%v\n%v", s.Cabecalho, s.Corpo)
}

// CabecalhoInventario armazena informações do arquivo de inventário
type CabecalhoInventario struct {
	CnpjEmissor    string `xml:"cnpjEmissor"`
	CpfTransmissor string `xml:"cpfTransmissor"`
	Data           string `xml:"data"`
}

func (s CabecalhoInventario) String() string {
	return fmt.Sprintf("Inventário :\n\nCNPJ : %v ; CPF : %v ; Data : %v\n", s.CnpjEmissor, s.CpfTransmissor, s.Data)
}

// CorpoInventario armazena informações dos registros de inventário
type CorpoInventario struct {
	Medicamentos Medicamentos `xml:"medicamentos"`
	Insumos      Insumos      `xml:"insumos"`
}

func (s CorpoInventario) String() string {
	return fmt.Sprintf("Corpo : \n\nMedicamentos : \n%v\nInsumos : \n%v\n", s.Medicamentos, s.Insumos)
}

// InsumoBasicoEntrada armazena os registros de entrada de insumos do inventário
type InsumoBasicoEntrada struct {
	ClasseTerapeutica    ClasseTerapeutica `xml:"classeTerapeutica"`
	CodigoInsumo         string            `xml:"codigoInsumo"`
	NumeroLoteInsumo     string            `xml:"numeroLoteInsumo"`
	InsumoCNPJFornecedor string            `xml:"insumoCNPJFornecedor"`
	QuantidadeInsumo     float32           `xml:"quantidadeInsumo"`
	UnidadeMedidaInsumo  TipoUnidadeInsumo `xml:"unidadeMedidaInsumo"`
}

func (s InsumoBasicoEntrada) String() string {
	return fmt.Sprintf("RegistroMS : %v, Lote : %v, Quantidade : %v, Unidade: %v\n", s.CodigoInsumo, s.NumeroLoteInsumo, s.QuantidadeInsumo, s.UnidadeMedidaInsumo)
}
