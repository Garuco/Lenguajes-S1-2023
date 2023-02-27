package main

import "fmt"

// Function that receives the list of elements, quantity of rotations, and direction of rotations. 0 = left, 1 = right.
/*func rotate(sliceptr *[]int, rotations int, dir int) []int {
	var temp []int

	if dir == 1 {
		temp = append(temp, (*sliceptr)[:len(*sliceptr)-rotations]...)
		temp = append((*sliceptr)[len(*sliceptr)-rotations:], temp...)
	} else {
		temp = append(temp, (*sliceptr)[rotations:]...)
		temp = append(temp, (*sliceptr)[:rotations]...)
	}
	return temp
}*/

// Function that creates a copy of a given list, and rotate it a given times to a given direction, returning the modified copy.
// Receives a pointer to the list, quantity of rotations, and direction of rotations. 0 = left, 1 = right.
func rotate2(list *[5]int, rotations int, dir int) [5]int {
	var temp int
	listTemp := *list

	if dir == 1 { // Rotation to the right.
		for i := 0; i < rotations; i++ { // Loop of given rotation times
			//Since we are rotating to the right, we need to save the last element of the list.
			temp = listTemp[len(listTemp)-1]
			// Loop to replace elements from the last to the second.
			for j := len(listTemp) - 1; j > 0; j-- {
				listTemp[j] = listTemp[j-1] //Move the element to the right.
			}
			listTemp[0] = temp //We replace the first element with the last one that we saved before.
		}
	} else { // Rotation to the left.
		//Same process as above, but from left to right, saving the first element, then moving elements to the left
		// and finally, replacing the last element with the first one that we saved before.
		for i := 0; i < rotations; i++ {
			temp = listTemp[0]
			for j := 0; j < len(listTemp)-1; j++ {
				listTemp[j] = listTemp[j+1]
			}
			listTemp[len(listTemp)-1] = temp
		}
	}
	return listTemp
}

func main() {
	var originalList = [5]int{1, 2, 3, 4, 5}
	var rightRotationList [5]int
	var leftRotationList [5]int

	var rightRotations = 2
	var leftRotations = 2

	var ptr = &originalList

	rightRotationList = rotate2(ptr, rightRotations, 1)
	leftRotationList = rotate2(ptr, leftRotations, 0)

	fmt.Println("Secuencia original:", originalList)
	fmt.Println("Secuencia con", rightRotations, "rotaciones a la derecha:", rightRotationList)
	fmt.Println("Secuencia con", leftRotations, "rotaciones a la izquierda:", leftRotationList)
}
