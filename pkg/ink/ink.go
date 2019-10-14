package ink

const (
	Red    ColorFamily = iota
	Blue   ColorFamily = iota
	Green  ColorFamily = iota
	Black  ColorFamily = iota
	Yellow ColorFamily = iota
	Orange ColorFamily = iota
	Brown  ColorFamily = iota
)

type ColorFamily int

type Ink struct {
	Id          int64 `gorm:"PRIMARY_KEY"`
	Name        string
	ColorFamily ColorFamily
}
