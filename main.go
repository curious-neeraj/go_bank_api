package main

func main() {

	// start the server
	server := NewAPIServer(":3000")
	server.Run()
}
