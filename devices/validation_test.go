package main

import (
	"testing"
)

// successfull test
func TestValidationInput0(t *testing.T) {
	var rules = make([]Rule, 1)
	rules = append(rules, Rule{Field: "string", Pattern: "exists", Message: "string is required"})

	var inputs = make(map[string]string)
	inputs["string"] = "ali"

	success, message := validateInput(inputs, rules)

	if !success {
		t.Errorf("failed in validation, output: %s", message)
	}
}

// failed test
func TestValidationInput1(t *testing.T) {
	var rules = make([]Rule, 1)
	rules = append(rules, Rule{Field: "string", Pattern: "exists", Message: "string is required"})

	var inputs = make(map[string]string)
	inputs["string"] = ""

	success, message := validateInput(inputs, rules)

	if success {
		t.Errorf("failed in validation, output: %s", message)
	}
}

func TestValidationInput2(t *testing.T) {
	var rules = make([]Rule, 1)
	rules = append(rules, Rule{Field: "string", Pattern: "isEmpty", Message: "string is not required"})

	var inputs = make(map[string]string)
	inputs["string"] = ""

	success, message := validateInput(inputs, rules)

	if !success {
		t.Errorf("failed in validation, output: %s", message)
	}
}

// failed test
func TestValidationInput3(t *testing.T) {
	var rules = make([]Rule, 1)
	rules = append(rules, Rule{Field: "string", Pattern: "isEmpty", Message: "string is not required"})

	var inputs = make(map[string]string)
	inputs["string"] = ""

	success, message := validateInput(inputs, rules)

	if !success {
		t.Errorf("failed in validation, output: %s", message)
	}
}

// failed test => testing input orders
func TestValidationInput4(t *testing.T) {
	var rules = make([]Rule, 1)
	rules = append(rules, Rule{Field: "a", Pattern: "isEmpty", Message: "a is not required"})
	rules = append(rules, Rule{Field: "b", Pattern: "isEmpty", Message: "b is not required"})

	var inputs = make(map[string]string)
	inputs["a"] = "1"
	inputs["b"] = "2"

	success, message := validateInput(inputs, rules)

	if success && message != "a is not required" {
		t.Errorf("failed in validation, output: %s", message)
	}
}
