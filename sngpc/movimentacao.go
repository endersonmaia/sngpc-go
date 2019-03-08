package sngpc

import "fmt"

//MensagemSNGPC
type MensagemSNGPC struct {
	Cabecalho CabecalhoMovimentacao `xml:"cabecalho"`
	Corpo     CorpoMovimentacao     `xml:"corpo"`
}

func (s MensagemSNGPC) String() string {
	return fmt.Sprintf("%v\n%v", s.Cabecalho, s.Corpo)
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
	CnpjEmissor    string `xml:"cnpjEmissor"`
	CpfTransmissor string `xml:"cpfTransmissor"`
	DataInicio     string `xml:"dataInicio"`
	DataFim        string `xml:"dataFim"`
}

func (s CabecalhoMovimentacao) String() string {
	return fmt.Sprintf("cnpj : %v ; cpf : %v ; inicio : %v, fim : %v\n", s.CnpjEmissor, s.CpfTransmissor, s.DataInicio, s.DataFim)
}

//Corpo
type CorpoMovimentacao struct {
	Medicamentos Medicamentos `xml:"medicamentos"`
	Insumos      Insumos      `xml:"insumos"`
}

func (s CorpoMovimentacao) String() string {
	out := "Corpo : \n"

	for _, s := range s.Medicamentos.EntradaMedicamentos {
		out += "EntradaMedicamentos : \n"
		out += fmt.Sprintf("NF : %v, Data: %v\n", s.NotaFiscalEntradaMedicamento.NumeroNotaFiscal, s.DataRecebimentoMedicamento)
		for _, m := range s.MedicamentoEntrada {
			out += fmt.Sprintf("Medicamento : %v, Lote : %v, Quantidade : %v, Classe : %v\n", m.RegistroMSMedicamento, m.NumeroLoteMedicamento, m.QuantidadeMedicamento, m.ClasseTerapeutica)
		}
	}

	for _, s := range s.Medicamentos.SaidaMedicamentoVendaAoConsumidor {
		out += "SaidaMedicamentoVendaAoConsumidor : \n"
		out += fmt.Sprintf("Data : %v\n", s.DataVendaMedicamento)
		for _, m := range s.MedicamentoVenda {
			out += fmt.Sprintf("Medicamento : %v, Lote : %v, Quantidade : %v\n", m.RegistroMSMedicamento, m.NumeroLoteMedicamento, m.QuantidadeMedicamento)
		}
	}
	return out
}

//Insumos
// <element name="insumos">
// <complexType>
//   <sequence>
// 	<element name="entradaInsumos" type="sngpc:ct_EntradaInsumo" minOccurs="0" maxOccurs="unbounded" />
// 	<element name="saidaInsumoVendaAoConsumidor" type="sngpc:ct_SaidaInsumoVenda" minOccurs="0" maxOccurs="unbounded" />
// 	<element name="saidaInsumoTransferencia" type="sngpc:ct_SaidaInsumoTransferencia" minOccurs="0" maxOccurs="unbounded" />
// 	<element name="saidaInsumoPerda" type="sngpc:ct_SaidaInsumoPerda" minOccurs="0" maxOccurs="unbounded" />
//   </sequence>
// </complexType>
// </element>
type Insumos struct {
	EntradaInsumos               []EntradaInsumo            `xml:"entradaInsumos"`
	SaidaInsumoVendaAoConsumidor []SaidaInsumoVenda         `xml:"saidaInsumoVendaAoConsumidor"`
	SaidaInsumoTransferencia     []SaidaInsumoTransferencia `xml:"saidaInsumoTransferencia"`
	SaidaInsumoPerda             []SaidaInsumoPerda         `xml:"saidaInsumoPerda"`
}

func (s Insumos) String() string {
	out := ""

	for _, e := range s.EntradaInsumos {
		out += fmt.Sprint(e.InsumoEntrada)
	}

	return out
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
