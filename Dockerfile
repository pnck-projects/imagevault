FROM golang:1.23-alpine AS build
WORKDIR /src
COPY . .
RUN apk update &&  \
    apk upgrade &&  \
    apk add --no-cache ca-certificates && \
    update-ca-certificates
RUN go mod download &&  \
    go mod verify &&  \
    go mod vendor && \
    go build -v -o /out/imagevault
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init

FROM scratch AS imagevault
WORKDIR /
USER 10001:10002
EXPOSE 8080
COPY --from=build /out/imagevault /usr/local/bin/imagevault
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/docs/swagger.json /docs/swagger.json
COPY version /version
ENTRYPOINT ["imagevault"]

