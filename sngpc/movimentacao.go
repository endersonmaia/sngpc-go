package sngpc

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"log"
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

// VendasToCSV formata a operação SaidaMedicamentoVendaAoConsumidor
func (s MensagemSNGPC) VendasToCSV(w *csv.Writer) error {
	header := []string{
		"tipoReceituarioMedicamento",
		"numeroNotificacaoMedicamento",
		"dataPrescricaoMedicamento",
		"nomePrescritor",
		"numeroRegistroProfissionalPrescritor",
		"conselhoProfissionalPrescritor",
		"UFConselhoPrescritor",
		"usoMedicamento",
		"nomeComprador",
		"tipoDocumentoComprador",
		"numeroDocumentoComprador",
		"orgaoExpedidorComprador",
		"UFEmissaoDocumentoComprador",
		"nomePaciente",
		"idadePaciente",
		"unidadeIdadePaciente",
		"sexoPaciente",
		"cidPaciente",
		"usoProlongadoMedicamento",
		"registroMSMedicamento",
		"numeroLoteMedicamento",
		"quantidadeMedicamento",
		"unidadeMedidaMedicamento",
		"dataVendaMedicamento",
	}

	err := w.Write(header)
	if err != nil {
		log.Panic(err)
		return err
	}

	for _, v := range s.Corpo.Medicamentos.SaidaMedicamentoVendaAoConsumidor {
		for _, m := range v.MedicamentoVenda {
			r := []string{
				v.TipoReceituarioMedicamento.String(),
				v.NumeroNotificacaoMedicamento,
				v.DataPrescricaoMedicamento,
				v.PrescritorMedicamento.NomePrescritor,
				v.PrescritorMedicamento.NumeroRegistroProfissional,
				v.PrescritorMedicamento.ConselhoProfissional.String(),
				v.PrescritorMedicamento.UFConselho.String(),
				v.UsoMedicamento.String(),
				v.CompradorMedicamento.NomeComprador,
				v.CompradorMedicamento.TipoDocumento.String(),
				v.CompradorMedicamento.NumeroDocumento,
				v.CompradorMedicamento.UFEmissaoDocumento.String(),
				v.PacienteMedicamento.Nome,
				v.PacienteMedicamento.Idade,
				v.PacienteMedicamento.UnidadeIdade.String(),
				v.PacienteMedicamento.Sexo.String(),
				v.PacienteMedicamento.Cid,
				m.UsoProlongado,
				m.RegistroMSMedicamento,
				m.NumeroLoteMedicamento,
				string(m.QuantidadeMedicamento),
				m.UnidadeMedidaMedicamento.String(),
				v.DataVendaMedicamento,
			}

			w.Write(r)
		}
	}

	w.Flush()

	return nil
}

// VendasToCSVAnonymized formata a operação SaidaMedicamentoVendaAoConsumidor
func (s MensagemSNGPC) VendasToCSVAnonymized(w *csv.Writer) error {
	header := []string{
		"tipoReceituarioMedicamento",
		"numeroNotificacaoMedicamento",
		"dataPrescricaoMedicamento",
		"conselhoProfissionalPrescritor",
		"UFConselhoPrescritor",
		"usoMedicamento",
		"UFEmissaoDocumentoComprador",
		"idadePaciente",
		"unidadeIdadePaciente",
		"sexoPaciente",
		"cidPaciente",
		"usoProlongadoMedicamento",
		"registroMSMedicamento",
		"numeroLoteMedicamento",
		"quantidadeMedicamento",
		"unidadeMedidaMedicamento",
		"dataVendaMedicamento",
	}

	err := w.Write(header)
	if err != nil {
		log.Panic(err)
		return err
	}

	for _, v := range s.Corpo.Medicamentos.SaidaMedicamentoVendaAoConsumidor {
		for _, m := range v.MedicamentoVenda {
			r := []string{
				v.TipoReceituarioMedicamento.String(),
				v.NumeroNotificacaoMedicamento,
				v.DataPrescricaoMedicamento,
				v.PrescritorMedicamento.ConselhoProfissional.String(),
				v.PrescritorMedicamento.UFConselho.String(),
				v.UsoMedicamento.String(),
				v.CompradorMedicamento.UFEmissaoDocumento.String(),
				v.PacienteMedicamento.Idade,
				v.PacienteMedicamento.UnidadeIdade.String(),
				v.PacienteMedicamento.Sexo.String(),
				v.PacienteMedicamento.Cid,
				m.UsoProlongado,
				m.RegistroMSMedicamento,
				m.NumeroLoteMedicamento,
				string(m.QuantidadeMedicamento),
				m.UnidadeMedidaMedicamento.String(),
				v.DataVendaMedicamento,
			}

			w.Write(r)
		}
	}

	w.Flush()

	return nil
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
