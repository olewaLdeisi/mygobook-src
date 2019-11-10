package main

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

type X struct {
	Name string
}

type Y struct {
	X
	Name string
}

func main() {
	/*
		var a Integer = 1
		fmt.Println(a)
		a.Less(2)
		fmt.Println(a)
	*/

	/*
		rect1 := new(Rect)
		rect2 := &Rect{}
		rect3 := &Rect{0, 0, 100, 200}
		rect4 := &Rect{
			width:  100,
			height: 200,
		}
	*/

	/*
		a := &Y{
			X:    X{Name: "XName"},
			Name: "YName",
		}
		fmt.Println(a.Name, a.X.Name)
	*/

	/*
		var a Integer = 1
		var b LessAdder = &a // 自动生成 func (a *Integer) Less(b Integer) bool{return (*a).Less(b)}
	*/
}
