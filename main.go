package main

import "fmt"
import "github.com/google/uuid"

func main()  {
	Hello()
	Hello2()
}

func Hello()  {
	fmt.Println( uuid.New())
}

func Hello2()  {
	fmt.Println("test")
	fmt.Println("test")
}


func Hello3()  {
	fmt.Println("test3")
}