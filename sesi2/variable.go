package main

import "fmt"

type Student struct {
	Name  string
	Class string
}

func (s *Student) SetName(newName string) {
	fmt.Println("change name from", s.Name, "to", newName)
	s.Name = newName
	fmt.Println("New name is", s.Name)
}

func main() {

	var firstName, lastName, _ = "Reyhan", "Jovie", 9
	firstName = setName("Hacktiv8")
	firstName = setName("Hacktiv8")
	fmt.Printf("Nama Lengkap : %+v-%#v\n", firstName, lastName)
	fmt.Println("Nama Lengkap :", firstName+"-"+lastName)

	name := ""
	name = "Hello"

	name2 := &name
	fmt.Println(*name2)
	fmt.Println(name)
	SetPointerName(&name, "bank neo")
	fmt.Println(name)
	fmt.Println(*name2)

	students := make(map[string]string, 0)
	students["name"] = "Hacktiv8"
	fmt.Println(students)
	fmt.Println(len(students))

	student := Student{}
	student.SetName("Golang")
	fmt.Println(student)
}

func setName(newName string) string {
	return newName
}

func SetPointerName(oldName *string, newName string) {
	*oldName = newName
}
