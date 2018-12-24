#!/bin/bash


echo "Creating the volume..."

kubectl apply -f ./kubernetes/persistent-volume.yml
kubectl apply -f ./kubernetes/persistent-volume-claim.yml


echo "Creating the database credentials..."

kubectl apply -f ./kubernetes/secret.yml


echo "Creating the postgres deployment and service..."

kubectl create -f ./kubernetes/postgres-deployment.yml
kubectl create -f ./kubernetes/postgres-service.yml



echo "Creating the Go deployment and service..."

kubectl create -f ./kubernetes/golang-deployment.yml
kubectl create -f ./kubernetes/golang-service.yml


echo "Adding the ingress..."

minikube addons enable ingress
kubectl apply -f ./kubernetes/minikube-ingress.yml


echo "Creating the vue deployment and service..."

kubectl create -f ./kubernetes/vue-deployment.yml
kubectl create -f ./kubernetes/vue-service.yml

echo "Creating Go-Lang HPA"
kubectl create -f ./kubernetes/golang-hpa.yml