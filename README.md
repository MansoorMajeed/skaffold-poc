# Skaffold POC

## Install Skaffold


https://skaffold.dev/docs/install/

```
# For macOS 
brew install skaffold
```

## Enable Kubernetes 

You can use minikube or docker desktop for local development


```
kubectl config get-contexts | grep docker-desktop
*         docker-desktop       docker-desktop           docker-desktop
```


## Start

```
skaffold dev --kube-context docker-desktop --port-forward
```