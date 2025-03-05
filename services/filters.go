package services

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"
	"github.com/rs/zerolog/log"
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
	"eq":    equal,
	"lt":    less,
	"lte":   less_e,
	"gt":    greater,
	"gte":   greater_e,
	"in":    in,
	"like":  like,
	"ilike": ilike,
	"LIKE":  like_,
	"ILIKE": ilike_,
}

// Builds the Where query applying the filters, runs it with GORM and returns the result as *gorm.DB and nil.
// If whereMap contains an invalid operator as key, returns error.
// Operators: =, >, <, >=, <=, in, ilike, like
func WhereResult(db *gorm.DB, s any) (*gorm.DB, *core.ApplicationError) {

	// Using reflect to work on struct s
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, core.TechnicalErrorWithCodeAndMessage("NOT-STRUCT", "input is not a struct")
	}
	t := v.Type()
	query := db
	// For each field of struct s
	for i := range v.NumField() {

		// Tag "filter"
		tag := t.Field(i).Tag.Get("filter")

		// If tag == "" then do nothing
		if tag != "" {
			log.Trace().Msgf("tag: %s", tag)

			// Get value of field
			var value string
			if str, ok := v.Field(i).Interface().(string); ok {
				value = str
			} else {
				value = fmt.Sprintf("%v", v.Field(i).Interface())
			}

			// If value == "" then do nothing
			if value != "" {
				var errW *core.ApplicationError
				query, errW = whereResult(query, tag, value)
				if errW != nil {
					return nil, errW
				}
			}

		}
	}

	return query, nil
}

func whereResult(db *gorm.DB, qp, value string) (*gorm.DB, *core.ApplicationError) {

	// Correct Query Parameter Key
	key, op := ExcludeLastUnderscore(qp)

	// Check if the operator is correct
	operator, ok := operators[op]
	if !ok {
		return nil, core.TechnicalErrorWithCodeAndMessage("INVALID-OPERATOR", "invalid operator")
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
	return db.Where(statement, param), nil

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
