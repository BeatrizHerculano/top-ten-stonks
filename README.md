## Instalação e execução do projeto

Baixe o projeto no caminho usado pelo seu "go install"  e execute o seguinte comando:
```
$ go install top-ten-stonks
```

Ao executar esse comando um outro comando sera criado para a execução do projeto, que será:
```
$ top-ten-stonks
```

Esse projeto usa o driver oficial do MongoDB para Go e um scrapper para go chamado Colly, ambos são instalados ao ultilizar o comando "go install" para o projeto. 

## Configuração
Para conectar um banco de dados cole a URL para um base de dados MongoDB do atlas no arquivo main.go, onde le-se "your Atlas connection URI".

