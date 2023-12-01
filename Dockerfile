# Use an official Go runtime as a parent image
FROM golang:1.19.2

# Copy the Go source code into the container
COPY cricket_scoreboard.go .

# Build the Go application
RUN go build cricket_scoreboard.go

# Make port 80 available to the world outside this container
EXPOSE 80

# Run the Go application when the container launches
CMD ["./cricket_scoreboard"]
