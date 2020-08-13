package calculator

func Calculate(operation Operation) Operation {
	switch operation.TypeOf {
	case "sum":
		operation.Result = sum(operation.First, operation.Second)

	case "sub":
		operation.Result = sub(operation.First, operation.Second)

	case "divide":
		operation.Result = divide(operation.First, operation.Second)
	default:
	}
	return operation
}
func sum(first float64, second float64) float64 {
	return first + second
}
func sub(first float64, second float64) float64 {
	return first - second
}
func divide(first float64, second float64) float64 {
	return first / second
}
