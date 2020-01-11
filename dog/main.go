//Package dog takes human years and turns them into dog years (1 human year = 7 dog years)
package dog

//Years takes an integer of human years and returns an integer of dog years
func Years(i int) int {
	dogYears := i * 7
	return dogYears
}
