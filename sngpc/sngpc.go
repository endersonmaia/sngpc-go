package sngpc

import (
	"fmt"
)

func (s MensagemSNGPC) String() string {
	return fmt.Sprintf("%v\n%v", s.Cabecalho, s.Corpo)
}

func (s CabecalhoMovimentacao) String() string {
	return fmt.Sprintf("cnpj : %v ; cpf : %v ; inicio : %v, fim : %v\n", s.CnpjEmissor, s.CpfTransmissor, s.DataInicio, s.DataFim)
}

func (s CorpoMovimentacao) String() string {
	out := "Corpo : \n"

	for _, s := range s.Medicamentos.EntradaMedicamentos {
		out += "EntradaMedicamentos : \n"
		out += fmt.Sprintf("NF : %v, Data: %v\n", s.NotaFiscalEntradaMedicamento.NumeroNotaFiscal, s.DataRecebimentoMedicamento)
		for _, m := range s.MedicamentoEntrada {
			out += fmt.Sprintf("Medicamento : %v, Lote : %v, Quantidade : %v\n", m.RegistroMSMedicamento, m.NumeroLoteMedicamento, m.QuantidadeMedicamento)
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

//Inventario
func (s MensagemSNGPCInventario) String() string {
	return fmt.Sprintf("%v\n%v", s.Cabecalho, s.Corpo)
}

func (s CabecalhoInventario) String() string {
	return fmt.Sprintf("Inventário :\n\nCNPJ : %v ; CPF : %v ; Data : %v\n", s.CnpjEmissor, s.CpfTransmissor, s.Data)
}

func (s CorpoInventario) String() string {
	return fmt.Sprintf("Corpo : \n\nMedicamentos : \n%v\nInsumos : \n%v\n", s.Medicamentos, s.Insumos)
}

func (s Medicamentos) String() string {
	out := ""

	for _, e := range s.EntradaMedicamentos {
		out += fmt.Sprint(e.MedicamentoEntrada)
	}

	return out
}

func (s Insumos) String() string {
	out := ""

	for _, e := range s.EntradaInsumos {
		out += fmt.Sprint(e.InsumoEntrada)
	}

	return out
}

func (s MedicamentoEntrada) String() string {
	return fmt.Sprintf("RegistroMS : %v, Lote : %v, Quantidade : %v\n", s.RegistroMSMedicamento, s.NumeroLoteMedicamento, s.QuantidadeMedicamento)
}

func (s InsumoEntrada) String() string {
	return fmt.Sprintf("RegistroMS : %v, Lote : %v, Quantidade : %v\n", s.CodigoInsumo, s.NumeroLoteInsumo, s.QuantidadeInsumo)
}
