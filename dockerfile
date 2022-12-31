# create main.exe
FROM golang:latest as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
# EXPOSE 80
# ENTRYPOINT [ "./main" ]
# build distroless image
FROM gcr.io/distroless/static-debian11
COPY --from=build /app/main ./
EXPOSE 80
CMD ["/main"]