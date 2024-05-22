# Testador de Carga em Go

Este projeto implementa um testador de carga simples em Go. Ele permite que você envie um número definido de requisições HTTP para uma URL específica e rastreie o tempo de resposta e o código de status de cada requisição.

## Como Usar
- Navegue até o diretório raiz do projeto e construa a imagem usando o comando:

```docker build -t nome-da-imagem .```

### Executar aplicaçao de teste de carga simples

- Use o seguinte comando para executar o testador de carga:

```docker run nome-da-imagem --url=<URL> --requests=<Número de Requisições> --concurrency=<Concorrência>```

- Substitua `<URL>` pela URL do serviço que você deseja testar.
- Substitua `<Número de Requisições>` pelo total de requisições que você deseja realizar.
- Substitua `<Concorrência>` pelo número de requisições simultâneas.

## Resultados Esperados

Após a execução do teste de carga, o programa fornecerá um relatório no terminal, que incluirá:

- **Tempo Médio de Requisição**: O tempo médio levado para cada requisição em segundos.
- **Quantidade Total de Requisições**: O número total de requisições realizadas.
- **Distribuição dos Códigos de Status**: Uma contagem de quantas vezes cada código de status HTTP foi retornado.

## Exemplo de Comando

```docker run stress-test --url=https://fullcycle.com.br --requests=1000 --concurrency=100```
