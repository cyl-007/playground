package settings

var fault = 0

func GetFault() (code int, has bool) {
	code = fault
	has = fault != 0

	return
}

func SetFault(code int) {
	fault = code
}
