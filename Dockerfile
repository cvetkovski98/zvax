FROM olivercvetkovski/zvax-protoc:latest AS protoc
FROM golang:1.19 AS build

WORKDIR /app

# copy proto files
COPY --from=protoc /app/common ./common

WORKDIR /app/src

# copy go mod and sum files
COPY go.mod .
COPY go.sum .

# download dependencies
RUN go mod download

# copy source code
COPY . .

# build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -o /service .

# final stage
FROM scratch AS final

WORKDIR /app

# copy the executable
COPY --from=build /service .

# run the executable
ENTRYPOINT [ "/app/service" ]
