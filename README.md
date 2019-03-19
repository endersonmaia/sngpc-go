[![Build Status](https://travis-ci.org/endersonmaia/sngpc-go.svg?branch=master)](https://travis-ci.org/endersonmaia/sngpc-go)

# sngpc-go

Este projeto tem por objetivo inicial criar um parser para o XML do SNGPC.

Com isto, será possível remover todos os dados que possam identificar pessoas e/ou empresas, e sintetizar apenas dados demográfios e estatísticos.

Tais dados podem ser utilizados para pesquisas acerca da incidência das prescrições de substâncias e correlacioná-las com outras fontes de dados.

Segue um exemplo de um arquivo [SNGPC](http://portal.anvisa.gov.br/sngpc/desenvolvedores) de acorco com o [Guia de Geração do XML – SNGPC Versão 2.0](http://www.anvisa.gov.br/sngpc/Documentos2012/Manual_SNGPC_2.0_2.pdf).

```xml
<?xml version="1.0" encoding="iso-8859-1" ?>
<mensagemSNGPCInventario xmlns="urn:sngpc-schema">
    <cabecalho>
        <cnpjEmissor>05059874000138</cnpjEmissor>
        <cpfTransmissor>72586648153</cpfTransmissor>
        <data>2006-09-30</data>
    </cabecalho>
    <corpo>
        <medicamentos>
            <entradaMedicamentos>
                <medicamentoEntrada>
                    <classeTerapeutica>1</classeTerapeutica>
                    <registroMSMedicamento>1888888888888</registroMSMedicamento>
                    <numeroLoteMedicamento>200678</numeroLoteMedicamento>
                    <quantidadeMedicamento>1234</quantidadeMedicamento>
                    <unidadeMedidaMedicamento>1</unidadeMedidaMedicamento>
                </medicamentoEntrada>
            </entradaMedicamentos>
            <entradaMedicamentos>
                <medicamentoEntrada>
                    <classeTerapeutica>1</classeTerapeutica>
                    <registroMSMedicamento>1888888888888</registroMSMedicamento>
                    <numeroLoteMedicamento>200678</numeroLoteMedicamento>
                    <quantidadeMedicamento>1234</quantidadeMedicamento>
                    <unidadeMedidaMedicamento>1</unidadeMedidaMedicamento>
                </medicamentoEntrada>
            </entradaMedicamentos>
        </medicamentos>
        <insumos>
            <entradaInsumos>
                <insumoEntrada>
                    <classeTerapeutica>1</classeTerapeutica>
                    <codigoInsumo>00092</codigoInsumo>
                    <numeroLoteInsumo>A315</numeroLoteInsumo>
                    <insumoCNPJFornecedor>99900099900000</insumoCNPJFornecedor>
                    <quantidadeInsumo>300000.0</quantidadeInsumo>
                    <tipoUnidade>1</tipoUnidade>
                </insumoEntrada>
            </entradaInsumos>
            <entradaInsumos>
                <insumoEntrada>
                    <classeTerapeutica>1</classeTerapeutica>
                    <codigoInsumo>00092</codigoInsumo>
                    <numeroLoteInsumo>A315</numeroLoteInsumo>
                    <insumoCNPJFornecedor>99900099900000</insumoCNPJFornecedor>
                    <quantidadeInsumo>300000.0</quantidadeInsumo>
                    <tipoUnidade>1</tipoUnidade>
                </insumoEntrada>
            </entradaInsumos>
        </insumos>
    </corpo>
</mensagemSNGPCInventario>
```


## Dados

É possível verificar os tipos de dados disponíveis nos arquivos XML através deste [schema XML](http://sngpc.anvisa.gov.br/schema/sngpcSimpleTypes.xsd).

O ideia é trazer os dados enumeráveis em formato CSV já sintetizados.
