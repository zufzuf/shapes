package factory

type Shape interface {
	Name() string
	Area() int64
}

type DefaultShape struct{}

func NewDefaultShape() *DefaultShape {
	return &DefaultShape{}
}

func (DefaultShape) Name() string {
	return "undefined data"
}

func (d DefaultShape) Area() int64 {
	return 0
}

type Rectangle struct {
	Long int64 `json:"long"`
	Wide int64 `json:"wide"`
}

func NewRectangle() *Rectangle {
	return &Rectangle{}
}

func (Rectangle) Name() string {
	return "persegi panjang"
}

func (r Rectangle) Area() int64 {
	return r.Long * r.Wide
}

type Square struct {
	Side int64 `json:"side"`
}

func (Square) Name() string {
	return "persegi"
}

func (s Square) Area() int64 {
	return s.Side * s.Side
}

type Triangle struct {
	Base   int64 `json:"base"`
	Height int64 `json:"height"`
}

func (Triangle) Name() string {
	return "segitiga"
}

func (t Triangle) Area() int64 {
	return (t.Base * t.Height) / 2
}
