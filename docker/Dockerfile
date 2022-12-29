## Get golang image form docker hub

FROM golang:1.19-alpine

ARG PORT

# Define the working directory
WORKDIR /app

# Copy all file to container
COPY . .

# Install All dependence
RUN go mod tidy

# Build the go project to binary
RUN go build -o /app/final-project main.go

# Expose port 8080, then localhost can access the container
EXPOSE ${PORT}

# Run binary build of go project
ENTRYPOINT [ "/app/final-project" ]
