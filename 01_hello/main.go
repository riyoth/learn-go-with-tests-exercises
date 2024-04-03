package main

import "fmt"

const englishHelloPrefix = "Hello "
const defaultName = "world"

func Hello(name string) string{
	if name==""{
		
		return englishHelloPrefix+defaultName
	}
	return englishHelloPrefix+name
}

func main(){
	fmt.Println(Hello("chris"))
}
