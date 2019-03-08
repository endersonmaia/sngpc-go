package sngpc

import "encoding/xml"

//MensagemSNGPCInventario
type MensagemSNGPCInventario struct {
	XMLName   xml.Name            `xml:"mensagemSNGPCInventario"`
	Cabecalho CabecalhoInventario `xml:"cabecalho"`
	Corpo     CorpoInventario     `xml:"corpo"`
}

//CabecalhoInventario
type CabecalhoInventario struct {
	XMLName        xml.Name `xml:"cabecalho"`
	CnpjEmissor    string   `xml:"cnpjEmissor"`
	CpfTransmissor string   `xml:"cpfTransmissor"`
	Data           string   `xml:"data"`
}

//Corpo
type CorpoInventario struct {
	XMLName      xml.Name     `xmml:"corpo"`
	Medicamentos Medicamentos `xml:"medicamentos"`
	Insumos      Insumos      `xml:"insumos"`
}

//MensagemSNGPC
type MensagemSNGPC struct {
	XMLName   xml.Name              `xml:"mensagemSNGPC"`
	Cabecalho CabecalhoMovimentacao `xml:"cabecalho"`
	Corpo     CorpoMovimentacao     `xml:"corpo"`
}

//Corpo
type CorpoMovimentacao struct {
	XMLName      xml.Name     `xmml:"corpo"`
	Medicamentos Medicamentos `xml:"medicamentos"`
	Insumos      Insumos      `xml:"insumos"`
}

//CabecalhoMovimentacao
// <complexType>
// <sequence>
//   <element name="cnpjEmissor" type="sngpc:st_CNPJ" />
//   <element name="cpfTransmissor" type="sngpc:st_CPF" />
//   <element name="dataInicio" type="sngpc:st_data" />
//   <element name="dataFim" type="sngpc:st_data" />
// </sequence>
// </complexType>
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
	XMLName                                        xml.Name                                         `xml:"medicamentos"`
	EntradaMedicamentos                            []EntradaMedicamentos                            `xml:"entradaMedicamentos"`
	SaidaMedicamentoVendaAoConsumidor              []SaidaMedicamentoVendaAoConsumidor              `xml:"saidaMedicamentoVendaAoConsumidor"`
	SaidaMedicamentoTransferencia                  []SaidaMedicamentoTransferencia                  `xml:"saidaMedicamentoTransferencia"`
	SaidaMedicamentoPerda                          []SaidaMedicamentoPerda                          `xml:"saidaMedicamentoPerda"`
	EntradaMedicamentoTransformacao                []EntradaMedicamentoTransformacao                `xml:"entradaMedicamentoTransformacao"`
	SaidaMedicamentoTransformacaoVendaAoConsumidor []SaidaMedicamentoTransformacaoVendaAoConsumidor `xml:"saidaMedicamentoTransformacaoVendaAoConsumidor"`
	SaidaMedicamentoTransformacaoPerda             []SaidaMedicamentoTransformacaoPerda             `xml:"saidaMedicamentoTransformacaoPerda"`
}

//EntradaMedicamentos
// <complexType name = "ct_EntradaMedicamento">
// <sequence>
// 	<element name = "notaFiscalEntradaMedicamento" type = "sngpc:ct_NotaFiscal"/>
// 	<element name = "medicamentoEntrada" type = "sngpc:ct_MedicamentoEntrada" minOccurs = "0" maxOccurs = "unbounded"/>
// 	<element name = "dataRecebimentoMedicamento" type = "sngpc:st_data"/>
// </sequence>
// </complexType>
type EntradaMedicamentos struct {
	XMLName                      xml.Name             `xml:"entradaMedicamentos"`
	NotaFiscalEntradaMedicamento NotaFiscal           `xml:"notaFiscalEntradaMedicamento"`
	MedicamentoEntrada           []MedicamentoEntrada `xml:"medicamentoEntrada"`
	DataRecebimentoMedicamento   string               `xml:"dataRecebimentoMedicamento"`
}

//SaidaMedicamentoVendaAoConsumidor
// <complexType name = "ct_SaidaMedicamentoVendaAoConsumidor">
// <sequence>
// 	<element name = "tipoReceituarioMedicamento" type = "sngpc:st_TipoReceituario"/>
// 	<element name = "numeroNotificacaoMedicamento" type = "sngpc:st_Notificacao"/>
// 	<element name = "dataPrescricaoMedicamento" type = "sngpc:st_data"/>
// 	<element name = "prescritorMedicamento" type = "sngpc:ct_Prescritor"/>
// 	<element name = "usoMedicamento" type = "sngpc:st_TipoUsoMedicamento"/>
// 	<element name = "compradorMedicamento" type = "sngpc:ct_Comprador" minOccurs="0" maxOccurs="1" />
// 	<element name = "pacienteMedicamento" type = "sngpc:ct_paciente" minOccurs="0" maxOccurs="1" />
// 	<element name = "medicamentoVenda" type = "sngpc:ct_MedicamentoVenda" minOccurs = "0" maxOccurs = "unbounded" />
// 	<element name = "dataVendaMedicamento" type = "sngpc:st_data"/>
// </sequence>
// </complexType>
type SaidaMedicamentoVendaAoConsumidor struct {
	XMLName                      xml.Name           `xml:"saidaMedicamentoVendaAoConsumidor"`
	TipoReceituarioMedicamento   uint8              `xml:"tipoReceituarioMedicamento"`
	NumeroNotificacaoMedicamento string             `xml:"numeroNotificacaoMedicamento"`
	DataPrescricaoMedicamento    string             `xml:"dataPrescricaoMedicamento"`
	PrescritorMedicamento        Prescritor         `xml:"prescritorMedicamento"`
	UsoMedicamento               uint8              `xml:"usoMedicamento"`
	CompradorMedicamento         Comprador          `xml:"compradorMedicamento"`
	PacienteMedicamento          Paciente           `xml:"pacienteMedicamento"`
	MedicamentoVenda             []MedicamentoVenda `xml:"medicamentoVenda"`
	DataVendaMedicamento         string             `xml:"dataVendaMedicamento"`
}

//SaidaMedicamentoTransferencia
// <complexType name = "ct_SaidaMedicamentoTransferencia">
// <sequence>
// 	<element name = "notaFiscalTransferenciaMedicamento" type = "sngpc:ct_NotaFiscal" />
// 	<element name = "medicamentoTransferencia" type = "sngpc:ct_Medicamento" minOccurs = "0" maxOccurs = "unbounded"/>
// 	<element name = "dataTransferenciaMedicamento" type = "sngpc:st_data"/>
// </sequence>
// </complexType>
type SaidaMedicamentoTransferencia struct {
	XMLName                            xml.Name      `xml:"saidaMedicamentoTransferencia"`
	NotaFiscalTransferenciaMedicamento NotaFiscal    `xml:"notaFiscalTransferenciaMedicamento"`
	MedicamentoTransferencia           []Medicamento `xml:"medicamentoTransferencia"`
	DataTransferenciaMedicamento       string        `xml:"dataTransferenciaMedicamento"`
}

//SaidaMedicamentoPerda
// <complexType name = "ct_SaidaMedicamentoPerda">
// <sequence>
// 	<element name = "motivoPerdaMedicamento" type = "sngpc:st_TipoMotivoPerda"/>
// 	<element name = "medicamentoPerda" type = "sngpc:ct_Medicamento"/>
// 	<element name = "dataPerdaMedicamento" type = "sngpc:st_data"/>
// </sequence>
// </complexType>
type SaidaMedicamentoPerda struct {
	XMLName                xml.Name    `xml:"saidaMedicamentoPerda"`
	MotivoPerdaMedicamento uint8       `xml:"motivoPerdaMedicamento"`
	MedicamentoPerda       Medicamento `xml:"medicamentoPerda"`
	DataPerdaMedicamento   string      `xml:"dataPerdaMedicamento"`
}

//EntradaMedicamentoTransformacao
// <complexType name = "ct_EntradaMedicamentoTransformacao">
// <sequence>
// 	<element name = "medicamentoTransformacaoEntrada" type="sngpc:ct_MedicamentoTransformacao" minOccurs = "0" maxOccurs = "unbounded"/>
// 	<element name = "dataTransformacaoEntrada" type = "sngpc:st_data"/>
// </sequence>
// </complexType>
type EntradaMedicamentoTransformacao struct {
	XMLName                         xml.Name                   `xml:"entradaMedicamentoTransformacao"`
	MedicamentoTransformacaoEntrada []MedicamentoTransformacao `xml:"medicamentoTransformacaoEntrada"`
	DataTransformacaoEntrada        string                     `xml:"DataTransformacaoEntrada"`
}

//SaidaMedicamentoTransformacaoVendaAoConsumidor
// <complexType name = "ct_SaidaMedicamentoTransformacaoVendaAoConsumidor">
// <sequence>
// 	<element name = "tipoReceituarioMedicamento" type = "sngpc:st_TipoReceituario"/>
// 	<element name = "numeroNotificacaoMedicamento" type = "sngpc:st_Notificacao"/>
// 	<element name = "dataPrescricaoMedicamento" type = "sngpc:st_data"/>
// 	<element name = "prescritorMedicamento" type = "sngpc:ct_Prescritor"/>
// 	<element name = "usoMedicamento" type = "sngpc:st_TipoUsoMedicamento"/>
// 	<element name = "compradorMedicamento" type = "sngpc:ct_Comprador" minOccurs="0" maxOccurs="1"/>
// 	<element name = "pacienteMedicamento" type = "sngpc:ct_paciente" minOccurs="0" maxOccurs="1"/>
// 	<element name = "medicamentoVenda" type = "sngpc:ct_MedicamentoTransformacaoVenda" minOccurs = "0" maxOccurs = "unbounded"/>
// 	<element name = "dataVendaMedicamento" type = "sngpc:st_data"/>
// </sequence>
// </complexType>
type SaidaMedicamentoTransformacaoVendaAoConsumidor struct {
	XMLName                      xml.Name                        `xml:"saidaMedicamentoVendaAoConsumidor"`
	TipoReceituarioMedicamento   uint8                           `xml:"tipoReceituarioMedicamento"`
	NumeroNotificacaoMedicamento string                          `xml:"numeroNotificacaoMedicamento"`
	DataPrescricaoMedicamento    string                          `xml:"dataPrescricaoMedicamento"`
	PrescritorMedicamento        Prescritor                      `xml:"prescritorMedicamento"`
	UsoMedicamento               uint8                           `xml:"usoMedicamento"`
	CompradorMedicamento         Comprador                       `xml:"compradorMedicamento"`
	PacienteMedicamento          Paciente                        `xml:"pacienteMedicamento"`
	MedicamentoVenda             []MedicamentoTransformacaoVenda `xml:"medicamentoVenda"`
	DataVendaMedicamento         string                          `xml:"dataVendaMedicamento"`
}

//SaidaMedicamentoTransformacaoPerda
// <complexType name = "ct_SaidaMedicamentoTransformacaoPerda">
// <sequence>
// 	<!--<element name = "classeTerapeutica" type = "sngpc:st_classeTerapeutica"/>-->
// 	<element name = "motivoPerdaMedicamento" type = "sngpc:st_TipoMotivoPerda"/>
// 	<element name = "medicamentoPerda" type = "sngpc:ct_MedicamentoSaidaTransformacao"/>
// 	<element name = "dataPerdaMedicamento" type = "sngpc:st_data"/>
// </sequence>
// </complexType>
type SaidaMedicamentoTransformacaoPerda struct {
	ClasseTerapeutica      uint8                         `xml:"classeTerapeutica"`
	MotivoPerdaMedicamento uint8                         `xml:"motivoPerdaMedicamento"`
	MedicamentoPerda       MedicamentoSaidaTransformacao `xml:"medicamentoPerda"`
	DataPerdaMedicamento   string                        `xml:"dataPerdaMedicamento"`
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
	XMLName                  xml.Name `xml:"medicamentoVenda"`
	UsoProlongado            string   `xml:"usoProlongado"`
	RegistroMSMedicamento    string   `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string   `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    uint     `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento uint8    `xml:"unidadeMedidaMedicamento"`
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

//MedicamentoEntrada
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
	XMLName                  xml.Name `xml:"medicamentoEntrada"`
	ClasseTerapeutica        uint8    `xml:"classeTerapeutica"`
	RegistroMSMedicamento    string   `xml:"registroMSMedicamento"`
	NumeroLoteMedicamento    string   `xml:"numeroLoteMedicamento"`
	QuantidadeMedicamento    uint     `xml:"quantidadeMedicamento"`
	UnidadeMedidaMedicamento uint8    `xml:"unidadeMedidaMedicamento"`
}
