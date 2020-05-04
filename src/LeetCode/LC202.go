package LeetCode

import "strconv"

func IsHappy(n int) bool {
	if n <= 0 {
		return false
	}else if n > 1 {
		list := make([]int, 0)
		for {
			m := strconv.Itoa(n)
			mLen := len(m)
			mSum := 0
			for i := 0; i<mLen; i++{
				mInt, _ := strconv.Atoi(m[i:i+1])
				mSum = mInt * mInt + mSum
			}
			if mSum == 1 {
				return true
			}else{
				for item := range list{
					if mSum == item {
						return false
					}
				}
				list = append(list, mSum)
			}
			n = mSum
		}
	}
	return true
}