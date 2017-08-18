package weightconv

import "fmt"

type Pound float64
type Kilogram float64

func (c Pound) String() string {
	return fmt.Sprintf("%gP", c)
}

func (f Kilogram) String() string {
	return fmt.Sprintf("%gK", f)
}
