package practica

import (
	"fmt"
)

func getMaximumNumber(nums []int) int {
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if (nums[i] > max) {
			max = nums[i]
		}
	}
	return max
}

func Ej1() {
	// Realizar un algoritmo para determinar el mayor de 3 números. Y luego para determinar el mayor de 10 numeros.
	fmt.Print("EJERCICIO 1\n\n")

	nums1 := []int{1,2,3}
	max1 := getMaximumNumber(nums1)

	fmt.Println("El número más grande entre", nums1, " es:", max1)

	nums2 := []int{1,2,3,4,5,6,7,8,9,10}
	max2 := getMaximumNumber(nums2)

	fmt.Println("El número más grande entre", nums2, " es:", max2)
}
