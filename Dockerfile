FROM golang:1.16-alpine AS build

WORKDIR /src/

COPY ./main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/hello-mundo

FROM alpine
COPY --from=build /bin/hello-mundo /bin/hello-mundo
EXPOSE 8080
ENTRYPOINT ["/bin/hello-mundo"]
