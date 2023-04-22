package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	a := make([]bool, len(candies))
	for _, x := range candies {
		if x > max {
			max = x
		}
	}
	for i, _ := range candies {
		if candies[i]+extraCandies >= max {
			a[i] = true
		}
	}
	return a
}
