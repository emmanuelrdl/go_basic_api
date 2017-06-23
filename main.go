//https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql



package main

// import "os"

func main() {
    a := App{}
    a.Initialize(
        "emmanuelrudelle",
        "password",
        "go_api")

    a.Run(":8080")
}
