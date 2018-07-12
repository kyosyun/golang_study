package tempconv


func KtoC(k Kelvin) Celsius {
	return Celsius(k + 273.15)
}

func CtoK(c Celsius) Kelvin {
	return Kelvin(c - 273.15)
}

func KtoF(k Kelvin) Fahrenheit {
	return CtoF(KtoC(k))
}

func FtoK(f Fahrenheit) Kelvin {
	return CtoK(FtoC(f))
}

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}


func FtoC(f Fahrenheit) Celsius {
	return Celsius((f-32) * 5 /9 )
}