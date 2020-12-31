# ensure minikube is up and running before executing this

# delete existing kubernetes resources
kubectl delete all --all

# use minikube docker env
eval $(minikube docker-env)

# build docker images on minikube node to include latest changes
docker build -f ./deployment/namenode.dockerfile . -t namenode
docker build -f ./deployment/datanode.dockerfile . -t datanode

# create statefulset deployment so ipAddr is same
kubectl apply -f ./deployment/datanode.yml
kubectl apply -f ./deployment/namenode.yml

# create headless service for communication. no clusterip
kubectl apply -f ./deployment/datanodeService.yml
kubectl apply -f ./deployment/namenodeService.yml
