package main

import "tanveershafeeprottoy.com/fileuploaddemo/app"

func main() {
	application := &app.App{}
	application.Init()
	application.Run()
}
