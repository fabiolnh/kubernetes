# Kubernetes Study

## Steps:
1) Install Docker
2) Install Kubectl
3) Install Kind (A kubernetes docker solution. Usable for studying)
## Create Cluster:
```
kind create cluster --config=k8s/kind.yaml --name=fabiolnh-cluster
```
## Apply Configuration:
```
kubectl apply -f k8s/*.yaml
```

## Important:
1) RepicaSet does not change version of pods during apply. (we have to delete pods)