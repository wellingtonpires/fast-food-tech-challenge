#!/bin/bash

#Script para encerrar o compose por completo, criar nova build e realizar o start, assim criando uma implantação limpa

minikube start --ports=8080:8080
sudo chmod 777 /var/run/docker.sock
kubectl apply -f fastfoodapp-deployment.yaml,fastfoodapp-service.yaml,postgres-initdb-config.yaml,postgres-claim0-persistentvolumeclaim.yaml,postgres-deployment.yaml,postgres-service.yaml