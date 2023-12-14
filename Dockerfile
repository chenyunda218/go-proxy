FROM golang:1.21 as build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /main
EXPOSE 10000
CMD [ "/main" ]