package sngpc

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func (s MensagemSNGPC) String() string {
	return fmt.Sprintf("%v\n%v", s.Cabecalho, s.Corpo)
}

func (s CabecalhoMovimentacao) String() string {
	return fmt.Sprintf("cnpj : %v ; cpf : %v ; inicio : %v, fim : %v\n", s.CnpjEmissor, s.CpfTransmissor, s.DataInicio, s.DataFim)
}

func (s CorpoMovimentacao) String() string {
	text := `Corpo : 
{{ range .Medicamentos.EntradaMedicamentos -}}
EntradaMedicamentos : 
NF : {{ .NotaFiscalEntradaMedicamento.NumeroNotaFiscal }}, Data: {{ .DataRecebimentoMedicamento }}
{{ range .MedicamentoEntrada -}}
Medicamento : {{ .RegistroMSMedicamento }}, Lote : {{ .NumeroLoteMedicamento }}, Quantidade : {{ .QuantidadeMedicamento }}
{{ end }}
{{ end }}
{{- range .Medicamentos.SaidaMedicamentoVendaAoConsumidor }}
SaidaMedicamentoVendaAoConsumidor : 
Data : {{ .DataVendaMedicamento }}
{{ range .MedicamentoVenda -}}
Medicamento : {{ .RegistroMSMedicamento }}, Lote : {{ .NumeroLoteMedicamento }}, Quantidade : {{ .QuantidadeMedicamento }}
{{ end }}
{{ end }}`
	tmpl := template.Must(template.New("").Parse(text))
	var out bytes.Buffer
	tmpl.Execute(&out, s)

	return out.String()
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
	var pieces []string

	for _, e := range s.EntradaMedicamentos {
		pieces = append(pieces, fmt.Sprint(e.MedicamentoEntrada))
	}

	return strings.Join(pieces, "")
}

func (s Insumos) String() string {
	var pieces []string

	for _, e := range s.EntradaInsumos {
		pieces = append(pieces, e.InsumoEntrada.String())
	}

	return strings.Join(pieces, "")
}

func (s MedicamentoEntrada) String() string {
	return fmt.Sprintf("RegistroMS : %v, Lote : %v, Quantidade : %v\n", s.RegistroMSMedicamento, s.NumeroLoteMedicamento, s.QuantidadeMedicamento)
}

func (s InsumoEntrada) String() string {
	return fmt.Sprintf("RegistroMS : %v, Lote : %v, Quantidade : %v\n", s.CodigoInsumo, s.NumeroLoteInsumo, s.QuantidadeInsumo)
}
