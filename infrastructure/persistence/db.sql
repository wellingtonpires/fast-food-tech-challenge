ALTER USER postgres PASSWORD '123';
CREATE TABLE produtos(nome VARCHAR(50), valor REAL, categoria VARCHAR(50), descricao VARCHAR(200), id VARCHAR(50));
CREATE TABLE clientes(nome VARCHAR(50), cpf VARCHAR(11), senha VARCHAR(50), id SERIAL);
CREATE TABLE pedidos(lanche VARCHAR(50), acompanhamento VARCHAR(50), bebida VARCHAR(50), statuspagamento VARCHAR(20), statuspreparacao VARCHAR(20), codpedido VARCHAR(50), idcliente VARCHAR(50));