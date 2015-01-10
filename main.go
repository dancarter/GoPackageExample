package main

import g "./greet"

func main() {
  g.SayHi("World")
  g.SayHi("Daniel")

  // This is not an exported function, so these calls would not work
  // g.logger("Hello")
}
