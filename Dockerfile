FROM golang:1.12-alpine
#Copy the build's output binary from the previous build container
COPY hello hello
ENTRYPOINT ["hello"]