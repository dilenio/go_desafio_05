# Desafio 05 - Stress Test

Para buildar a aplicação use o comando docker abaixo:

```
docker build -t loadtester .
```

Para rodar a aplicação, use um dos comandos abaixo:

```
go run main.go --url=https://google.com --requests=100 --concurrency=20
```

ou

```
docker run loadtester --url=https://google.com --requests=100 --concurrency=20
```
