package main

import (
	"encoding/json"
)

// Rule is a single validation rule defined for a parameter
type Rule struct {
	Field   string // field name
	Pattern string // validation function (isEmpty, exists)
	Message string // validation message is case of failure
}

// validateRequest validates the input data received from Request
// @param Request
// @param []Rule
// @return bool (validation passed) and Response (error response, in case of validation failure)
func validateRequest(request Request, rules []Rule) (bool, Response) {
	var inputs map[string]string
	json.Unmarshal([]byte(request.Body), &inputs)

	// validate inputs
	var success, message = validateInput(inputs, rules)

	if !success { // in case of validation failure, returns the failure message
		errorResponse := ErrorResponse{
			Message: message,
		}
		result, _ := json.Marshal(errorResponse)
		return false, Response{Body: string(result), StatusCode: 400}
	}
	// in case of validate success returns an empty message
	return true, Response{}
}

func validateInput(inputs map[string]string, rules []Rule) (bool, string) {
	for _, rule := range rules { // execute validation rules of input fields
		switch rule.Pattern {
		case "isEmpty": // field have to be empty|nil
			if inputs[rule.Field] != "" {
				return false, rule.Message
			}
			break
		case "exists": // file have to be a valid string|int
			if inputs[rule.Field] == "" {
				return false, rule.Message
			}
			break
		}
	}
	return true, ""
}
