func convert(s string, numRows int) string {
    if numRows == 1 {
		return s
	} else {
		result := make([]byte, 0)
		i := 0
		for i < numRows {
			if gap, ok := gaps(i, numRows); ok {
				idx := i
				for idx < len(s) {
					result = append(result, s[idx])
					j := idx + gap
					if j < len(s) {
						result = append(result, s[j])
					}
					idx += (2*numRows - 2)
				}
			} else {
				idx := i
				for idx < len(s) {
					result = append(result, s[idx])
					idx += (2*numRows - 2)
				}
			}
			i++
		}
		return string(result)
	}
}

func gaps(i, numRows int) (int, bool) {
	gap := 2 * numRows  - 2
	if i==0 || i == numRows -1 {
		return gap, false
	}else{
		return (numRows - 1 - i) * 2, true
	}
}