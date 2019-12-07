package t

// Max 获取最大的数
func Max(args ...interface{}) int {
	if len(args) == 0 {
		return 0
	}
	var max = New(args[0]).Int()
	for k, arg := range args {
		if k == 0 {
			continue
		}
		tmp := New(arg).Int()
		if tmp > max {
			max = tmp
		}
	}
	return max
}

// Min 获取最小的数
func Min(args ...interface{}) int {
	if len(args) == 0 {
		return 0
	}
	var min = New(args[0]).Int()
	for k, arg := range args {
		if k == 0 {
			continue
		}
		tmp := New(arg).Int()
		if tmp < min {
			min = tmp
		}
	}
	return min
}

