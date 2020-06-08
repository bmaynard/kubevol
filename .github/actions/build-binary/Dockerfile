# Specify the version of Go to use
FROM golang:1.13

# Copy all the files from the host into the container
WORKDIR /src
COPY . .

# Enable Go modules
ENV GO111MODULE=on

# Compile the action
RUN go build -o /bin/action

# Specify the container's entrypoint as the action
ENTRYPOINT ["/bin/action"]