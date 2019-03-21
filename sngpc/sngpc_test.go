package sngpc

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestFromXML(t *testing.T) {
	is := is.New(t)

	s := MensagemSNGPC{}
	err := fromXML(strings.NewReader(msgSNGPCEntrada), &s)

	is.NoErr(err)
	is.Equal(s.Cabecalho.CnpjEmissor, "00000000000000")
	is.Equal(len(s.Corpo.Medicamentos.EntradaMedicamentos[0].MedicamentoEntrada), 2)
	is.Equal(s.Corpo.Medicamentos.EntradaMedicamentos[0].DataRecebimentoMedicamento, `2019-01-01`)
}

func TestWrongSNGPCType(t *testing.T) {
	is := is.New(t)

	s := MensagemSNGPCInventario{}
	err := fromXML(strings.NewReader(msgSNGPCEntrada), &s)
	is.True(err != nil) // shold error when trying do load MensagemSNGPCInventario{} from a MensagemSNGPC type XML
}

var msgSNGPCEntrada = `
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
				<medicamentoEntrada>
					<classeTerapeutica>1</classeTerapeutica>
					<registroMSMedicamento>1023503520123</registroMSMedicamento>
					<numeroLoteMedicamento>LOTE002</numeroLoteMedicamento>
					<quantidadeMedicamento>6</quantidadeMedicamento>
					<unidadeMedidaMedicamento>1</unidadeMedidaMedicamento>
				</medicamentoEntrada>
				<dataRecebimentoMedicamento>2019-01-01</dataRecebimentoMedicamento>
			</entradaMedicamentos>
		</medicamentos>
	</corpo>
</mensagemSNGPC>
`
