Test the application on your local Kubernetes cluster

```bash
# Setup the local Kubernetes Cluster
$ minikube start --cpus=4 --memory=4096
# Enable dashboard
minikube addons enable dashboard
# Enable ingress
minikube addons enable ingress


# Initialize dapr
$ dapr init -k
# Check if the dapr installation worked
$ kubectl get pods --namespace dapr-system

# Add the bitnami helm charts repository
$ helm repo add bitnami https://charts.bitnami.com/bitnami
# Start a redis cluster on your local Kubernetes cluster
$ helm install redis bitnami/redis


```