package test

type data struct {
	x int
	Y int
}

func CreateData() *data {
	return new(data)
}
