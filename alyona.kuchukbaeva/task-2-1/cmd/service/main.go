package main

const (
	minTemp = 15
	maxTemp = 30
)

type conditioner struct {
	low  int
	high int
}

func newConditioner() conditioner {
	return conditioner{low: minTemp, high: maxTemp}
}

func main() {

}
