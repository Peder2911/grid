// A simple framework for running stepwise simulations on gridded data.  A
// simulator is simply a function that takes and yields a cube. This way, each
// element in the simulation can be written as a standalone function that is
// nice, simple and testable.
package sim

import (
   "github.com/peder2911/grid/mat"
   "math/rand"
)

// Compose simulator functions into a single function.
func Simulator(funcs... func(mat.Cube) mat.Cube) func(mat.Cube) mat.Cube {
   return func (cube mat.Cube) mat.Cube {
      for _,fn := range funcs {
         cube = fn(cube)
      }
      return cube
   }
}

// Simple function for seeding a cube with random integers
func Intseed(cube mat.Cube) mat.Cube {
   xsize, ysize, _ := cube.Size()
   for x := 0; x < xsize; x ++ {
      for y := 0; y < ysize; y ++ {
         cube[x][y][0] = rand.Intn(2)
      }
   }
   return cube
}

// Conways game of life
//
// Each cell with two neighbours will carry on as it is.
// A dead cell with three neighbours will come alive.
// A live cell with three neighbours will live.
// Otherwise, the cell dies.
//
// Sometimes creates fun patterns. A useful exercise to prove the concept of
// simulator functions.
//
func Gol(cube mat.Cube) mat.Cube{
   xsize, ysize, _ := cube.Size()
   for x := 0; x < xsize; x ++ {
      for y := 0; y < ysize; y ++ {
         current := cube[x][y][0]
         liveneighbours := 0
         xneighbours := []int{x-1,x,x+1}
         yneighbours := []int{y-1,y,y+1}
         if x == 0 {
            xneighbours = xneighbours[1:]
         }
         if x == xsize-1 {
            xneighbours = xneighbours[:2]
         }
         if y == 0 {
            yneighbours = yneighbours[1:]
         }
         if y == ysize-1 {
            yneighbours = yneighbours[:2]
         }
         for _,xnv := range xneighbours {
            for _,ynv := range yneighbours {
               if ! (xnv == x && ynv == y) {
                  liveneighbours += cube[xnv][ynv][0]
               }
            }
         }

         var result int
         if liveneighbours == 2 {
            result = current
         } else if liveneighbours == 3 {
            result = 1
         } else {
            result = 0
         }
         cube[x][y][0] = result
      }
   }
   return cube
}
