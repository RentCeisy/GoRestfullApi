FROM golang:latest
COPY . .
WORKDIR $GOPATH/src/github.com/RentCeisy/GoRestfullApi/cmd/GoRestfullApi
RUN go get -d -v ./../../...
EXPOSE 8080
