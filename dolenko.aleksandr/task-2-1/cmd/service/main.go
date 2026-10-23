package main

const (
	MaxTemp = 30
	MinTemp = 15
)

type conditioner struct {
	low  int
	high int
}

func newConditioner() *conditioner {
	return &conditioner{low: MaxTemp, high: MinTemp}
}

func main() {

}
