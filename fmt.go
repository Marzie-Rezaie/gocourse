package main

import "fmt"

func main() {

	var name string
	var age int
	fmt.Println("please enter name and age:")
	fmt.Scan(&name, &age)
	// fmt.Scanln(&name, &age)
	// fmt.Scanf("%s %d", &name, &age)

	fmt.Printf("name is:%s\n", name)
	fmt.Printf("age is:%d\n", age)

	err := checkAge(13)
	if err != nil {
		fmt.Println(err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is not legal", age)
	}
	return nil
}
