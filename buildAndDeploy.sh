#!/bin/bash

#Script para encerrar o compose por completo, criar nova build e realizar o start, assim criando uma implantação limpa

#sudo docker rm -vf $(docker ps -aq)
#docker rmi -f $(docker images -aq)
sudo docker compose down -v
sudo docker compose build
sudo docker compose up