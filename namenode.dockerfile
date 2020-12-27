FROM golang

WORKDIR /usr/src/app

COPY . .

RUN go mod download

CMD [ "go", "run", "./namenode/start/start.go"]

EXPOSE 8001/tcp