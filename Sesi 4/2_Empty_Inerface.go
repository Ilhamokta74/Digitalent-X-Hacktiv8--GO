package main

func EmptyInterface() {
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Sudirman"

	randomValues = 28

	randomValues = true

	randomValues = []string{"Aralie", "Nanda"}
}

func EmptyTypeAssertion() {
	var v interface{}

	v = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}
}

func EmptyMapSlice() {
	rs := []interface{}{1, "Aralie", true, 2, "Amanda", true}

	rm := map[string]interface{}{
		"Name":   "Aralie",
		"Status": true,
		"Age":    23,
	}

	_, _ = rs, rm
}

func main() {
	// EmptyInterface()
	// EmptyTypeAssertion()
	// EmptyMapSlice()
}
