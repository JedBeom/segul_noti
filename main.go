package main

func main() {
	c := jobInit()
	c.Start()
	<-make(chan int)
}
