# Labs Go Expert - Weather API
## Objetivo

Desenvolver um sistema em Go que receba um CEP, identifique a cidade e retorne o clima atual, apresentando as temperaturas em graus Celsius, Fahrenheit e Kelvin. Esse sistema será implementado e publicado no Google Cloud Run.

## Requisitos

- O sistema deve receber um **CEP** válido de 8 dígitos.
- O sistema deve realizar uma consulta para obter a localização com base no CEP fornecido e, a partir disso, recuperar as temperaturas atuais formatadas em:
  - **Celsius (C)**
  - **Fahrenheit (F)**
  - **Kelvin (K)**
  
### Respostas esperadas:

#### Em caso de sucesso:
- **Código HTTP**: 200
- **Corpo da resposta**:
  ```json
  {
    "temp_C": 28.5,
    "temp_F": 83.3,
    "temp_K": 301.5
  }

#### Em caso de falha (CEP inválido, mas no formato correto):
- **Código HTTP**: 422
- **Mensagem**: "invalid zipcode"

#### Em caso de falha (CEP não encontrado):
- **Código HTTP**: 404
- **Mensagem**: "can not find zipcode"

## Deploy

O sistema deverá ser publicado no Google Cloud Run.

## Dicas

- Utilize a API [viaCEP](https://viacep.com.br/) (ou similar) para encontrar a localização a partir do CEP fornecido.
- Utilize a [WeatherAPI](https://www.weatherapi.com/) (ou similar) para recuperar as informações de temperatura.
  
### Conversão de Temperatura
- **Celsius para Fahrenheit**: `F = C * 1.8 + 32`
- **Celsius para Kelvin**: `K = C + 273`

## Entrega

- Código-fonte completo da implementação.
- Testes automatizados que demonstrem o funcionamento correto do sistema.
- Utilize `docker`/`docker-compose` para facilitar a execução e os testes.
- Deploy realizado no Google Cloud Run (utilizando o *free tier*), com o endereço ativo para acesso público.