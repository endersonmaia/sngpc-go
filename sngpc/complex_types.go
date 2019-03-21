package sngpc

import "fmt"

type Medicamento struct {
	RegistroMSMedicamento    string                   `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string                   `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    string                   `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento UnidadeMedidaMedicamento `xml:"unidadeMedidaMedicamento"`
}

//Medicamentos
//<element name="medicamentos">
// 	<complexType>
// 	  <sequence>
// 		<element name="entradaMedicamentos" type="sngpc:ct_EntradaMedicamento" minOccurs="0" maxOccurs="unbounded" />
// 		<element name="saidaMedicamentoVendaAoConsumidor" type="sngpc:ct_SaidaMedicamentoVendaAoConsumidor" minOccurs="0" maxOccurs="unbounded" />
// 		<element name="saidaMedicamentoTransferencia" type="sngpc:ct_SaidaMedicamentoTransferencia" minOccurs="0" maxOccurs="unbounded" />
// 		<element name="saidaMedicamentoPerda" type="sngpc:ct_SaidaMedicamentoPerda" minOccurs="0" maxOccurs="unbounded" />
// 		<element name="entradaMedicamentoTransformacao" type="sngpc:ct_EntradaMedicamentoTransformacao" minOccurs="0" maxOccurs="unbounded" />
// 		<element name="saidaMedicamentoTransformacaoVendaAoConsumidor" type="sngpc:ct_SaidaMedicamentoTransformacaoVendaAoConsumidor" minOccurs="0" maxOccurs="unbounded" />
// 		<element name="saidaMedicamentoTransformacaoPerda" type="sngpc:ct_SaidaMedicamentoTransformacaoPerda" minOccurs="0" maxOccurs="unbounded" />
// 	  </sequence>
// 	</complexType>
//</element>
type Medicamentos struct {
	EntradaMedicamentos                            []EntradaMedicamentos                            `xml:"entradaMedicamentos"`
	SaidaMedicamentoVendaAoConsumidor              []SaidaMedicamentoVendaAoConsumidor              `xml:"saidaMedicamentoVendaAoConsumidor"`
	SaidaMedicamentoTransferencia                  []SaidaMedicamentoTransferencia                  `xml:"saidaMedicamentoTransferencia"`
	SaidaMedicamentoPerda                          []SaidaMedicamentoPerda                          `xml:"saidaMedicamentoPerda"`
	EntradaMedicamentoTransformacao                []EntradaMedicamentoTransformacao                `xml:"entradaMedicamentoTransformacao"`
	SaidaMedicamentoTransformacaoVendaAoConsumidor []SaidaMedicamentoTransformacaoVendaAoConsumidor `xml:"saidaMedicamentoTransformacaoVendaAoConsumidor"`
	SaidaMedicamentoTransformacaoPerda             []SaidaMedicamentoTransformacaoPerda             `xml:"saidaMedicamentoTransformacaoPerda"`
}

func (s Medicamentos) String() string {
	out := ""

	for _, e := range s.EntradaMedicamentos {
		out += fmt.Sprint(e.MedicamentoEntrada)
	}

	return out
}

//MedicamentoVenda
// <complexType name="ct_MedicamentoVenda">
// <sequence>
//   <element name ="usoProlongado" type="sngpc:st_SimNaoNull" />
//   <element name="registroMSMedicamento" type="sngpc:st_RegistroMS" />
//   <element name="numeroLoteMedicamento" type="sngpc:st_Lote" />
//   <element name="quantidadeMedicamento" type="sngpc:st_QuantidadeMedicamento" />
//   <element name="unidadeMedidaMedicamento" type="sngpc:st_UnidadeMedidaMedicamento" />
// </sequence>
// </complexType>
type MedicamentoVenda struct {
	UsoProlongado            string                   `xml:"usoProlongado"`
	RegistroMSMedicamento    string                   `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string                   `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    uint                     `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento UnidadeMedidaMedicamento `xml:"unidadeMedidaMedicamento"`
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

//NotaFiscal
// <complexType name="ct_NotaFiscal">
// <sequence>
// 	<element name="numeroNotaFiscal" type="sngpc:st_NumeroNotaFiscal" />
// 	<element name="tipoOperacaoNotaFiscal" type="sngpc:st_TipoOperacaoNotaFiscal" />
// 	<element name="dataNotaFiscal" type="sngpc:st_data" />
// 	<element name="cnpjOrigem" type="sngpc:st_CNPJ" />
// 	<element name="cnpjDestino" type="sngpc:st_CNPJ" />
// </sequence>
// </complexType>
type NotaFiscal struct {
	NumeroNotaFiscal       string                 `xml:"numeroNotaFiscal"`
	TipoOperacaoNotaFiscal TipoOperacaoNotaFiscal `xml:"tipoOperacaoNotaFiscal"`
	DataNotaFiscal         string                 `xml:"dataNotaFiscal"`
	CnpjOrigem             string                 `xml:"cnpjOrigem"`
	CnpjDestino            string                 `xml:"cnpjDestino"`
}

// MedicamentoSaidaTransformacao
type MedicamentoSaidaTransformacao struct {
	RegistroMSMedicamento   string            `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento   string            `xml:"numeroLoteMedicamento"`
	QuantidadeInsumo        float32           `xml:"quantidadeInsumo"`
	UnidadeDeMedidaDoInsumo TipoUnidadeInsumo `xml:"unidadeDeMedidaDoInsumo"`
}

// MedicamentoTransformacao
type MedicamentoTransformacao struct {
	RegistroMSMedicamento    string                   `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string                   `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    string                   `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento UnidadeMedidaMedicamento `xml:"unidadeMedidaMedicamento"`
	QuantidadeInsumo         float32                  `xml:"quantidadeInsumo"`
	UnidadeDeMedidaDoInsumo  TipoUnidadeInsumo        `xml:"unidadeDeMedidaDoInsumo"`
}

// MedicamentoTransformacaoVenda
type MedicamentoTransformacaoVenda struct {
	UsoProlongado                              string                    `xml:"usoProlongado"`
	RegistroMSMedicamento                      string                    `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento                      string                    `xml:"numeroLoteMedicamento"`
	QuantidadeDeInsumoPorUnidadeFarmacotecnica float32                   `xml:"quantidadeDeInsumoPorUnidadeFarmacotecnica"`
	UnidadeDeMedidaDoInsumo                    TipoUnidadeInsumo         `xml:"unidadeDeMedidaDoInsumo"`
	UnidadeFarmacotecnica                      TipoUnidadeFarmacotecnica `xml:"unidadeFarmacotecnica"`
	QuantidadeDeUnidadesFarmacotecnicas        float32                   `xml:"quantidadeDeUnidadesFarmacotecnicas"`
}

// Insumo
type Insumo struct {
	CodigoInsumo         string `xml:"codigoInsumo"`
	NumeroLoteInsumo     string `xml:"numeroLoteInsumo"`
	InsumoCNPJFornecedor string `xml:"insumoCNPJFornecedor"`
}

// InsumoEntrada
type InsumoEntrada struct {
	InsumoEntrada           Insumo            `xml:"insumoEntrada"`
	QuantidadeInsumoEntrada float32           `xml:"quantidadeInsumoEntrada"`
	TipoUnidadeEntrada      TipoUnidadeInsumo `xml:"tipoUnidadeEntrada"`
}

// InsumoTransferencia
type InsumoTransferencia struct {
	InsumoTransferencia           Insumo            `xml:"insumoTransferencia"`
	QuantidadeInsumoTransferencia float32           `xml:"quantidadeInsumoTransferencia"`
	TipoUnidadeTransferencia      TipoUnidadeInsumo `xml:"tipoUnidadeTransferencia"`
}

// InsumoPerda
type InsumoPerda struct {
	InsumoPerda           Insumo            `xml:"insumoPerda"`
	QuantidadeInsumoPerda float32           `xml:"quantidadeInsumoPerda"`
	TipoUnidadePerda      TipoUnidadeInsumo `xml:"tipoUnidadePerda"`
}

// InsumoVendaAoConsumidor
type InsumoVendaAoConsumidor struct {
	UsoProlongado                              uint8                     `xml:"usoProlongado"`
	InsumoVendaAoConsumidor                    Insumo                    `xml:"insumoVendaAoConsumidor"`
	QuantidadeDeInsumoPorUnidadeFarmacotecnica string                    `xml:"quantidadeDeInsumoPorUnidadeFarmacotecnica"`
	UnidadeDeMedidaDoInsumo                    TipoUnidadeInsumo         `xml:"unidadeDeMedidaDoInsumo"`
	UnidadeFarmacotecnica                      TipoUnidadeFarmacotecnica `xml:"unidadeFarmacotecnica"`
	QuantidadeDeUnidadesFarmacotecnicas        float32                   `xml:"quantidadeDeUnidadesFarmacotecnicas"`
}

// Comprador
type Comprador struct {
	NomeComprador      string         `xml:"nomeComprador"`
	TipoDocumento      TipoDocumento  `xml:"tipoDocumento"`
	NumeroDocumento    string         `xml:"numeroDocumento"`
	OrgaoExpedidor     OrgaoExpedidor `xml:"orgaoExpedidor"`
	UFEmissaoDocumento UF             `xml:"UFEmissaoDocumento"`
}

// Prescritor
type Prescritor struct {
	NomePrescritor             string               `xml:"nomePrescritor"`
	NumeroRegistroProfissional string               `xml:"numeroRegistroProfissional"`
	ConselhoProfissional       ConselhoProfissional `xml:"conselhoProfissional"`
	UFConselho                 UF                   `xml:"UFConselho"`
}

// ct_Retorno
// <complexType name="ct_Retorno">
// <sequence>
//   <element name="hash" type="sngpc:st_Nome" />
//   <element name="dataInicioRetorno" type="sngpc:st_data" />
//   <element name="dataFimRetorno" type="sngpc:st_data" />
//   <element name="dataTransmissao" type="sngpc:st_data" />
//   <element name="dataValidacao" type="sngpc:st_data" />
//   <sequence>
// 	<element name="descricaoValidacao" type="sngpc:st_DescricaoValidacao" minOccurs="0" maxOccurs="unbounded" />
//   </sequence>
//   <element name="validado" type="sngpc:st_simNao" />
// </sequence>
// </complexType>

// Paciente
type Paciente struct {
	Nome         string       `xml:"nome"`
	Idade        string       `xml:"idade"`
	UnidadeIdade UnidadeIdade `xml:"unidadeIdade"`
	Sexo         Sexo         `xml:"sexo"`
	Cid          string       `xml:"cid"`
}

// InsumoBasico
type InsumoBasico struct {
	CodigoInsumo         string            `xml:"codigoInsumo"`
	NumeroLoteInsumo     string            `xml:"numeroLoteInsumo"`
	InsumoCNPJFornecedor string            `xml:"insumoCNPJFornecedor"`
	QuantidadeInsumo     float32           `xml:"quantidadeInsumo"`
	TipoUnidade          TipoUnidadeInsumo `xml:"tipoUnidade"`
}
