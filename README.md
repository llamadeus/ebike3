# Ebike3
This is the repository for the third and final assignment of the Software Architecture and Platforms course at the University of Bologna (Campus Cesena).

The goal of this assignment is to build an "EBike application" based on an event-driven microservices architecture, applying the event sourcing pattern where considered useful. Finally, the application should be deployed on a distributed infrastructure using Kubernetes. The exact assignment description can be found [here](https://github.com/sap-2024-2025/Assignment-03).

## Running the Application
The application is developed from the ground up using [minikube](https://minikube.sigs.k8s.io/), which lets you run your Kubernetes cluster on your local machine, and [DevSpace](https://devspace.sh), which is a tool that allows you to develop and deploy applications on Kubernetes clusters in a few seconds. Therefore, the easiest way to run the application is to use minikube and DevSpace.

Make sure your Kubernetes cluster is running before proceeding.

### Generate the JWT Keypair
To generate the JWT keypair, run the following command:
```bash
outdir="certificates"
mkdir -p "$outdir"
openssl genrsa -out "$outdir/jwt-key.pem" 2048
openssl rsa -in "$outdir/jwt-key.pem" -pubout -out "$outdir/jwt-key.pub"
```

Next, create the according Kubernetes secrets:
```bash
kubectl create secret generic jwt-private-key --from-file=certificates/jwt-key.pem --namespace=ebike3
kubectl create secret generic jwt-public-key --from-file=certificates/jwt-key.pub --namespace=ebike3
```

This will create two secrets `jwt-private-key` and `jwt-public-key` in the `ebike3` namespace.

### Start minikube
To start minikube, run the following command:
```bash
minikube start
minikube dashboard --url  # Optional, starts the dashboard and prints the URL
```

### Start DevSpace
Now, you can start DevSpace by running the following command:
```bash
devspace dev
```

This will start the application in a Kubernetes cluster and deploy it to the cluster. The frontend should now be available at `http://localhost:3000` and the backend at `http://localhost:4000`.

**Note**: It can take quite some time for the application to start up as it is running in development mode. Check the Minikube dashboard for the status of the pods.
