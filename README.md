# Documentação do Serviço de Envio de Mensagens

Este serviço permite o envio de mensagens SMS e WhatsApp para clientes que abandonaram carrinhos de compras em uma plataforma. Ele utiliza o framework Gin para criar endpoints HTTP e a API do Twilio para enviar mensagens.

## Endpoints

### Enviar SMS

- **URL:** `/sendSms`
- **Método:** POST
- **Descrição:** Este endpoint permite o envio de mensagens SMS para os clientes que abandonaram carrinhos de compras.
- **Corpo da Solicitação (JSON):**
  ```json
  {
    "email": "email-do-cliente",
    "phone": "número-de-telefone",
    "products": ["produto1", "produto2", "..."]
  }
  ```

### Enviar WhatsApp

- **URL:** `/sendWhats`
- **Método:** POST
- **Descrição:** Este endpoint permite o envio de mensagens WHATS para os clientes que abandonaram carrinhos de compras.
- **Corpo da Solicitação (JSON):**
  ```json
  {
    "email": "email-do-cliente",
    "phone": "número-de-telefone",
    "products": ["produto1", "produto2", "..."]
  }
  ```

## Configuração
Certifique-se de configurar as seguintes variáveis de ambiente antes de executar o serviço:

```
TWILIO_PHONE_NUMBER: Número de telefone Twilio para enviar SMS.
TWILIO_ACCOUNT_SID: SID da conta Twilio.
TWILIO_AUTH_TOKEN: Token de autenticação da conta Twilio.
```
Dependências
Este serviço depende das seguintes bibliotecas e pacotes:
```
Gin: Framework HTTP para Go.
Twilio Go: Biblioteca Twilio para Go.
Certifique-se de instalar essas dependências antes de executar o serviço.
```


## Como rodar:

```azure
go run .
```

## Configurar API:
````env
export TWILIO_ACCOUNT_SID=<SEU TOKEN AQUI>
export TWILIO_AUTH_TOKEN=<SEU TOKEN AQUI
export TWILIO_PHONE_NUMBER=<SEU TOKEN AQUI
````
