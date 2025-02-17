package lists

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func (p *Product) New(title string, id int, price float64) {
	p.title = title
	p.id = id
	p.price = price
}

func (p Product) Display() {
	fmt.Print("Title: ", p.title, ", Id: ", p.id, ", Price: ", p.price, "\n")
}
func main() {
	// 1
	hobbies := [3]string{"sing", "read", "sleep"}
	fmt.Println(hobbies)
	// 2
	fmt.Println(hobbies[0])
	newList := hobbies[1:3]
	fmt.Println(newList)
	// 3
	slice1 := hobbies[0:2]
	fmt.Println(slice1)
	slice2 := hobbies[:2]
	fmt.Println(slice2)
	// 4
	re_slice := slice1[1:3]
	fmt.Println(re_slice)
	// 5
	dynamic_array := []string{"master_go", "code_4_life"}
	fmt.Println(dynamic_array)
	// 6
	dynamic_array[1] = "code_4_abit"
	dynamic_array = append(dynamic_array, "live_in_happy")
	fmt.Println(dynamic_array)
	// 7

	var product, product1, product2 Product
	product.New("Candy", 0, 3.5)
	product1.New("Cake", 1, 2.3)

	products := []Product{
		{
			"KitKat",
			5,
			15.2,
		},
		{
			"Alibaba",
			7,
			12.3,
		},
	}

	product2.New("keyboard", 3, 5.4)
	// products = append(products, product)
	// products = append(products, product1)
	products = append(products, product2)

	for i := 0; i < len(products); i++ {
		products[i].Display()
	}
}
