FROM golang:1.21

WORKDIR /project/go-docker/

# COPY go.mod, go.sum and download the dependencies
COPY go.* ./
RUN go mod download

# COPY All things inside the project and build
COPY . .
RUN go build -o /project/go-docker/build/myapp .

EXPOSE 8084

ENV GIN_MODE=debug

ENTRYPOINT [ "/project/go-docker/build/myapp" ]