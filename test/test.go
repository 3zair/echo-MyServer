package main

import (
	"log"
)

func protect(g func()) {
	defer func() {
		log.Println("done")
		if err := recover(); err != nil {
			log.Printf("run time panic:%v", err)
		}
	}()
	log.Println("start")
	g()
}

//func main() {
//	protect(func() {
//		fmt.Println("hello world")
//	})
//}