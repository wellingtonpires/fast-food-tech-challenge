#!/bin/bash

#Script para encerrar o minikube por completo, criar nova build e realizar o start, assim criando uma implantação limpa

minikube delete
minikube start --ports=8080:8080
sudo chmod 777 /var/run/docker.sock
minikube addons enable metrics-server
kubectl apply -f fastfoodapp-deployment.yaml,fastfoodapp-service.yaml,postgres-initdb-config.yaml,postgres-claim0-persistentvolumeclaim.yaml,postgres-deployment.yaml,postgres-service.yaml
kubectl autoscale deployment fastfoodapp --cpu-percent=80 --min=1 --max=10
kubectl autoscale deployment postgres --cpu-percent=80 --min=1 --max=10