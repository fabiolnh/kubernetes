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
- Pod: A Kubernetes Object that can have one or more containers
- ReplicaSet: it defines the number of a pod and recreates it when it is down
- Deployment: RepicaSet does not change versions of pods during applyment (you have to stop and start again), that's why we use Deployment
- Service: First Gateway for the Application. Acts like a Load Balancer. It generates a network for the pods to communicate between them, too.
- Namespace: Conexts. A way to organize the things in kubernetes (ex: each namespace for each squad)
- Ingress: The unique entry to the application. In Ingress, we can configure some routes to select which service is to redirect the entry, based on the path, like a reserve proxy. 
- ServiceAccount: A way to secure what the pod may access. If someone breaks into the pod, if the pod is using the default ServiceAccount of the Kubernetes, it will get access to everything in kubernetes. The best way is to create a different ServiceAccount to limit what it can do. Another way to use the ServiceAccount is to link a role to access some service in the cloud 
- HPA: Horizontal Pod Autoscaling
- Rolling Update: A way to update pods with zero-downtime
- ConfigMap: Useful to: inject env variables, create physical files into containers, work with sensible data, etc.
- Probes: Ex: If the container is down (if /health returns a code differently from 200) it restarts the container. (POC: https://github.com/fabiolnh/kubernetes/blob/main/k8s/03-deployment-noconfigmap-liveness.yaml)
- OBS: The object "Secret" is not so safe.
- Other important things are annotated inside the files
