FROM golang:alpine
RUN apk update && apk add git
RUN mkdir /yuki-api
WORKDIR /yuki-api
ADD . /yuki-api
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /yuki-api/app .
CMD ["./app"]