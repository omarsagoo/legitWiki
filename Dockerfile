# Golang version - Docker Hub (hub.docker.com)
FROM golang:alpine3.12

# Enviornment variables
ENV APP_NAME totally-legit-wiki
ENV PORT 8080

# Open system port
EXPOSE ${PORT}

# Working directory
WORKDIR /go/src/${APP_NAME}

COPY . /go/src/${APP_NAME}

# Install dependecies from mod file
RUN go mod download

# Compile the binary to run the progra
ENTRYPOINT CompileDaemon --build="go build -o totally-legit-wiki" --command=./totally-legit-wiki