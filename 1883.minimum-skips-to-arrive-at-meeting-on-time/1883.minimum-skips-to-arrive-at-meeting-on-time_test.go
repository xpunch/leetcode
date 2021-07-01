package xcode

import "testing"

func Test1(t *testing.T) {
	dist, speed, hoursBefore := []int{1, 3, 2}, 4, 2
	if output := minSkips(dist, speed, hoursBefore); output != 1 {
		t.Fatal(output)
	}
}

func Test2(t *testing.T) {
	dist, speed, hoursBefore := []int{7, 3, 5, 5}, 2, 10
	if output := minSkips(dist, speed, hoursBefore); output != 2 {
		t.Fatal(output)
	}
}
func Test3(t *testing.T) {
	dist, speed, hoursBefore := []int{7, 3, 5, 5}, 1, 10
	if output := minSkips(dist, speed, hoursBefore); output != -1 {
		t.Fatal(output)
	}
}
