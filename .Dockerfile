FROM golang:1.21.1-alpine AS build

WORKDIR /common
COPY . .

RUN apk add --no-cache git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o common

FROM alpine:latest

WORKDIR /common
COPY --from=build /common/common /common/common

EXPOSE 8091

CMD ["./common"]
