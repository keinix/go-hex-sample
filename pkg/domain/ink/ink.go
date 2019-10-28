package ink

const (
	Red    ColorFamily = iota + 1
	Blue   ColorFamily = iota + 1
	Green  ColorFamily = iota + 1
	Black  ColorFamily = iota + 1
	Yellow ColorFamily = iota + 1
	Orange ColorFamily = iota + 1
	Brown  ColorFamily = iota + 1
)

type ColorFamily int

type Ink struct {
	Id          int64       `gorm:"PRIMARY_KEY" json:"id"`
	Name        string      `json:"name"`
	ColorFamily ColorFamily `json:"colorFamily"`
}
