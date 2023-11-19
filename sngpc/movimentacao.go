package sngpc

import (
	"encoding/xml"
	"fmt"
)

// MensagemSNGPC armazena o conteúdo de um arquivo de movimentação do SNGPC
// O arquivo contem informações da movimentação de medicamentos e insumos
// de um determinado período
type MensagemSNGPC struct {
	XMLName   xml.Name              `xml:"mensagemSNGPC"`
	Cabecalho CabecalhoMovimentacao `xml:"cabecalho"`
	Corpo     CorpoMovimentacao     `xml:"corpo"`
}

func (s MensagemSNGPC) String() string {
	return fmt.Sprintf("%v\n%v", s.Cabecalho, s.Corpo)
}

// CabecalhoMovimentacao armazena informações da movimentação
type CabecalhoMovimentacao struct {
	CnpjEmissor    string `xml:"cnpjEmissor"`
	CpfTransmissor string `xml:"cpfTransmissor"`
	DataInicio     string `xml:"dataInicio"`
	DataFim        string `xml:"dataFim"`
}

func (s CabecalhoMovimentacao) String() string {
	return fmt.Sprintf("cnpj : %v ; cpf : %v ; inicio : %v, fim : %v\n", s.CnpjEmissor, s.CpfTransmissor, s.DataInicio, s.DataFim)
}

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

// Insumos armazena os diversos tipos de movimentações de insumos
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

func (s Medicamentos) String() string {
	out := ""

	for _, e := range s.EntradaMedicamentos {
		out += fmt.Sprint(e.MedicamentoEntrada)
	}

	return out
}
