# Sistema de temperatura por CEP

Desafio do Lab para Pós graduação em GoLang - Fullcycle

## Informações para rodar localmente

Para rodar a aplicação, basta executar:

```
docker-compose up -d
```

A Aplicação estará disponível em:

```
http://localhost:8080
```

Para ter acesso a rota raiz, basta colocar ao CEP corretamente da seguinte forma, substituindo o {cep} pelo número do cep, sem espaços, hifem ou pontuação: http://localhost:8080/{cep}

Exemplo:

```
http://localhost:8080/45208643
```

## Informações do deploy no Google Cloud

Para executar a aplicação em ambiente web, basta acessar o endereço, substituindo o {cep} pelo número do cep, sem espaços, hifem ou pontuação:
https://desafio-05-lab-temperatura-ex52perbeq-uc.a.run.app/{cep}

Exemplo:

```
https://desafio-05-lab-temperatura-ex52perbeq-uc.a.run.app/45208643
```

## Testes

Os testes podem ser rodados a partir do comando: go test -v

Exemplo:

```
go test -v
```
