package main

import "fmt"
import "github.com/google/uuid"

func main()  {
	Hello()
}

func Hello()  {
	fmt.Println( uuid.New())
}

func Hello2()  {
	fmt.Println("test")
}