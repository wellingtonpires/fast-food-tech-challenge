package persistence

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/wellingtonpires/fast-food-tech-challenge/domain/entity/pedido"
	"github.com/wellingtonpires/fast-food-tech-challenge/domain/entity/cliente"
)

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

func BuscaPedidos() (pedidos []pedido.Pedido){
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Conexáo aberta!")
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

func CadastraCliente(cli cliente.Cliente){
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Conexáo aberta!")
	}
	defer con.Close()
	_, err = con.Exec(`INSERT INTO cliente VALUES ($1, $2, $3)`, cli.Nome, cli.Cpf, cli.Senha)
	if err != nil {
		fmt.Println(err.Error())
	}
}
