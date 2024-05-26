FROM golang:1.22

WORKDIR /app

COPY . ./
RUN go mod download
RUN make

EXPOSE 8080

CMD ["exec", "dist/wikara"]
