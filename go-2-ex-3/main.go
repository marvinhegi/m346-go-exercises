package main

import "fmt"

func main() {
	modules := make(map[string]int)
	modules["Module 104"] = 104
	modules["Module 117"] = 117
	modules["Module 346"] = 346

	fmt.Println("Modul 104:", 104)
	fmt.Println("Modul 117:", 117)
	fmt.Println("Modul 346:", 346)

	delete(modules, "Module 117")
	modules["Module 199"] = 199
	modules["Module 346"] = 222

	fmt.Println(modules)
}
