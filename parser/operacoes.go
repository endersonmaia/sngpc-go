package parser

//EntradaMedicamentos
// <complexType name = "ct_EntradaMedicamento">
// <sequence>
// 	<element name = "notaFiscalEntradaMedicamento" type = "sngpc:ct_NotaFiscal"/>
// 	<element name = "medicamentoEntrada" type = "sngpc:ct_MedicamentoEntrada" minOccurs = "0" maxOccurs = "unbounded"/>
// 	<element name = "dataRecebimentoMedicamento" type = "sngpc:st_data"/>
// </sequence>
// </complexType>
type EntradaMedicamentos struct {
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

//EntradaInsumo
// <complexType name = "ct_EntradaInsumo">
// 	<sequence>
// 		<element name = "notaFiscalEntradaInsumo" type = "sngpc:ct_NotaFiscal"/>
// 		<element name = "insumoEntrada" type = "sngpc:ct_InsumoBasicoEntrada" minOccurs = "0" maxOccurs = "unbounded" />
// 		<element name = "dataRecebimentoInsumo" type = "sngpc:st_data"/>
// 	</sequence>
// </complexType>
type EntradaInsumo struct {
	NotaFiscalEntradaInsumo NotaFiscal            `xml:"notaFiscalEntradaInsumo"`
	InsumoEntrada           []InsumoBasicoEntrada `xml:"insumoEntrada"`
	dataRecebimentoInsumo   string                `xml:"dataRecebimentoInsumo"`
}

// SaidaInsumoVenda
// <complexType name = "ct_SaidaInsumoVenda">
// <sequence>
// 	<element name = "tipoReceituarioInsumo" type = "sngpc:st_TipoReceituario"/>
// 	<element name = "numeroNotificacaoInsumo" type = "sngpc:st_Notificacao"/>
// 	<element name = "dataPrescricaoInsumo" type = "sngpc:st_data"/>
// 	<element name = "prescritorInsumo" type = "sngpc:ct_Prescritor"/>
// 	<element name = "usoInsumo" type = "sngpc:st_TipoUsoMedicamento"/>
// 	<element name = "compradorInsumo" type = "sngpc:ct_Comprador" minOccurs="0" maxOccurs="1"/>
// 	<element name = "pacienteInsumo" type = "sngpc:ct_paciente" minOccurs="0" maxOccurs="1"/>
// 	<element name = "substanciaInsumoVendaAoConsumidor" type = "sngpc:ct_InsumoVendaAoConsumidor" minOccurs = "0" maxOccurs = "unbounded"/>
// 	<element name = "dataVendaInsumo" type = "sngpc:st_data"/>
// </sequence>
// </complexType>
type SaidaInsumoVenda struct {
	TipoReceituarioInsumo             uint8                   `xml:"tipoReceituarioInsumo"`
	NumeroNotificacaoInsumo           string                  `xml:"numeroNotificacaoInsumo"`
	DataPrescricaoInsumo              string                  `xml:"dataPrescricaoInsumo"`
	PrescritorInsumo                  string                  `xml:"prescritorInsumo"`
	UsoInsumo                         string                  `xml:"usoInsumo"`
	CompradorInsumo                   Comprador               `xml:"compradorInsumo"`
	PacienteInsumo                    Paciente                `xml:"pacienteInsumo"`
	SubstanciaInsumoVendaAoConsumidor InsumoVendaAoConsumidor `xml:"substanciaInsumoVendaAoConsumidor"`
	DataVendaInsumo                   string                  `xml:"dataVendaInsumo"`
}

//SaidaInsumoTransferencia
// <complexType name = "ct_SaidaInsumoTransferencia">
// <sequence>
//   <element name = "notaFiscalTransferenciaInsumo" type = "sngpc:ct_NotaFiscal"/>
//   <element name = "insumoTransferencia" type = "sngpc:ct_InsumoBasico" minOccurs="0" maxOccurs="unbounded" />
//   <element name = "dataTransferenciaInsumo" type = "sngpc:st_data"/>
// </sequence>
// </complexType>
type SaidaInsumoTransferencia struct {
	NotaFiscalTransferenciaInsumo NotaFiscal     `xml:"notaFiscalTransferenciaInsumo"`
	InsumoTransferencia           []InsumoBasico `xml:"insumoTransferencia"`
	DataTransferenciaInsumo       string         `xml:"dataTransferenciaInsumo"`
}

//SaidaInsumoPerda
// <complexType name = "ct_SaidaInsumoPerda">
// <sequence>
// 	<element name = "motivoPerdaInsumo" type = "sngpc:st_TipoMotivoPerda"/>
// 	<element name = "substanciaInsumoPerda" type = "sngpc:ct_InsumoPerda"/>
// 	<element name = "dataPerdaInsumo" type = "sngpc:st_data"/>
// 	<element name = "insumoCNPJFornecedor" type = "sngpc:st_CNPJ"/>
// </sequence>
// </complexType>
type SaidaInsumoPerda struct {
	MotivoPerdaInsumo     uint8       `xml:"motivoPerdaInsumo"`
	SubstanciaInsumoPerda InsumoPerda `xml:"substanciaInsumoPerda"`
	DataPerdaInsumo       string      `xml:"dataPerdaInsumo"`
	InsumoCNPJFornecedor  string      `xml:"insumoCNPJFornecedor"`
}

//MedicamentoInventario
// <complexType name = "ct_Medicamento_Inventario">
// <sequence>
// 	<element name = "medicamentoEntrada" type = "sngpc:ct_MedicamentoEntrada" />
// </sequence>
// </complexType>
type MedicamentoInvenatario struct {
	MedicamentoEntrada MedicamentoEntrada `xml:"medicamentoEntrada"`
}

//InsumoInventario
// <complexType name = "ct_Insumo_Inventario">
// <sequence>
// 	<element name = "insumoEntrada" type = "sngpc:ct_InsumoBasicoEntrada"/>
// </sequence>
// </complexType>
type InsumoInventario struct {
	InsumoEntrada InsumoBasicoEntrada `xml:"insumoEntrada"`
}
