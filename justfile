#!/usr/bin/env just --justfile

# Generates the keypair for the gateway
gen-jwt-key:
    #!/usr/bin/env bash
    set -euxo pipefail
    outdir="certificates"
    mkdir -p "$outdir"
    openssl genrsa -out "$outdir/jwt-key.pem" 2048
    openssl rsa -in "$outdir/jwt-key.pem" -pubout -out "$outdir/jwt-key.pub"

# Prepares the Kubernetes cluster for the ebike3 application
kubectl-prepare:
  kubectl delete secret jwt-private-key --ignore-not-found --namespace=ebike3
  kubectl delete secret jwt-public-key --ignore-not-found --namespace=ebike3

  kubectl create secret generic jwt-private-key --from-file=certificates/jwt-key.pem --namespace=ebike3
  kubectl create secret generic jwt-public-key --from-file=certificates/jwt-key.pub --namespace=ebike3
