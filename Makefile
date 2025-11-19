# Define the binary output name
BINARY_NAME=redbook.exe

# Build the binary
build:
	go build -o $(BINARY_NAME) redbook.go

# Run the program
run: build
	./$(BINARY_NAME)

# Clean the generated binary
clean:
	go clean
	rm -f $(BINARY_NAME)

