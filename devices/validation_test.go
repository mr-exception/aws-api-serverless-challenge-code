package main

import (
	"testing"
)

func TestValidationInput1(t *testing.T) {
	var rules = make([]Rule, 1)
	rules = append(rules, Rule{Field: "string", Pattern: "exists", Message: "string is required"})

	var inputs = make(map[string]string)
	inputs["string"] = "ali"

	success, message := validateInput(inputs, rules)

	if !success {
		t.Errorf("failed in validation, output: %s", message)
	}
}
