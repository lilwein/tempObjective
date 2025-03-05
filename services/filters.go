package services

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"
	"gorm.io/gorm"
)

// Operators
const (
	equal     = "="
	less      = "<"
	less_e    = "<="
	greater   = ">"
	greater_e = ">="
	in        = "in"
	like      = "LIKE"
	ilike     = "ILIKE"
	like_     = "like"
	ilike_    = "ilike"
)

// var operators = []string{
// 	equal,
// 	less,
// 	less_e,
// 	greater,
// 	greater_e,
// 	in,
// 	like,
// 	ilike,
// }

var operators = map[string]string{
	"eq": equal,
}

// Builds the Where query applying the filters, runs it with GORM and returns the result as *gorm.DB and nil.
// If whereMap contains an invalid operator as key, returns error.
// Operators: =, >, <, >=, <=, in, ilike, like
func WhereResult(s any, db *gorm.DB) *core.ApplicationError {

	// Using reflect to work on struct s
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return core.TechnicalErrorWithCodeAndMessage("NOT-STRUCT", "input is not a struct")
	}
	t := v.Type()

	// For each field of struct s
	for i := range v.NumField() {

		// Tag "filter"
		tag := t.Field(i).Tag.Get("filter")

		// If tag == "" then do nothing
		if tag != "" {

			// Get value of field
			var value string
			if str, ok := v.Field(i).Interface().(string); ok {
				value = str
			} else {
				value = fmt.Sprintf("%v", v.Field(i).Interface())
			}

			// If value == "" then do nothing
			if err := whereResult(db, tag, value); err != nil {
				return err
			}

		}
	}

	return nil
}

func whereResult(db *gorm.DB, qp, value string) *core.ApplicationError {

	// Correct Query Parameter Key
	key, op := ExcludeLastUnderscore(qp)

	// Check if the operator is correct
	operator, ok := operators[op]
	if !ok {
		return core.TechnicalErrorWithCodeAndMessage("INVALID-OPERATOR", "invalid operator")
	}

	// Parameter
	param := strings.ToLower(value)

	// Build statement
	statement := fmt.Sprintf("%s %s ?", key, operator)
	// fmt.Println("statement:\t", statement)

	// Modify parameter if operator is LIKE
	if operator == "ILIKE" || operator == "LIKE" {
		param = "%" + param + "%"
	}

	// Build query with gorm
	db = db.Where(statement, param)

	return nil
}

// Returns the s string without the last token after the last underscore, if this exists
func ExcludeLastUnderscore(s string) (string, string) {
	tokens := strings.Split(s, "_")
	if len(tokens) > 1 {
		key, op := strings.Join(tokens[:len(tokens)-1], "_"), tokens[len(tokens)-1]
		fmt.Printf("key: %s\top: %s\n", key, op)

		return key, op
	}
	return "", ""
}
