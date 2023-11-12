package exercise

// return the nearest num^2 to x
func NumWithClosestSquareTo(x float64) float64 {
	const INCREMENT = 0.01

	num := 0.0

	// do this within a loop
	// get the squared number
	// check if its greater than x
	// if it is then return previous num
	// else continue
	for proceed := true; proceed; {
		if square := num * num; square > x {
			proceed = false
			num -= INCREMENT
		} else if square == x {
			proceed = false
		} else {
			num += INCREMENT
		}
	}
	
	return num
}