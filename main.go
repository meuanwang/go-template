package main

import ( 
	"project-demo/Router"
)

func main() {
	r := router.RouterInit()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}