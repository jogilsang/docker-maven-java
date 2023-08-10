FROM golang:1.20-alpine
#Copy the build's output binary from the previous build container
COPY hello.go .
RUN go build hello.go
#COPY hello hello
# ENTRYPOINT ["hello"]
CMD ["./hello"]