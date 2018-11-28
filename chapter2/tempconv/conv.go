package tempconv

/*
(c * 9/5 + 32)
converts celsius to fahrenheit
*/
func CtoF(c Celsius) (f Fahrenheit) {
	f = Fahrenheit(c*9/5 + 32)
	return
}

/*
((f - 32) * 5 / 9)
convert fahrenheit to celsius
*/
func FtoC(f Fahrenheit) (c Celsius) {
	c = Celsius((f - 32) * 5 / 9)
	return
}
