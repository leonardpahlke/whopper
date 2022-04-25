# Local application testing


1. Setup local Kubernetes cluster
2. Initialize Dapr
3. Install redis
4. Add server pod deployments
5. Use the whopper client to test the deployed application

## 1. Setup local Kubernetes cluster
Start a local Kubernetes cluster (see [minikube getting started](https://minikube.sigs.k8s.io/docs/start/))
```bash
# Setup the local Kubernetes Cluster
$ minikube start --cpus=4 --memory=4096
# Enable dashboard
$ minikube addons enable dashboard
# Enable ingress
$ minikube addons enable ingress
```

## 2. Initialize Dapr
Initialize dapr (see [dapr docs](https://docs.dapr.io/operations/hosting/kubernetes/kubernetes-deploy/#install-with-dapr-cli))
```bash
# Initialize dapr
$ dapr init -k
# Check if the dapr installation worked
$ kubectl get pods --namespace dapr-system
```

## 3. Install redis
Add a statestore pod to the Kubernetes cluster (see [bitnami redis helm chart](https://github.com/bitnami/charts/tree/master/bitnami/redis))

Since this application has pods that store data persistently a redis k/v-store is getting installed via a helm chart on the cluster. This pod is getting registered as statestore in dapr.
```bash
# Add the bitnami helm charts repository
$ helm repo add bitnami https://charts.bitnami.com/bitnami

# Start a redis cluster on your local Kubernetes cluster
$ helm install redis bitnami/redis
```

## 4. Add server pod deployments
Next up the server pods are getting deployed on the Kubernetes cluster

```bash
```

## 5. Use the whopper client to test the deployed application

```bash
```