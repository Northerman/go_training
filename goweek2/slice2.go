// There're 2 hw 1)Slice 2)wordcount

package main

import "golang.org/x/tour/pic"

// Go run this in playground
func Pic(dx, dy int) [][]uint8 {
	pic := [][]uint8{}
	for y := 0; y < dy; y++ {
		row := []uint8{}
		for x := 0; x < dx; x++ {
			row = append(row, uint8((x+y)/2)) //same as python row.append(....)
		}
		pic = append(pic, row)
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
