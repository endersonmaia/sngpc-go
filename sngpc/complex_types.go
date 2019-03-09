package sngpc

import "fmt"

//Medicamento
// <complexType name="ct_Medicamento">
//     <sequence>
//       <element name="registroMSMedicamento" type="sngpc:st_RegistroMS" />
//       <element name="numeroLoteMedicamento" type="sngpc:st_Lote" />
//       <element name="quantidadeMedicamento" type="sngpc:st_QuantidadeMedicamento" />
//       <element name="unidadeMedidaMedicamento" type="sngpc:st_UnidadeMedidaMedicamento" />
//     </sequence>
//   </complexType>
type Medicamento struct {
	RegistroMSMedicamento    string `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    string `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento string `xml:"unidadeMedidaMedicamento"`
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
	UsoProlongado            string `xml:"usoProlongado"`
	RegistroMSMedicamento    string `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    uint   `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento uint8  `xml:"unidadeMedidaMedicamento"`
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
	ClasseTerapeutica        uint8  `xml:"classeTerapeutica"`
	RegistroMSMedicamento    string `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    uint   `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento uint8  `xml:"unidadeMedidaMedicamento"`
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
	NumeroNotaFiscal       string `xml:"numeroNotaFiscal"`
	TipoOperacaoNotaFiscal uint8  `xml:"tipoOperacaoNotaFiscal"`
	DataNotaFiscal         string `xml:"dataNotaFiscal"`
	CnpjOrigem             string `xml:"cnpjOrigem"`
	CnpjDestino            string `xml:"cnpjDestino"`
}

//MedicamentoSaidaTransformacao
// <complexType name="ct_MedicamentoSaidaTransformacao">
// <sequence>
//   <element name="registroMSMedicamento" type="sngpc:st_RegistroMS" />
//   <element name="numeroLoteMedicamento" type="sngpc:st_Lote" />
//   <element name="quantidadeInsumo" type="sngpc:st_QuantidadeInsumo" />
//   <element name="unidadeDeMedidaDoInsumo" type="sngpc:st_TipoUnidadeInsumo" />
// </sequence>
// </complexType>
type MedicamentoSaidaTransformacao struct {
	RegistroMSMedicamento   string  `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento   string  `xml:"numeroLoteMedicamento"`
	QuantidadeInsumo        float32 `xml:"quantidadeInsumo"`
	UnidadeDeMedidaDoInsumo uint8   `xml:"unidadeDeMedidaDoInsumo"`
}

//MedicamentoTransformacao
// <complexType name="ct_MedicamentoTransformacao">
// <sequence>
//   <!--<element name="classeTerapeutica" type = "sngpc:st_classeTerapeutica" />-->
//   <element name="registroMSMedicamento" type="sngpc:st_RegistroMS" />
//   <element name="numeroLoteMedicamento" type="sngpc:st_Lote" />
//   <element name="quantidadeMedicamento" type="sngpc:st_QuantidadeMedicamento" />
//   <element name="unidadeMedidaMedicamento" type="sngpc:st_UnidadeMedidaMedicamento" />
//   <element name="quantidadeInsumo" type="sngpc:st_QuantidadeInsumo" />
//   <element name="unidadeDeMedidaDoInsumo" type="sngpc:st_TipoUnidadeInsumo" />
// </sequence>
// </complexType>
type MedicamentoTransformacao struct {
	RegistroMSMedicamento    string  `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string  `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    string  `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento string  `xml:"unidadeMedidaMedicamento"`
	QuantidadeInsumo         float32 `xml:"quantidadeInsumo"`
	UnidadeDeMedidaDoInsumo  string  `xml:"unidadeDeMedidaDoInsumo"`
}

//MedicamentoTransformacaoVenda
// <complexType name="ct_MedicamentoTransformacaoVenda">
// <sequence>
//   <element name="usoProlongado" type="sngpc:st_SimNaoNull"/>
//   <element name="registroMSMedicamento" type="sngpc:st_RegistroMS" />
//   <element name="numeroLoteMedicamento" type="sngpc:st_Lote" />
//   <element name="quantidadeDeInsumoPorUnidadeFarmacotecnica" type="sngpc:st_QuantidadeDeInsumoPorUnidadeFarmacotecnica" />
//   <element name="unidadeDeMedidaDoInsumo" type="sngpc:st_TipoUnidadeInsumo" />
//   <element name="unidadeFarmacotecnica" type="sngpc:st_TipoUnidadeFarmacotecnica" />
//   <element name="quantidadeDeUnidadesFarmacotecnicas" type="sngpc:st_QuantidadeDeUnidadesFarmacotecnicas" />
// </sequence>
// </complexType>
type MedicamentoTransformacaoVenda struct {
	UsoProlongado                              string  `xml:"usoProlongado"`
	RegistroMSMedicamento                      string  `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento                      string  `xml:"numeroLoteMedicamento"`
	QuantidadeDeInsumoPorUnidadeFarmacotecnica float32 `xml:"quantidadeDeInsumoPorUnidadeFarmacotecnica"`
	UnidadeDeMedidaDoInsumo                    uint8   `xml:"unidadeDeMedidaDoInsumo"`
	UnidadeFarmacotecnica                      uint8   `xml:"unidadeFarmacotecnica"`
	QuantidadeDeUnidadesFarmacotecnicas        float32 `xml:"quantidadeDeUnidadesFarmacotecnicas"`
}

//Insumo
// <complexType name="ct_Insumo">
// <sequence>
//   <element name="codigoInsumo" type="sngpc:st_CodigoDCB" />
//   <element name="numeroLoteInsumo" type="sngpc:st_Lote" />
//   <element name="insumoCNPJFornecedor" type="sngpc:st_CNPJ" />
// </sequence>
// </complexType>
type Insumo struct {
	CodigoInsumo         string `xml:"codigoInsumo"`
	NumeroLoteInsumo     string `xml:"numeroLoteInsumo"`
	InsumoCNPJFornecedor string `xml:"insumoCNPJFornecedor"`
}

//InsumoEntrada
// <complexType name="ct_InsumoEntrada">
//     <sequence>
//       <element name="insumoEntrada" type="sngpc:ct_Insumo" />
//       <element name="quantidadeInsumoEntrada" type="sngpc:st_QuantidadeInsumo" />
//       <element name="tipoUnidadeEntrada" type="sngpc:st_TipoUnidadeInsumo" />
//     </sequence>
//   </complexType>
type InsumoEntrada struct {
	InsumoEntrada           Insumo  `xml:"insumoEntrada"`
	QuantidadeInsumoEntrada float32 `xml:"quantidadeInsumoEntrada"`
	TipoUnidadeEntrada      uint8   `xml:"tipoUnidadeEntrada"`
}

//InsumoTransferencia
// <complexType name="ct_InsumoTransferencia">
// <sequence>
//   <element name="insumoTransferencia" type="sngpc:ct_Insumo" />
//   <element name="quantidadeInsumoTransferencia" type="sngpc:st_QuantidadeInsumo" />
//   <element name="tipoUnidadeTransferencia" type="sngpc:st_TipoUnidadeInsumo" />
// </sequence>
// </complexType>
type InsumoTransferencia struct {
	InsumoTransferencia           Insumo  `xml:"insumoTransferencia"`
	QuantidadeInsumoTransferencia float32 `xml:"quantidadeInsumoTransferencia"`
	TipoUnidadeTransferencia      uint8   `xml:"tipoUnidadeTransferencia"`
}

//InsumoPerda
// <complexType name="ct_InsumoPerda">
// <sequence>
//   <element name="insumoPerda" type="sngpc:ct_Insumo" />
//   <element name="quantidadeInsumoPerda" type="sngpc:st_QuantidadeInsumo" />
//   <element name="tipoUnidadePerda" type="sngpc:st_TipoUnidadeInsumo" />
// </sequence>
// </complexType>
type InsumoPerda struct {
	InsumoPerda           Insumo  `xml:"insumoPerda"`
	QuantidadeInsumoPerda float32 `xml:"quantidadeInsumoPerda"`
	TipoUnidadePerda      uint8   `xml:"tipoUnidadePerda"`
}

//InsumoVendaAoConsumidor
// <complexType name="ct_InsumoVendaAoConsumidor">
// <sequence>
//   <element name = "usoProlongado" type = "sngpc:st_SimNaoNull"/>
//   <element name="insumoVendaAoConsumidor" type="sngpc:ct_Insumo" />
//   <element name="quantidadeDeInsumoPorUnidadeFarmacotecnica" type="sngpc:st_QuantidadeDeInsumoPorUnidadeFarmacotecnica" />
//   <element name="unidadeDeMedidaDoInsumo" type="sngpc:st_TipoUnidadeInsumo" />
//   <element name="unidadeFarmacotecnica" type="sngpc:st_TipoUnidadeFarmacotecnica" />
//   <element name="quantidadeDeUnidadesFarmacotecnicas" type="sngpc:st_QuantidadeDeUnidadesFarmacotecnicas" />
// </sequence>
// </complexType>
type InsumoVendaAoConsumidor struct {
	UsoProlongado                              uint8  `xml:"usoProlongado"`
	InsumoVendaAoConsumidor                    Insumo `xml:"insumoVendaAoConsumidor"`
	QuantidadeDeInsumoPorUnidadeFarmacotecnica string `xml:"quantidadeDeInsumoPorUnidadeFarmacotecnica"`
	UnidadeDeMedidaDoInsumo                    string `xml:"unidadeDeMedidaDoInsumo"`
	UnidadeFarmacotecnica                      string `xml:"unidadeFarmacotecnica"`
	QuantidadeDeUnidadesFarmacotecnicas        string `xml:"quantidadeDeUnidadesFarmacotecnicas"`
}

//Comprador
// <complexType name="ct_Comprador">
// <sequence>
//   <element name="nomeComprador" type="sngpc:st_Nome" />
//   <element name="tipoDocumento" type="sngpc:st_TipoDocumento" />
//   <element name="numeroDocumento" type="sngpc:st_NumeroDocumento" />
//   <element name="orgaoExpedidor" type="sngpc:st_OrgaoExpedidor" />
//   <element name="UFEmissaoDocumento" type="sngpc:st_UF" />
// </sequence>
// </complexType>
type Comprador struct {
	NomeComprador      string `xml:"nomeComprador"`
	TipoDocumento      uint8  `xml:"tipoDocumento"`
	NumeroDocumento    string `xml:"numeroDocumento"`
	OrgaoExpedidor     string `xml:"orgaoExpedidor"`
	UFEmissaoDocumento string `xml:"UFEmissaoDocumento"`
}

//Prescritor
// <complexType name="ct_Prescritor">
//     <sequence>
//       <element name="nomePrescritor" type="sngpc:st_Nome" />
//       <element name="numeroRegistroProfissional" type="sngpc:st_NumeroDocumento" />
//       <element name="conselhoProfissional" type="sngpc:st_ConselhoProfissional" />
//       <element name="UFConselho" type="sngpc:st_UF" />
//     </sequence>
//   </complexType>
type Prescritor struct {
	NomePrescritor             string `xml:"nomePrescritor"`
	NumeroRegistroProfissional string `xml:"numeroRegistroProfissional"`
	ConselhoProfissional       string `xml:"conselhoProfissional"`
	UFConselho                 string `xml:"UFConselho"`
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

//Paciente
// <complexType name="ct_paciente">
// <sequence>
//   <element name="nome" type="sngpc:st_Nome" />
//   <element name="idade" type="sngpc:st_Idade" />
//   <element name="unidadeIdade" type="sngpc:st_UnidadeIdade" />
//   <element name="sexo" type="sngpc:st_Sexo" />
//   <element name="cid" type="sngpc:st_Cid" />
// </sequence>
// </complexType>
type Paciente struct {
	Nome         string `xml:"nome"`
	Idade        string `xml:"idade"`
	UnidadeIdade string `xml:"unidadeIdade"`
	Sexo         string `xml:"sexo"`
	Cid          string `xml:"cid"`
}

//InsumoBasico
// <complexType name="ct_InsumoBasico">
// <sequence>
//   <element name="codigoInsumo" type="sngpc:st_CodigoDCB" />
//   <element name="numeroLoteInsumo" type="sngpc:st_Lote" />
//   <element name="insumoCNPJFornecedor" type="sngpc:st_CNPJ" />
//   <element name="quantidadeInsumo" type="sngpc:st_QuantidadeInsumo" />
//   <element name="tipoUnidade" type="sngpc:st_TipoUnidadeInsumo" />
// </sequence>
// </complexType>
type InsumoBasico struct {
	CodigoInsumo         string  `xml:"codigoInsumo"`
	NumeroLoteInsumo     string  `xml:"numeroLoteInsumo"`
	InsumoCNPJFornecedor string  `xml:"insumoCNPJFornecedor"`
	QuantidadeInsumo     float32 `xml:"quantidadeInsumo"`
	TipoUnidade          uint8   `xml:"tipoUnidade"`
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
	UnidadeMedidaInsumo  uint8             `xml:"unidadeMedidaInsumo"`
}

func (s InsumoBasicoEntrada) String() string {
	return fmt.Sprintf("RegistroMS : %v, Lote : %v, Quantidade : %v\n", s.CodigoInsumo, s.NumeroLoteInsumo, s.QuantidadeInsumo)
}
