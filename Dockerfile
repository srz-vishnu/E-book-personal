FROM golang:1.23.5-alpine
WORKDIR /e-book
COPY . ./
RUN go mod download
RUN go build -v -o /output .
EXPOSE 8080
CMD ["/output"]


