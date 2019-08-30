package main

var a1 string
var done bool

func setup() {
	a1 = "jello,world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	print(a1)
}
