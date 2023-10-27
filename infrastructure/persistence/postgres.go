package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/wellingtonpires/fast-food-tech-challenge/domain/entity/cliente"
	"github.com/wellingtonpires/fast-food-tech-challenge/domain/entity/pedido"
	"github.com/wellingtonpires/fast-food-tech-challenge/domain/entity/produto"
)

const conexaoAberta = "Conexáo aberta com o banco!"

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Conectado ao banco!")
	}

	err = db.Ping()

	return db, err
}

func BuscaPedidos() (pedidos []pedido.Pedido) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}
	defer con.Close()
	rows, err := con.Query(`SELECT * FROM pedidos`)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p pedido.Pedido
		err = rows.Scan(&p.Lanche, &p.Acompanhamento, &p.Bebida, &p.IdCliente, &p.NumPedido)
		pedidos = append(pedidos, p)
	}
	return pedidos
}

func CadastraCliente(cli cliente.Cliente) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}
	defer con.Close()
	_, err = con.Exec(`INSERT INTO cliente VALUES ($1, $2, $3)`, cli.Nome, cli.Cpf, cli.Senha)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DadosCliente(cli cliente.Cliente) (dadosCliente []cliente.Cliente) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}
	defer con.Close()
	rows, err := con.Query(`SELECT cpf, senha FROM cliente`)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var c cliente.Cliente
		err = rows.Scan(&c.Cpf, &c.Senha)
		dadosCliente = append(dadosCliente, c)
	}
	return dadosCliente
}

func CadastraProduto(p produto.Produto) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}
	defer con.Close()
	_, err = con.Exec(`INSERT INTO produtos VALUES ($1, $2, $3, $4, $5)`, p.Id, p.Nome, p.Valor, p.Categoria, p.Descricao)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func AtualizaProduto(p produto.Produto) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}
	defer con.Close()
	_, err = con.Exec(`UPDATE produtos SET nome = $1, valor = $2, categoria = $3, descricao = $4 WHERE id = $5;`, p.Nome, p.Valor, p.Categoria, p.Descricao, p.Id)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ExcluiProduto(p produto.Produto) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}
	defer con.Close()
	_, err = con.Exec(`DELETE FROM produtos WHERE id = $1`, p.Id)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ConsultaProdutoPorCategoria(cat string) (prod []produto.Produto) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}
	defer con.Close()
	rows, err := con.Query(`SELECT * FROM produtos WHERE categoria = $1`, cat)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p produto.Produto
		err = rows.Scan(&p.Id, &p.Nome, &p.Valor, &p.Descricao, &p.Categoria)
		prod = append(prod, p)
	}
	return prod
}
