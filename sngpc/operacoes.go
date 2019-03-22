package sngpc

// EntradaMedicamentos
type EntradaMedicamentos struct {
	NotaFiscalEntradaMedicamento NotaFiscal           `xml:"notaFiscalEntradaMedicamento"`
	MedicamentoEntrada           []MedicamentoEntrada `xml:"medicamentoEntrada"`
	DataRecebimentoMedicamento   string               `xml:"dataRecebimentoMedicamento"`
}

// SaidaMedicamentoVendaAoConsumidor
type SaidaMedicamentoVendaAoConsumidor struct {
	TipoReceituarioMedicamento   TipoReceituario    `xml:"tipoReceituarioMedicamento"`
	NumeroNotificacaoMedicamento string             `xml:"numeroNotificacaoMedicamento"`
	DataPrescricaoMedicamento    string             `xml:"dataPrescricaoMedicamento"`
	PrescritorMedicamento        Prescritor         `xml:"prescritorMedicamento"`
	UsoMedicamento               TipoUsoMedicamento `xml:"usoMedicamento"`
	CompradorMedicamento         Comprador          `xml:"compradorMedicamento"`
	PacienteMedicamento          Paciente           `xml:"pacienteMedicamento"`
	MedicamentoVenda             []MedicamentoVenda `xml:"medicamentoVenda"`
	DataVendaMedicamento         string             `xml:"dataVendaMedicamento"`
}

// SaidaMedicamentoTransferencia
type SaidaMedicamentoTransferencia struct {
	NotaFiscalTransferenciaMedicamento NotaFiscal    `xml:"notaFiscalTransferenciaMedicamento"`
	MedicamentoTransferencia           []Medicamento `xml:"medicamentoTransferencia"`
	DataTransferenciaMedicamento       string        `xml:"dataTransferenciaMedicamento"`
}

// SaidaMedicamentoPerda
type SaidaMedicamentoPerda struct {
	MotivoPerdaMedicamento uint8       `xml:"motivoPerdaMedicamento"`
	MedicamentoPerda       Medicamento `xml:"medicamentoPerda"`
	DataPerdaMedicamento   string      `xml:"dataPerdaMedicamento"`
}

// EntradaMedicamentoTransformacao
type EntradaMedicamentoTransformacao struct {
	MedicamentoTransformacaoEntrada []MedicamentoTransformacao `xml:"medicamentoTransformacaoEntrada"`
	DataTransformacaoEntrada        string                     `xml:"DataTransformacaoEntrada"`
}

// SaidaMedicamentoTransformacaoVendaAoConsumidor
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

// SaidaMedicamentoTransformacaoPerda
type SaidaMedicamentoTransformacaoPerda struct {
	ClasseTerapeutica      uint8                         `xml:"classeTerapeutica"`
	MotivoPerdaMedicamento uint8                         `xml:"motivoPerdaMedicamento"`
	MedicamentoPerda       MedicamentoSaidaTransformacao `xml:"medicamentoPerda"`
	DataPerdaMedicamento   string                        `xml:"dataPerdaMedicamento"`
}

// EntradaInsumo
type EntradaInsumo struct {
	NotaFiscalEntradaInsumo NotaFiscal            `xml:"notaFiscalEntradaInsumo"`
	InsumoEntrada           []InsumoBasicoEntrada `xml:"insumoEntrada"`
	DataRecebimentoInsumo   string                `xml:"dataRecebimentoInsumo"`
}

// SaidaInsumoVenda
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

// SaidaInsumoTransferencia
type SaidaInsumoTransferencia struct {
	NotaFiscalTransferenciaInsumo NotaFiscal     `xml:"notaFiscalTransferenciaInsumo"`
	InsumoTransferencia           []InsumoBasico `xml:"insumoTransferencia"`
	DataTransferenciaInsumo       string         `xml:"dataTransferenciaInsumo"`
}

// SaidaInsumoPerda
type SaidaInsumoPerda struct {
	MotivoPerdaInsumo     uint8       `xml:"motivoPerdaInsumo"`
	SubstanciaInsumoPerda InsumoPerda `xml:"substanciaInsumoPerda"`
	DataPerdaInsumo       string      `xml:"dataPerdaInsumo"`
	InsumoCNPJFornecedor  string      `xml:"insumoCNPJFornecedor"`
}

// MedicamentoInventario
type MedicamentoInvenatario struct {
	MedicamentoEntrada MedicamentoEntrada `xml:"medicamentoEntrada"`
}

// InsumoInventario
type InsumoInventario struct {
	InsumoEntrada InsumoBasicoEntrada `xml:"insumoEntrada"`
}
