package control

func IfOnly(age int) string {
	if age >= 18 {
		return "成年了"
	}

	return "他还是一个孩子"
}

func IfElse(age int) string {
	if age >= 18 {
		return "成年了"
	} else {
		return "他还是一个孩子"
	}
}

func IfElseIf(age int) string {
	if age >= 18 {
		return "成年了"
	} else if age >= 12 {
		return "骚年"
	} else {
		return "他还是一个孩子"
	}
}

func IfElseIfV1(age int) string {
	if age >= 18 {
		return "成年了"
	} else if age >= 12 {
		return "骚年"
	}
	return "他还是一个孩子"
}

func IfNewVariable(start int, end int) string {
	// 这个distance作用域只是在这个 if else 流程控制里
	if distance := start - end; distance > 100 {
		println(distance)
		return "距离太远了"
	} else {
		println(distance)
		return "距离比较近"
	}

	// 编译错误
	//println(distance)
}
