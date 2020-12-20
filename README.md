# distributed-fs

WIP 

## build and run project

#### generate protobuf files
`protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/dfs.proto`


#### docker 
`docker build -t namenode .`

`docker run --name namenode --mount type=bind,source="$(pwd)",target=/usr/src/app -p target=8001,published=8001 -dit namenode`


#### start data node server
`go run dataNode/start/start.go`

#### start name node server
`go run namenode/start/start.go`


## resources
HDFS architecture 
http://itm-vm.shidler.hawaii.edu/HDFS/

GFS paper
https://research.google.com/archive/gfs-sosp2003.pdf