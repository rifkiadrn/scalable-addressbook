#!/bin/bash


echo "Creating the volume..."

kubectl delete -f ./kubernetes/persistent-volume.yml
kubectl delete -f ./kubernetes/persistent-volume-claim.yml


echo "Creating the database credentials..."

kubectl delete -f ./kubernetes/secret.yml


echo "Creating the postgres deployment and service..."

kubectl delete -f ./kubernetes/postgres-deployment.yml
kubectl delete -f ./kubernetes/postgres-service.yml



echo "Creating the Go deployment and service..."

kubectl delete -f ./kubernetes/golang-deployment.yml
kubectl delete -f ./kubernetes/golang-service.yml


echo "Adding the ingress..."

minikube addons enable ingress
kubectl apply -f ./kubernetes/minikube-ingress.yml


echo "Creating the vue deployment and service..."

kubectl delete -f ./kubernetes/vue-deployment.yml
kubectl delete -f ./kubernetes/vue-service.yml

echo "Creating Go-Lang HPA"
kubectl delete -f ./kubernetes/golang-hpa.yml