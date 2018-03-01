FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/tcarlton2000/cookbook-api/api

# Install our dependencies
RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN go get github.com/xeipuuv/gojsonschema
RUN go get github.com/martinlindhe/unit

# Install api binary globally within container
RUN go install github.com/tcarlton2000/cookbook-api/api

# Set binary as entrypoint
ENTRYPOINT /go/bin/api

# Expose default port (8080)
EXPOSE 8080
