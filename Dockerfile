FROM golang:tip-trixie AS base

# define work directory
WORKDIR /app

# copy the sourcecode
COPY . .

# build exec
RUN cd /app/cmd/server && go mod vendor && CGO_ENABLED=0 GOOS=linux go build -o backend-service

FROM alpine:3.22

WORKDIR /app

COPY --from=base app/cmd/server/backend-service .
COPY --from=base app/template .

# EXPOSE 8080 is the port that the REST API will be exposed on
EXPOSE 8080

CMD [ "./backend-service" ]