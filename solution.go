package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type CustomInt int

const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0

func CalcSquare(sideLen float64, sidesNum CustomInt) float64 {
	var result float64
	sideLenQ := sideLen * sideLen

	switch sidesNum {
	case 3:
		result = math.Sqrt(3) / 4 * sideLenQ
	case 4:
		result = sideLenQ
	case 0:
		result = math.Pi * sideLenQ
	}

	return result
}
