package sngpc

import "fmt"

//MensagemSNGPCInventario
type MensagemSNGPCInventario struct {
	Cabecalho CabecalhoInventario `xml:"cabecalho"`
	Corpo     CorpoInventario     `xml:"corpo"`
}

//CabecalhoInventario
type CabecalhoInventario struct {
	CnpjEmissor    string `xml:"cnpjEmissor"`
	CpfTransmissor string `xml:"cpfTransmissor"`
	Data           string `xml:"data"`
}

//Corpo
type CorpoInventario struct {
	Medicamentos Medicamentos `xml:"medicamentos"`
	Insumos      Insumos      `xml:"insumos"`
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
