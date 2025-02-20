package main

func main() {
	number := &[]int{5, 6, 7}
	process(number)
	// fmt.Println((*number)[0])
}

func process(number *[]int) {
	for i, _ := range *number {
		(*number)[i] *= 2
	}

}
