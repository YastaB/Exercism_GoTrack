package space

type Planet string

var orbitScale = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {

	/*if math.Abs(31.69-31.70979198376458) > 0.01 {
		println("Exceeds")
	}*/
	println("Seconds: ", seconds, " Planet: ", planet)
	println("Planet Scale: ", orbitScale[planet])
	result := (seconds / (3600.0 * 8765.82 * orbitScale[planet]))
	//result = math.Floor(result*1000) / 1000
	println(result)
	return result
}
