# STEP-1
FROM golang:1.20-alpine

# STEP-2
COPY hello.go .
COPY static/ static/

# STEP-3
RUN go build hello.go

# STEP-4
CMD ["./hello"]
