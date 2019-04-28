package main

import (
   "fmt"
   "./kinsol"
   "./nvector"
)

// https://github.com/modelon/Assimulo/blob/master/tests/solvers/test_kinsol.py


func main() {
   pointer := kinsol.KINCreate()
   fmt.Println( pointer )
}
