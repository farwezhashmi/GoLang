package main // In Go all the code is in packages 

import "fmt" // This is one of package , 
// in which the basic I/O functions available similarlly like <stdio.h> in C 

func main() {  // This is the main block of our code which is excuted main

	fmt.Println("Hello,World!")   // We use dot operator call the func where it is from which package

}
// To run this programme ,we need create a go.mod file which holds the project name and Go's version
// To run the code use "go run <file_name.go>"