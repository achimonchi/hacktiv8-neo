package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCubeVolume(t *testing.T) {
	cube := NewCube(10)
	expected := 1_000.0

	require.Equal(t, cube.Volume(), expected)
}
func TestCubeLuas(t *testing.T) {
	cube := NewCube(10)
	expected := 600.0

	require.Equal(t, cube.Luas(), expected)
}
func TestCubeKeliling(t *testing.T) {
	cube := NewCube(10)
	expected := 120.0

	require.Equal(t, cube.Keliling(), expected)
}

func init() {

	// runtime.GOMAXPROCS(2)
}
func BenchmarkVolumePtr(b *testing.B) {
	cube := NewCubePtr(10)
	for i := 0; i < b.N; i++ {
		cube.VolumePtr()
	}

}
func BenchmarkVolume(b *testing.B) {
	cube := NewCube(10)
	for i := 0; i < b.N; i++ {
		cube.Volume()
	}
}

func BenchmarkLuasPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCubePtr(10).LuasPtr()
	}

}

func BenchmarkLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCube(10).Luas()
	}
}
