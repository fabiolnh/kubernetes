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
  * It is similar to "Replication Controller", however, ReplicaSet is newer (apps/v1) and has the "selector", to select the labels
- Deployment: It is how the application will be deployed. RepicaSet does not change versions of pods during applyment (you have to stop and start again), that's why we use Deployment. With deployment, we can doo rolling update, pause, resume, etc.
- StatefulSet: Usualy the same as ReplicaSet, however it is used with persistent storage
- DaemonSet: One pod will be running in each node (ex: 3 nodes will have 3 pods, 1 for each node) 
- Namespace: Contexts. A way to organize the things in kubernetes (ex: each namespace for each squad)
- Service: First Gateway for the Application. Acts like a Load Balancer. It generates a network for the pods to communicate between them, too.
- Ingress: The unique entry to the application. In Ingress, we can configure some routes to select which service is to redirect the entry, based on the path, like a reserve proxy.
- The "Service" object is used to expose internal services of the cluster, while the "Ingress" object is used to expose services to the world.
- Types of "Service":
  * ClusterIP: Internal IP (only accessible within the cluster).
  * LoadBalancer: ClusterIP + External Load Balancer (commonly provided by the cloud). This allows you to access your application from the external world.
- ServiceAccount: A way to secure what the pod may access. If someone breaks into the pod, if the pod is using the default ServiceAccount of the Kubernetes, it will get access to everything in kubernetes. The best way is to create a different ServiceAccount to limit what it can do. Another way to use the ServiceAccount is to link a role to access some service in the cloud 
- HPA: Horizontal Pod Autoscaling
- Rolling Update: A deploy way. By default, this is the default. You do not need to specify. Use it with livenessProbe and readnessProbe to get Zero-Downtime in deploys and in runtime.
- ConfigMap: Useful to: inject env variables, create physical files into containers, work with sensible data, etc.
- Probes: Ex: If the container is down (if /health returns a code differently from 200) it restarts the container. (POC: https://github.com/fabiolnh/kubernetes/blob/main/k8s/03-deployment-noconfigmap-liveness.yaml)
- OBS: The object "Secret" is not so safe.
- Other important things are annotated inside the files
