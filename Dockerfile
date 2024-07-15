FROM mirror.gcr.io/library/golang:1.22-alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o /go/bin/ssh

############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/ssh /go/bin/ssh
# Run the ssh binary.
ENTRYPOINT ["/go/bin/ssh"]

