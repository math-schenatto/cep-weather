# API de Consulta de Temperatura por CEP

Este projeto é uma API desenvolvida em Go que recebe um CEP válido, consulta a localização correspondente e retorna a temperatura atual em Celsius, Fahrenheit e Kelvin. O sistema foi projetado para ser implantado no Google Cloud Run.

## Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Funcionalidades](#funcionalidades)
- [Requisitos](#requisitos)
- [Como Usar](#como-usar)
- [Instalação e Execução Local](#instalação-e-execução-local)
- [Deploy no Google Cloud Run](#deploy-no-google-cloud-run)
- [Testes Automatizados](#testes-automatizados)

## Sobre o Projeto

O objetivo deste projeto é fornecer uma API que recebe um CEP válido, consulta a localização correspondente e retorna a temperatura atual em três escalas: Celsius, Fahrenheit e Kelvin. A API foi desenvolvida em Go e utiliza as APIs [ViaCEP](https://viacep.com.br/) para consulta de CEP e [WeatherAPI](https://www.weatherapi.com/) para obter os dados de temperatura.

## Funcionalidades

- Recebe um CEP válido de 8 dígitos.
- Consulta a localização correspondente ao CEP.
- Retorna a temperatura atual em Celsius, Fahrenheit e Kelvin.
- Responde adequadamente em casos de sucesso ou falha.

## Requisitos

- Go 1.21 ou superior.
- Conta no Google Cloud para deploy no Cloud Run.
- Chave de API do [WeatherAPI](https://www.weatherapi.com/).

## Como Usar

### Endpoint

A API está disponível no seguinte endpoint: 
```bash
GET /weather?cep={CEP}
```

### Exemplo de Requisição

```bash
curl "https://cloud-run-cep-weather-gyvwutyyia-uc.a.run.app/weather?cep=13330-250"
```

### Resposta de Sucesso (HTTP 200)

```bash
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

### Respostas de Falha

- CEP inválido (HTTP 422):

```bash
{
  "message": "invalid zipcode"
}
```

- CEP não encontrado (HTTP 404):
```bash
{
  "message": "can not find zipcode"
}
```

## Instalação e execução local
O projeto pode ser executado localmente usando o docker-compose. Siga os passos abaixo:

Passos para Executar com Docker Compose

1. Clone o repositório:
```bash
git clone https://github.com/math-schenatto/cep-weather.git
cd cep-weather
```
2. Execute o projeto com Docker Compose:
```bash
docker-compose up --build
```

3. Acesse a API localmente:
```bash
curl "http://localhost:8080/weather?cep=13330-250"
```

## Deploy no Google Cloud Run

O projeto foi implantado no Google Cloud Run e está acessível através da seguinte URL:

```bash
https://cloud-run-cep-weather-gyvwutyyia-uc.a.run.app/weather?cep=13330-250
```

## Testes Automatizados
O projeto inclui testes automatizados para garantir o funcionamento correto da API. Para executar os testes, use o comando:

```bash
go test -timeout 30s -run ^TestWeatherHandler$ cep-weather/tests
```