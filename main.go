package main

// _ imports the package without requiring explicit use, but .
// imports the package at the root level allowing use of funcs
// in the package without having to refer to the package
import . "./greet"

func main() {
  SayHi("World")
  SayHi("Daniel")

  // This is not an exported function, so these calls would not work
  // greet.logger("Hello")
  // logger("Hello")
}
