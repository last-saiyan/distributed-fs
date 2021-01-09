# distributed-fs

WIP 

## build and run project

#### generate protobuf files
`protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/dfs.proto`  

#### run tests
`cd namenode && go test -v`  
`cd dataNode && go test -v`


#### docker images
`docker build -f ./deployment/namenode.dockerfile . -t namenode`  
`docker build -f ./deployment/datanode.dockerfile . -t datanode`


to run docker image with fs mount  
`docker run --name namenode --mount type=bind,source="$(pwd)",target=/usr/src/app -p target=8001,published=8001 -dit namenode`


#### start data node server
`go run dataNode/start/start.go`

#### start name node server
`go run namenode/start/start.go`


#### create a cluster on minikube
`./start.sh`


## resources
HDFS architecture  
http://itm-vm.shidler.hawaii.edu/HDFS/  
GFS paper  
https://research.google.com/archive/gfs-sosp2003.pdf