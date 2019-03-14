package sngpc

//ClasseTerapeutica
// <simpleType name="st_classeTerapeutica">
// <restriction base="string">
//   <!--Antimicrobiano-->
//   <enumeration value="1" />
//   <!--Sujeito a controle Especial-->
//   <enumeration value="2" />
// </restriction>
// </simpleType>
type ClasseTerapeutica uint8

const (
	Antimicrobiano           ClasseTerapeutica = 1
	SujeitoAControleEspecial ClasseTerapeutica = 2
)

type UnidadeMedidaMedicamento uint8

const (
	Caixas  UnidadeMedidaMedicamento = 1
	Frascos UnidadeMedidaMedicamento = 2
)

type TipoUnidadeInsumo uint8

const (
	Grama     TipoUnidadeInsumo = 1
	Mililitro TipoUnidadeInsumo = 2
	Unidade   TipoUnidadeInsumo = 3
)
