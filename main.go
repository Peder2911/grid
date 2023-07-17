// A simple simulator. Currently only has an implementation of conways game
// of life, but is expandable.
package main

import (
   "github.com/peder2911/grid/mat"
   "github.com/peder2911/grid/sim"
   "fmt"
   //"time"
)

func main(){
   cube := sim.Intseed(mat.NewCube(64,128,4))
   simulator := sim.Simulator(sim.Gol)
   for i := 1; i < 100000 ; i++ {
      cube = simulator(cube)
      fmt.Println(cube)
      //time.Sleep(time.Millisecond*50)
   }
   fmt.Println("Done!")
}
