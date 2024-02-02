package main

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() {
	return n.Message
}

func SaveData(id string, data interface{}) error {
	if id == "" {
		return &validationError{"validaton error"}
	}
	if id != "Alif" {
		return &notFoundError{"data not found"}
	}
}

func main() {

}
