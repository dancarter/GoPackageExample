package main

import "./greet"

func main() {
	greet.SayHi("World")
	greet.SayHi("Daniel")

	// This is not an exported function, so these calls would not work
	// greet.logger("Hello")
}
