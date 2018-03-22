func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0
	for left <= right {
		h := min(height[left], height[right])
		w := right - left
		if maxWater < w * h {
			maxWater = w *h
		}
		if height[left] < height[right]{
			left ++
		}else{
			right --
		}
	}
	return maxWater
}

func min(a, b int) int {
	if a > b {
		return b
	}else{
		return a
	}
}