package helper

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
