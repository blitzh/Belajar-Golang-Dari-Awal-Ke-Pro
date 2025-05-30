package main

// Internal fungsi dalam satu package dan dalam satu file
func intPrivateFunction() string {
	return "This is a private function"
}

func intPublicFunction() string {
	return "This is a public function"
}

// fund dengan parameter single
func intFunctionWithSingleParam(param string) string {
	return "This function has a single parameter: " + param
}

// fungsi dengan parameter multiple
func intFunctionWithMultipleParams(param1 string, param2 int, isOk bool) string {
	if isOk {
		return "This function has multiple parameters: " + param1 + " and " + string(param2) + " and isOk: true"
	} else {
		return "This function has multiple parameters: " + param1 + " and " + string(param2) + " and isOk: false"
	}
}

func main() {

	// panggil fungsi internal disini

	// pangiil fungsi dari pack contohpack disini

}
