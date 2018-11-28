/*
package example:
tempconv performs Celsius Fahrenheit conversions

to make your package is working, you need project structure like
GOPATH/yourpackagefolder/yourpackagefiles
because go is looking for there for packages.
*/

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC = -273.15
	BoilingC      = 100
	FreezingC     = 0
)

func (c Celsius) String() (str string) {
	str = fmt.Sprintf("%g°C", c)
	return
}

func (f Fahrenheit) String() (str string) {
	str = fmt.Sprintf("%g°C", f)
	return
}
