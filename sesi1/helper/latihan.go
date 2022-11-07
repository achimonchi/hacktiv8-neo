package helper

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	1. func CheckIsPrime(int)(bool)
	2. func GenerateID(prefix string, barang []Barang) (int)
		GenerateID("BAR",barang) => BAR-3
		pastikan, ID nya uniq

	unit test
	1. TestIsPrime_Success
	2. TestIsPrime_Fail
	3. TestGenerateID_Success
		- harus uniq
		a. Test Case 1
			input :
			barang => [
				{id:BAR-1,name:"baju"},
				{id:BAR-2,name:"baju"},
			].

			output :
				BAR-3
		b. Test Case 2
			input :
			barang => [
				{id:BAR-2,name:"baju"},
				{id:BAR-3,name:"baju"},
			].

			output :
				BAR-1
*/

func CheckIsPrime(num int) (isPrime bool) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

type Barang struct {
	Id   string
	Name string
}

func GenerateID(prefix string, barang *[]Barang) string {
	newId := ""
	for i, b := range *barang {
		id := strings.Split(b.Id, "BAR-")[1]
		idInt, _ := strconv.Atoi(id)
		if idInt != (i + 1) {
			newId = fmt.Sprintf("BAR-%v", i)
		}
	}

	if newId == "" {
		newId = fmt.Sprintf("%v-%v", prefix, len(*barang)+1)

	}
	return newId
}
