package main

func main() {
	r := Router()           // Invoking application
	r.Run("localhost:8080") // specifing address and port no for application run
}
