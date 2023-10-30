FROM  golang:1.21
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /fastfood
EXPOSE 8080
CMD ["/fastfood"]