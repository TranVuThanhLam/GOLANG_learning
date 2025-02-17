package maps

import "fmt"

func main() {
	url := map[string]string{
		"google":  "https://google.com",
		"amazone": "https://aws.com",
	}

	url["facebook"] = "https://facebook.com"
	fmt.Print(url)

}
