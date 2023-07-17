
package mat
import (
   "strings"
)

type Column []int
type Matrix []Column
type Cube []Matrix

func NewMatrix(x int, y int) Matrix {
   var mat Matrix = make(Matrix, x)
   for i := 0; i < x; i ++ {
      mat[i] = make(Column, y)
   }
   return mat
}

func NewCube(x int, y int, z int) Cube {
   var cube Cube = make(Cube, x)
   for i := 0; i < x; i ++ {
      cube[i] = NewMatrix(y, z) 
   }
   return cube
}


func (c Cube) Size() (int, int, int){
   return len(c), len(c[0]), len(c[0][0])
}

func (c Cube) String() string{
   repr := ""
   xsize, ysize, _ := c.Size()
   repr += "┌" + strings.Repeat("─",ysize) + "┐"
   for x := 0; x < xsize ; x++ {
      repr += "\n|"
      for y := 0; y < ysize ; y++ {
         value := c[x][y][0]
         if value == 0 {
            repr += " "
         } else {
            repr += "●"
         }
      }
      repr += "│"
   }
   repr += "\n└" + strings.Repeat("─",ysize) + "┘"
   return repr
}
