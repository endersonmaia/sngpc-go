package sngpc

import (
	"encoding/xml"
	"fmt"
)

// MensagemSNGPCInventario
type MensagemSNGPCInventario struct {
	XMLName   xml.Name            `xml:"mensagemSNGPCInventario"`
	Cabecalho CabecalhoInventario `xml:"cabecalho"`
	Corpo     CorpoInventario     `xml:"corpo"`
}

func (s MensagemSNGPCInventario) String() string {
	return fmt.Sprintf("%v\n%v", s.Cabecalho, s.Corpo)
}

type CabecalhoInventario struct {
	CnpjEmissor    string `xml:"cnpjEmissor"`
	CpfTransmissor string `xml:"cpfTransmissor"`
	Data           string `xml:"data"`
}

func (s CabecalhoInventario) String() string {
	return fmt.Sprintf("Inventário :\n\nCNPJ : %v ; CPF : %v ; Data : %v\n", s.CnpjEmissor, s.CpfTransmissor, s.Data)
}

type CorpoInventario struct {
	Medicamentos Medicamentos `xml:"medicamentos"`
	Insumos      Insumos      `xml:"insumos"`
}

func (s CorpoInventario) String() string {
	return fmt.Sprintf("Corpo : \n\nMedicamentos : \n%v\nInsumos : \n%v\n", s.Medicamentos, s.Insumos)
}

//MedicamentoEntrada type struct
// <complexType name="ct_MedicamentoEntrada">
// <sequence>
//   <element name="classeTerapeutica" type="sngpc:st_classeTerapeutica" />
//   <element name="registroMSMedicamento" type="sngpc:st_RegistroMS" />
//   <element name="numeroLoteMedicamento" type="sngpc:st_Lote" />
//   <element name="quantidadeMedicamento" type="sngpc:st_QuantidadeMedicamento" />
//   <element name="unidadeMedidaMedicamento" type="sngpc:st_UnidadeMedidaMedicamento" />
// </sequence>
// </complexType>
type MedicamentoEntrada struct {
	ClasseTerapeutica        ClasseTerapeutica        `xml:"classeTerapeutica"`
	RegistroMSMedicamento    string                   `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string                   `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    uint                     `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento UnidadeMedidaMedicamento `xml:"unidadeMedidaMedicamento"`
}

func (s MedicamentoEntrada) String() string {
	return fmt.Sprintf("RegistroMS : %v, Lote : %v, Quantidade : %v\n", s.RegistroMSMedicamento, s.NumeroLoteMedicamento, s.QuantidadeMedicamento)
}

//InsumoBasicoEntrada
// <complexType name="ct_InsumoBasicoEntrada">
// <sequence>
//   <element name="classeTerapeutica" type="sngpc:st_classeTerapeutica" />
//   <element name="codigoInsumo" type="sngpc:st_CodigoDCB" />
//   <element name="numeroLoteInsumo" type="sngpc:st_Lote" />
//   <element name="insumoCNPJFornecedor" type="sngpc:st_CNPJ" />
//   <element name="quantidadeInsumo" type="sngpc:st_QuantidadeInsumo" />
//   <element name="tipoUnidade" type="sngpc:st_TipoUnidadeInsumo" />
// </sequence>
// </complexType>
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
