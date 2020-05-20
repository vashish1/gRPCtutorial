FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /cli
WORKDIR  /cli

COPY . .

#RUN GO GET COMMAND TO GET ALL THE PACKAGES
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o UserCli


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /cli
WORKDIR /cli

COPY --from=builder /cli/UserCli .

CMD ["./UserCli"]