## Fase 2

### Pré requisitos:
* kubectl v1.29.0  
* minikube v1.32.0  
* go 1.21.3  

### Como implementar:
Executar o script **buildAndDeploy_FASE2.sh**

Ao executar o script, será deletado o cluster existente do Minikube e criado um novo, ajuste nas permissões do docker.sock, habilita o metrics-server, realiza o deploy da aplicação no cluster e configuração do HPA.

Importante aguardar alguns minutos para conclusão da criação dos pods antes de testar.

### Como testar:
Na pasta raiz existe o arquivo **Testes-FASE-2.postman_collection.json** com chamadas prontas para todas as APIs.

### Vídeo demonstração - Deploy e testes:
[Tech Challenge Fase 2](https://youtu.be/3eS7t2aHkI4)

### Descrição das APIs envolvidas na Fase 2:

`/pedido/novo-pedido`  
_Checkout do pedido retornando um código de identificação_

`/pagamento/confirmacao`  
_Atualiza o status de pagamento e em caso de sucesso no pagamento, o status do pedido passa para "Em preparação"_

`/pedido/consulta-pagamento?codigopedido={codigopedido}`  
_Consulta o status do pagamento de um pedido_

`/listapedidos/todos`  
_Lista todos os pedidos realizados_

---
---

## Fase 1

### Como implementar:
Executar o script **buildAndDeploy_FASE1.sh** ou **docker compose build && docker compose up**.

### Como testar:
Na pasta raiz existe o arquivo **Testes-FASE-1.postman_collection.json** com chamadas prontas para todas as APIS.

**Obs.:** Na API de confirmação de pagamento, é necessário que seja informado o código do pedido. Para isso, basta copiar o **codPedido** gerado na API /pedido/novo-pedido e colar no JSON da API /pagamento/confirmação.

### Documentação DDD - Miro:
https://miro.com/app/board/uXjVNZqJmc0=/?share_link_id=344727336283

---

Wellington de Souza Pires - 4SOAT - RM350887
