==>file main.go 

Imports:
fmt: Used for formatting and printing output.
os: Provides functions for interacting with the operating system, including environment variables.

Main Function:
os.Environ() retrieves a slice of strings containing all environment variables.
The for loop iterates over each environment variable.
Split("=", 1) splits each environment variable into a key and value pair using the "=" delimiter.
fmt.Printf("%s=%s\n", key, value) prints the key-value pair to the console.

***how to run
To run this program, save it as a .go file (e.g., env_vars.go) and compile it using go build
go build env_vars.go
Then, execute the compiled binary:
./env_vars