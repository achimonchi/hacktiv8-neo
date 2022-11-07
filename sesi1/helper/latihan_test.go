package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIsPrime_Success(t *testing.T) {
	expect := true
	actual := CheckIsPrime(19)

	if expect != actual {
		t.Errorf("fail! expect:%v got:%v", expect, actual)
	}
}
func TestCheckIsPrime_Fail(t *testing.T) {
	expect := false
	actual := CheckIsPrime(21)

	if expect != actual {
		t.Errorf("fail! expect:%v got:%v", expect, actual)
	}
}

func TestGenerateID_Success(t *testing.T) {
	barangs := []Barang{
		{
			Id:   "BAR-1",
			Name: "Barang 1",
		},
		{
			Id:   "BAR-2",
			Name: "Barang 2",
		},
	}

	expect := "BAR-3"
	actual := GenerateID("BAR", &barangs)

	if expect != actual {
		t.Errorf("fail! expect:%v got:%v", expect, actual)
	}
}

func TestGenerateID_Missng2Success(t *testing.T) {
	barangs := []Barang{
		{
			Id:   "BAR-1",
			Name: "Barang 1",
		},
		{
			Id:   "BAR-3",
			Name: "Barang 3",
		},
		{
			Id:   "BAR-4",
			Name: "Barang 4",
		},
	}

	expect := "BAR-2"
	actual := GenerateID("BAR", &barangs)
	assert.Equal(t, expect, actual)
}
