package contohpack

func PublicFunction() string {
	return "This is a public function from contohpack"
}

func PrivateFunction() string {
	return "This is a private function from contohpack"
}

func FunctionWithSingleParam(param string) string {
	return "This function has a single parameter: " + param
}

func FunctionWithMultipleParams(param1 string, param2 int, isOk bool) string {
	if isOk {
		return "This function has multiple parameters: " + param1 + " and " + string(param2) + " and isOk: true"
	} else {
		return "This function has multiple parameters: " + param1 + " and " + string(param2) + " and isOk: false"
	}
}

func FunctionWithVariadicParams(params ...string) string {
	result := "This function has variadic parameters: "
	for _, param := range params {
		result += param + " "
	}
	return result
}

func FunctionWithReturnValues() (string, int) {
	return "This function returns a string and an int", 42
}

func FunctionWithNamedReturnValues() (result string, count int) {
	result = "This function returns named values"
	count = 10
	return
}
