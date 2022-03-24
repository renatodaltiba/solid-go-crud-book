<h1 align="center">
  CRUD SIMPLES UTILIZANDO CONCEITOS DO SOLID
</h1>

<p align="center">
 <a href="#sobre-o-projeto">Sobre o Projeto</a> •
 <a href="#tecnologias">Tecnologias</a> •
 <a href="#configurações-necessárias">Configurações necessárias</a> •
</p>

## Sobre o projeto

Esse simples projeto tem como objetivo o aprendizado em cima do Go, fora implementado conceitos como SRP, DIP, LSP entre outros.
O mesmo executa um CRUD simples de uma tabela books dentro do banco de dados.

- [Go lang](https://go.dev/)
- [Postgres SQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Docker-compose](https://docs.docker.com/compose/)

## Configurações necessárias

## **Requisitos Opção 01 - Docker**

Necessário realizar as instalações.

- [Go lang](https://go.dev/)

### **Clonando o projeto**

```bash
  # Execute o comando git clone para realizar o clone do repositório
  $ git clone https://github.com/renatodaltiba/solid-go-crud-book

  # Entre na pasta do repositório clonado
  $ cd solid-go-crud-book
```
### **Iniciando a database**
```bash
  $ docker-compose up
```
### **Inserir a tabela books no banco**
  Execute manualmente o código

```bash
  CREATE TABLE public.books (
  description character varying(255) NULL,
  author character varying(255) NULL,
  title character varying(255) NULL,
  id integer NOT NULL
);
ALTER TABLE
  public.books
ADD
  CONSTRAINT books_pkey PRIMARY KEY (id)
  ```

### **Iniciando o projeto**

```bash
  # Execute yarn para instalar as dependências
  $ go run ./cmd/server/main.go

```

