package main

import "fmt"
import "github.com/google/uuid"

func main()  {
	Hello()
	Hello2()
	Hello3()
}

func Hello()  {
	fmt.Println( uuid.New())
}

func Hello2()  {
	fmt.Println("test1-729")
	fmt.Println("test")
}


func Hello3()  {
	fmt.Println("728---")
}
