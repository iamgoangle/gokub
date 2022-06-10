package strTo

import "strconv"

func ToInt32(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}

	return int32(i)
}
