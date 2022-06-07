# Busca de IPs e Servidores de um endereço web </h1>

![Versão](https://img.shields.io/badge/Vers%C3%A3o-1.0.0-lightgrey) 
![Status](https://img.shields.io/badge/status-otimizando%2Ffinalizado-success)
![Linguagem](https://img.shields.io/badge/language-golang-informational)

- Recebe um endereço web e busca na internet o ip público e o servidor onde esse endereço está hospedado.
- Caso não seja passada nenhuma flag de busca, retornará de um endereço padrão (google.com)

## 🛠️ Como rodar o projeto

- Clone o projeto com o comando: `git clone https://github.com/fabianorauzer-egsys/golang-find-ip-servers` <br>
- É necessário ter o Go instalado na máquina. 
Seguir esse <a href="https://wiki.egsys.com.br/desenvolvimento/como-instalar-go" target="blank"> passo a passo</a>.
- Dentro do projeto, execute o comando `go run main.go ip`, se tudo estiver correto, deverá ter um retorno dos IPs Google, que foi setado como padrão nesse projeto.
- Caso dê algum outro tipo de retorno, deve estar faltando alguma dependência do Go na sua máquina, mas o próprio terminal deverá informar os comandos a serem executados nesse caso.
- A síntaxe do comando para realizar a busca por IP é: <b>go run main.go ip --host ENDEREÇO_WEB_AQUI</b>
- A síntaxe do comando para realizar a busca por Nome do servidor é: <b>go run main.go servers --host ENDEREÇO_WEB_AQUI</b>

## 🔨 Funcionalidades e Saídas do projeto 
- Exemplo da busca de IPs: `go run main.go ip --host amazon.com.br` <br> 
 A saída seriam os IPs públicos do endereço passado:<br>
  <ul><b>IP</b>: <i>54.239.26.87</i></ul>
  <ul><b>IP</b>: <i>72.21.203.171</i></ul>
  <ul><b>IP</b>: <i>52.94.225.243</i></ul><br>
- Exemplo da busca de nome do(s) servidor(es): `go run main.go servers --host amazon.com.br`<br> 
A saída seriam os nomes dos servidores onde esse endereço está hospedado:<br>
    <ul><b>Server</b>: <i>pdns1.ultradns.net.</i></ul>
    <ul><b>Server</b>: <i>pdns4.ultradns.org.</i></ul>
    <ul><b>Server</b>: <i>ns1.p31.dynect.net.</i></ul>
    
## Contribuidores
[<img src="https://avatars.githubusercontent.com/u/60713792?v=4" width=115><br><sub>Fabiano Rauzer</sub>](https://github.com/fabianorauzer-egsys)  
