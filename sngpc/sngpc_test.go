package sngpc

import (
	"strings"
	"testing"
)

var sngpcEntrada = `
<?xml version="1.0" encoding="ISO-8859-1"?>
<mensagemSNGPC xmlns="urn:sngpc-schema">
	<cabecalho>
		<cnpjEmissor>00000000000000</cnpjEmissor>
		<cpfTransmissor>00000000000</cpfTransmissor>
		<dataInicio>2019-01-01</dataInicio>
		<dataFim>2019-01-01</dataFim>
	</cabecalho>
	<corpo>
		<medicamentos>
			<entradaMedicamentos>
			<notaFiscalEntradaMedicamento>
				<numeroNotaFiscal>123456</numeroNotaFiscal>
				<tipoOperacaoNotaFiscal>1</tipoOperacaoNotaFiscal>
				<dataNotaFiscal>2019-01-01</dataNotaFiscal>
				<cnpjOrigem>00000000000000</cnpjOrigem>
				<cnpjDestino>00000000000000</cnpjDestino>
			</notaFiscalEntradaMedicamento>
			<medicamentoEntrada>
				<classeTerapeutica>1</classeTerapeutica>
				<registroMSMedicamento>1023503520122</registroMSMedicamento>
				<numeroLoteMedicamento>LOTE001</numeroLoteMedicamento>
				<quantidadeMedicamento>6</quantidadeMedicamento>
				<unidadeMedidaMedicamento>1</unidadeMedidaMedicamento>
			</medicamentoEntrada>
		</medicamentos>
	</corpo>
</mensagemSNGPC>
`

func TestFromXML(t *testing.T) {
	s := MensagemSNGPC{}
	fromXML(strings.NewReader(sngpcEntrada), &s)

	if s.Cabecalho.CnpjEmissor != "00000000000000" {
		t.Errorf("MensagemSNGPC.Cabecalho.CnpjEmissor errado, got: %v, want: %v.", s.Cabecalho.CnpjEmissor, "00000000000000")
	}
}
