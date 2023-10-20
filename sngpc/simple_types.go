package sngpc

// ClasseTerapeutica
//
//	1: "Antimicrobiano"
//	2: "Sujeito a Controle Especial"
type ClasseTerapeutica uint8

var classeTerapeutica = map[ClasseTerapeutica]string{
	1: "Antimicrobiano",
	2: "Sujeito a Controle Especial",
}

// Retorna a descrição da classe terapêutica.
func (s ClasseTerapeutica) String() string {
	return classeTerapeutica[s]
}

// st_data
// st_CPF
// st_CNPJ
// st_RazaoSocial
// st_AES
// st_simNao
// st_SimNaoNull
// st_NumeroNotaFiscal
// st_QuantidadeMedicamento
// st_QuantidadeGramasMedicamento
// st_Notificacao
// st_RegistroMS
// st_CodigoDCB
// st_Nome
// st_NumeroDocumento
// st_Lote
// st_QuantidadeInsumo
// st_QuantidadeDeInsumoPorUnidadeFarmacotecnica
// st_QuantidadeDeUnidadesFarmacotecnicas

// TipoReceituario
//
//	1: "Receita de Controle Especial em 2 vias (Receita Branca)"
//	2: "Notificação de Receita B (Notificação Azul",
//	3: "Notificação de Receita Especial (Notificação Branca)",
//	4: "Notificação de Receita A (Notificação Amarela)",
//	5: "Receita Antimicrobiano em 2 vias",
type TipoReceituario uint8

var tipoReceituario = map[TipoReceituario]string{
	1: "Receita de Controle Especial em 2 vias (Receita Branca)",
	2: "Notificação de Receita B (Notificação Azul",
	3: "Notificação de Receita Especial (Notificação Branca)",
	4: "Notificação de Receita A (Notificação Amarela)",
	5: "Receita Antimicrobiano em 2 vias",
}

// Retorna a descrição do Tipo de Receituário.
func (s TipoReceituario) String() string {
	return tipoReceituario[s]
}

// TipoUsoMedicamento
//	1: "Humano"
// 	2: "Veterinário"
type TipoUsoMedicamento uint8

var tipoUsoMedicamento = map[TipoUsoMedicamento]string{
	1: "Humano",
	2: "Veterinário",
}

// Retorna a descrição do Tipo de Uso do Medicamento.
func (s TipoUsoMedicamento) String() string {
	return tipoUsoMedicamento[s]
}

// TipoOperacaoNotaFiscal
//	1: "Compra"
//	2: "Transferência"
//	3: "Venda"
type TipoOperacaoNotaFiscal uint8

var tipoOperacaoNotaFiscal = map[TipoOperacaoNotaFiscal]string{
	1: "Compra",
	2: "Transferência",
	3: "Venda",
}

// Retorna a descrição do Tipo de Operação da Nota Fiscal
func (s TipoOperacaoNotaFiscal) String() string {
	return tipoOperacaoNotaFiscal[s]
}

// ConselhoProfissional
//
//	"CRM" : "Conselho Regional de Medicina"
//	"CRMV": "Conselho Regional de Medicina Veterinária"
//	"CRO" : "Conselho Regional de Odontologia"
//	"CRF" : "Conselho Regional de Farmácia"
//	"RMS" : "Programa Mais Medicos - Registro Ministério da Saúde"
type ConselhoProfissional string

var conselhoProfissional = map[ConselhoProfissional]string{
	"CRM":  "Conselho Regional de Medicina",
	"CRMV": "Conselho Regional de Medicina Veterinária",
	"CRO":  "Conselho Regional de Odontologia",
	"CRF":  "Conselho Regional de Farmácia",
	"RMS":  "Programa Mais Medicos - Registro Ministério da Saúde",
}

// Retorna a descrição do Conselho Profissinal informado.
func (s ConselhoProfissional) String() string {
	return conselhoProfissional[s]
}

// UF - Unidade Federativa
//	"AC": "Acre"
//	"AL": "Alagoas"
//	"AM": "Amazonas"
//	"AP": "Amapa"
//	"BA": "Bahia"
//	"CE": "Ceara"
//	"DF": "Distrito Federal"
//	"ES": "Espirito Santo"
//	"GO": "Goias"
//	"MA": "Maranhao"
//	"MG": "Minas Gerais"
//	"MS": "Mato Grosso do Sul"
//	"MT": "Mato Grosso"
//	"PA": "Para"
//	"PB": "Paraiba"
//	"PE": "Pernambuco"
//	"PI": "Piaui"
//	"PR": "Parana"
//	"RJ": "Rio de Janeiro"
//	"RN": "Rio Grande do Norte"
//	"RO": "Rondonia"
//	"RR": "Roraima"
//	"RS": "Rio Grande do Sul"
//	"SC": "Santa Catarina"
//	"SE": "Sergipe"
//	"SP": "Sao Paulo"
//	"TO": "Tocantins"
type UF string

var uf = map[UF]string{
	"AC": "Acre",
	"AL": "Alagoas",
	"AM": "Amazonas",
	"AP": "Amapa",
	"BA": "Bahia",
	"CE": "Ceara",
	"DF": "Distrito Federal",
	"ES": "Espirito Santo",
	"GO": "Goias",
	"MA": "Maranhao",
	"MG": "Minas Gerais",
	"MS": "Mato Grosso do Sul",
	"MT": "Mato Grosso",
	"PA": "Para",
	"PB": "Paraiba",
	"PE": "Pernambuco",
	"PI": "Piaui",
	"PR": "Parana",
	"RJ": "Rio de Janeiro",
	"RN": "Rio Grande do Norte",
	"RO": "Rondonia",
	"RR": "Roraima",
	"RS": "Rio Grande do Sul",
	"SC": "Santa Catarina",
	"SE": "Sergipe",
	"SP": "Sao Paulo",
	"TO": "Tocantins",
}

// Retorna o nome da Unidade Federativa correspondente à sigla informada
func (s UF) String() string {
	return uf[s]
}

// TipoMotivoPerda
// 1:  "Furto / Roubo"
// 2:  "Avaria"
// 3:  "Vencimento"
// 4:  "Apreensão / Recolhimento pela Visa"
// 5:  "Perda no processo"
// 6:  "Coleta para controle de qualidade"
// 7:  "Perda de exclusão da portaria 344"
// 8:  "Por desvio de qualidade"
// 9:  "Recolhimento do Fabricante"
// 10: "Devolução ao fornecedor/distribuidora"
type TipoMotivoPerda uint8

var tipoMotivoPerda = map[TipoMotivoPerda]string{
	1:  "Furto / Roubo",
	2:  "Avaria",
	3:  "Vencimento",
	4:  "Apreensão / Recolhimento pela Visa",
	5:  "Perda no processo",
	6:  "Coleta para controle de qualidade",
	7:  "Perda de exclusão da portaria 344",
	8:  "Por desvio de qualidade",
	9:  "Recolhimento do Fabricante",
	10: "Devolução ao fornecedor/distribuidora",
}

// Retorna a descrição do Motivo da Perda
func (s TipoMotivoPerda) String() string {
	return tipoMotivoPerda[s]
}

// TipoUnidadeInsumo
//	1: "Grama"
//	2: "Mililitro"
//	3: "Unidade (U)"
type TipoUnidadeInsumo uint8

var tipoUnidadeInsumo = map[TipoUnidadeInsumo]string{
	1: "Grama",
	2: "Mililitro",
	3: "Unidade (U)",
}

// Retorna a descrição do Tipo de Unidade de Insumo
func (s TipoUnidadeInsumo) String() string {
	return tipoUnidadeInsumo[s]
}

// TipoUnidadeFarmacotecnica
//	1: "Grama"
//	2: "Cápsula"
//	3: "Comprimido"
//	4: "Mililitro"
type TipoUnidadeFarmacotecnica uint8

var tipoUnidadeFarmacotecnica = map[TipoUnidadeFarmacotecnica]string{
	1: "Grama",
	2: "Cápsula",
	3: "Comprimido",
	4: "Mililitro",
}

// Retorna a descrição do Tipo de Unidade Farmacotêcnica
func (s TipoUnidadeFarmacotecnica) String() string {
	return tipoUnidadeFarmacotecnica[s]
}

// TipoDocumento
//	1:  "CARTEIRA DE REGISTRO PROFISSIONAL"
//	2:  "CARTEIRA DE IDENTIDADE"
//	4:  "PEDIDO DE AUTORIZAÇÃO DE TRABALHO"
//	5:  "CERTIDÃO DE NASCIMENTO"
//	6:  "CERTIDÃO DE CASAMENTO"
//	7:  "CERTIFICADO DE RESERVISTA"
//	8:  "CARTA PATENTE"
//	10: "CERTIFICADO DE DISPENSA DE INCORPORAÇÃO"
//	11: "CARTEIRA DE IDENTIDADE DO ESTRANGEIRO"
//	13: "PASSAPORTE"
//	14: "PROTOCOLO DA POLÍCIA FEDERAL"
//	19: "INSCRIÇÃO ESTADUAL"
//	20: "INSCRIÇÃO MUNICIPAL"
//	21: "ALVARÁ/LICENÇA SANITÁRIA MUNICIPAL"
//	22: "ALVARA/LICENÇA SANITÁRIA ESTADUAL"
//	38: "AUTORIZAÇÃO DE FUNCIONAMENTO DE EMPRESA"
//	39: "AUTORIZAÇÃO ESPECIAL DE FUNCIONAMENTO"
//	40: "AUTORIZAÇÃO ESPECIAL SIMPLIFICADA"
//	50: "CARTEIRA DE TRABALHO E PREVIDÊNCIA SOCIAL"
//	62: "CADASTRO NACIONAL DE PESSOA JURIDICA"
type TipoDocumento uint8

var tipoDocumento = map[TipoDocumento]string{
	1:  "CARTEIRA DE REGISTRO PROFISSIONAL",
	2:  "CARTEIRA DE IDENTIDADE",
	4:  "PEDIDO DE AUTORIZAÇÃO DE TRABALHO",
	5:  "CERTIDÃO DE NASCIMENTO",
	6:  "CERTIDÃO DE CASAMENTO",
	7:  "CERTIFICADO DE RESERVISTA",
	8:  "CARTA PATENTE",
	10: "CERTIFICADO DE DISPENSA DE INCORPORAÇÃO",
	11: "CARTEIRA DE IDENTIDADE DO ESTRANGEIRO",
	13: "PASSAPORTE",
	14: "PROTOCOLO DA POLÍCIA FEDERAL",
	19: "INSCRIÇÃO ESTADUAL",
	20: "INSCRIÇÃO MUNICIPAL",
	21: "ALVARÁ/LICENÇA SANITÁRIA MUNICIPAL",
	22: "ALVARA/LICENÇA SANITÁRIA ESTADUAL",
	38: "AUTORIZAÇÃO DE FUNCIONAMENTO DE EMPRESA",
	39: "AUTORIZAÇÃO ESPECIAL DE FUNCIONAMENTO",
	40: "AUTORIZAÇÃO ESPECIAL SIMPLIFICADA",
	50: "CARTEIRA DE TRABALHO E PREVIDÊNCIA SOCIAL",
	62: "CADASTRO NACIONAL DE PESSOA JURIDICA",
}

// Retorna a descrição do Tipo do Documento
func (s TipoDocumento) String() string {
	return tipoDocumento[s]
}

// OrgaoExpedidor
//	"CRA":     "CONSELHO REGIONAL DE ADMINISTRAÇÃO"
//	"CRE":     "CONSELHO REGIONAL DE ECONOMIA"
//	"CREA":    "CONSELHO REG.DE ENG. ARQUIT. E AGRONOMIA"
//	"CRF":     "CONSELHO REGIONAL DE FARMÁCIA"
//	"DGPC":    "DIRETORIA GERAL DA POLÍCIA CIVIL"
//	"DPF":     "DEPARTAMENTO DE POLÍCIA FEDERAL"
//	"IDAMP":   "INSTITUTO IDENTIF. AROLDO MENDES PAIVA"
//	"IFP":     "INSTITUTO FÉLIX PACHECO"
//	"IN":      "IMPRENSA NACIONAL"
//	"JUNTA":   "JUNTA"
//	"MAER":    "MINISTÉRIO DA AERONÁUTICA"
//	"MEX":     "MINISTÉRIO DO EXÉRCITO"
//	"MM":      "MINISTÉRIO DA MARINHA"
//	"OAB":     "ORDEM DOS ADVOGADOS DO BRASIL"
//	"SEJSP":   "SECRETARIA DE EST. DA JUSTIÇA E SEG. PUB"
//	"SES":     "SECRETARIA DE ESTADO DA SEGURANÇA"
//	"SESP":    "SECRETARIA DO ESTADO SEG. PÚBLICA"
//	"SJS":     "SECRETARIA DA JUSTIÇA E DA SEGURANÇA"
//	"SJTC":    "SECR. DA JUST. DO TRAB. E DA CIDADANIA"
//	"SSIPT":   "SECR.  DE SEG. E INFORM. POLÍCIA TÉCNICA"
//	"SSP":     "SECRETARIA DE SEGURANÇA PÚBLICA"
//	"VACIV":   "VARA CIVIL"
//	"VAMEN":   "VARA DE MENORES"
//	"PM":      "POLICIA MILITAR"
//	"ITB":     "INSTITUTO TAVARES BURIL"
//	"CRM":     "CONSELHO REGIONAL DE MEDICINA"
//	"CBM":     "CORPO DE BOMBEIROS MILITAR"
//	"DIC":     "DETRAN - DIRETORIA DE IDENTIFICAÇÃO CIVIL"
//	"CFP":     "CONSELHO FEDERAL DE PSICOLOGIA"
//	"CRO":     "CONSELHO REGIONAL DE ODONTOLOGIA"
//	"COREN":   "CONSELHO REGIONAL DE ENFERMARIA"
//	"CFN":     "CONSELHO FEDERAL DE NUTRICIONISTAS"
//	"MRE":     "MINISTÉRIO DAS RELAÇÕES EXTERIORES"
//	"CRCI":    "CONSELHO REG. DE CORRETORES DE IMÓVEIS"
//	"CRB":     "CONSELHO REGIONAL DE BIOLOGIA"
//	"CRN":     "CONSELHO REGIONAL DE NUTRIÇÃO"
//	"CFE":     "CONSELHO FEDERAL DE ENFERMAGEM"
//	"CRC":     "CONSELHO REGIONAL DE CONTABILIDADE"
//	"CRP":     "CONSELHO REGIONAL DE PSICOLOGIA"
//	"CRQ":     "CONSELHO REGIONAL DE QUIMICA"
//	"ANVISA":  "AGÊNCIA NACIONAL DE VIGILÂNCIA SANITÁRIA"
//	"GOVEST":  "GOVERNO DO ESTADO"
//	"PREF":    "PREFEITURA"
//	"CRBM":    "CONSELHO REGIONAL DE BIOMEDICINA"
//	"IPF":     "INSTITUTO PEREIRA FAUSTINO"
//	"CREFITO": "CONSELHO REGIONAL DE FISIOTERAPIA E TERAPIA OCUPACIONAL"
//	"CRMV":    "CONSELHO REGIONAL DE MEDICINA VETERINÁRIA"
//	"MTE":     "MINISTÉRIO DO TRABALHO E EMPREGO"
//	"CRFA":    "CONSELHO REGIONAL DE FONOAUDIOLOGIA"
//	"CORECON": "CONSELHO REGIONAL DE ECONOMIA"
//	"SDS":     "SECRETARIA DE DESENVOLVIMENTO SOCIAL"
//	"SRF":     "SECRETARIA DA RECEITA FEDERAL"
type OrgaoExpedidor string

var orgaoExpedidor = map[OrgaoExpedidor]string{
	"CRA":     "CONSELHO REGIONAL DE ADMINISTRAÇÃO",
	"CRE":     "CONSELHO REGIONAL DE ECONOMIA",
	"CREA":    "CONSELHO REG.DE ENG. ARQUIT. E AGRONOMIA",
	"CRF":     "CONSELHO REGIONAL DE FARMÁCIA",
	"DGPC":    "DIRETORIA GERAL DA POLÍCIA CIVIL",
	"DPF":     "DEPARTAMENTO DE POLÍCIA FEDERAL",
	"IDAMP":   "INSTITUTO IDENTIF. AROLDO MENDES PAIVA",
	"IFP":     "INSTITUTO FÉLIX PACHECO",
	"IN":      "IMPRENSA NACIONAL",
	"JUNTA":   "JUNTA",
	"MAER":    "MINISTÉRIO DA AERONÁUTICA",
	"MEX":     "MINISTÉRIO DO EXÉRCITO",
	"MM":      "MINISTÉRIO DA MARINHA",
	"OAB":     "ORDEM DOS ADVOGADOS DO BRASIL",
	"SEJSP":   "SECRETARIA DE EST. DA JUSTIÇA E SEG. PUB",
	"SES":     "SECRETARIA DE ESTADO DA SEGURANÇA",
	"SESP":    "SECRETARIA DO ESTADO SEG. PÚBLICA",
	"SJS":     "SECRETARIA DA JUSTIÇA E DA SEGURANÇA",
	"SJTC":    "SECR. DA JUST. DO TRAB. E DA CIDADANIA",
	"SSIPT":   "SECR.  DE SEG. E INFORM. POLÍCIA TÉCNICA",
	"SSP":     "SECRETARIA DE SEGURANÇA PÚBLICA",
	"VACIV":   "VARA CIVIL",
	"VAMEN":   "VARA DE MENORES",
	"PM":      "POLICIA MILITAR",
	"ITB":     "INSTITUTO TAVARES BURIL",
	"CRM":     "CONSELHO REGIONAL DE MEDICINA",
	"CBM":     "CORPO DE BOMBEIROS MILITAR",
	"DIC":     "DETRAN - DIRETORIA DE IDENTIFICAÇÃO CIVIL",
	"CFP":     "CONSELHO FEDERAL DE PSICOLOGIA",
	"CRO":     "CONSELHO REGIONAL DE ODONTOLOGIA",
	"COREN":   "CONSELHO REGIONAL DE ENFERMARIA",
	"CFN":     "CONSELHO FEDERAL DE NUTRICIONISTAS",
	"MRE":     "MINISTÉRIO DAS RELAÇÕES EXTERIORES",
	"CRCI":    "CONSELHO REG. DE CORRETORES DE IMÓVEIS",
	"CRB":     "CONSELHO REGIONAL DE BIOLOGIA",
	"CRN":     "CONSELHO REGIONAL DE NUTRIÇÃO",
	"CFE":     "CONSELHO FEDERAL DE ENFERMAGEM",
	"CRC":     "CONSELHO REGIONAL DE CONTABILIDADE",
	"CRP":     "CONSELHO REGIONAL DE PSICOLOGIA",
	"CRQ":     "CONSELHO REGIONAL DE QUIMICA",
	"ANVISA":  "AGÊNCIA NACIONAL DE VIGILÂNCIA SANITÁRIA",
	"GOVEST":  "GOVERNO DO ESTADO",
	"PREF":    "PREFEITURA",
	"CRBM":    "CONSELHO REGIONAL DE BIOMEDICINA",
	"IPF":     "INSTITUTO PEREIRA FAUSTINO",
	"CREFITO": "CONSELHO REGIONAL DE FISIOTERAPIA E TERAPIA OCUPACIONAL",
	"CRMV":    "CONSELHO REGIONAL DE MEDICINA VETERINÁRIA",
	"MTE":     "MINISTÉRIO DO TRABALHO E EMPREGO",
	"CRFA":    "CONSELHO REGIONAL DE FONOAUDIOLOGIA",
	"CORECON": "CONSELHO REGIONAL DE ECONOMIA",
	"SDS":     "SECRETARIA DE DESENVOLVIMENTO SOCIAL",
	"SRF":     "SECRETARIA DA RECEITA FEDERAL",
}

// Retorna a descrição do Orgão Expedidor
func (s OrgaoExpedidor) String() string {
	return orgaoExpedidor[s]
}

// st_DescricaoValidacao

// st_Idade
type Idade uint8

// UnidadeIdade
//	1: "Anos"
//	2: "Meses
type UnidadeIdade uint8

var unidadeIdade = map[UnidadeIdade]string{
	1: "Anos",
	2: "Meses",
}

// Retorna a descrição da Unidade da Idade
func (s UnidadeIdade) String() string {
	return unidadeIdade[s]
}

// Sexo
//	1: "Masculino"
//	2: "Feminino"
type Sexo uint8

var sexo = map[Sexo]string{
	1: "Masculino",
	2: "Feminino",
}

// Retorna a descrição do Sexo
func (s Sexo) String() string {
	return sexo[s]
}

// st_Cid

// UnidadeMedidaMedicamento
//	1: "Caixas"
//	2: "Frascos"
type UnidadeMedidaMedicamento uint8

var unidadeMedidaMedicamento = map[UnidadeMedidaMedicamento]string{
	1: "Caixas",
	2: "Frascos",
}

// Retorna a descrição da Unidade de Medida do Medicamento
func (s UnidadeMedidaMedicamento) String() string {
	return unidadeMedidaMedicamento[s]
}
