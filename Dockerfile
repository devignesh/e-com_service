FROM golang:alpine

WORKDIR /app
COPY . .
RUN go mod download
RUN ls 
RUN go build -o josh-com .
EXPOSE 8080
CMD ["./blogs-api"]