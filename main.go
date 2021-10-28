package main

import "shop/web"

func main() {
	if web.Connect() {
		web.Run()
	}
}
