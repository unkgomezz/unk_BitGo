FROM golang
WORKDIR /app
RUN git clone 
WORKDIR /app/btcgo
RUN rm -rf .git
RUN go mod tidy
RUN go build -o btcgo ./src
CMD ["./btcgo"]