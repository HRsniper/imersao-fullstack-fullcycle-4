# Imersão Full Stack & FullCycle 4.0 - Fincycle - Back-end das transações

## Descrição

Repositório do back-end das transações feito com Nest.js

**Importante**: A aplicação do Apache Kafka, Keycloak deve estar rodando primeiro.

## Rodar a aplicação

### Configurar /etc/hosts

A comunicação entre as aplicações se dá de forma direta através da rede da máquina.
Para isto é necessário configurar um endereços que todos os containers Docker consigam acessar.

Acrescente no seu /etc/hosts (para Windows o caminho é C:\Windows\system32\drivers\etc\hosts):

```
127.0.0.1 host.docker.internal tenant1.docker.internal tenant2.docker.internal
```

Em todos os sistemas operacionais é necessário abrir o programa para editar o _hosts_ como Administrator da máquina ou root.

Execute os comandos:

```
docker-compose up
```

Acessar http://localhost:3000.
