package helpers

import (
	//"fmt"
	"strings"
    "reflect"
    "github.com/asaskevich/govalidator"
)

func ValidationError(s interface{}, err error) map[string]string {
    switch err.(type) {
    case govalidator.Errors:
        // Use reflect to get the raw struct element
        typ := reflect.TypeOf(s).Elem()
        if typ.Kind() != reflect.Struct {
            return nil
        }

        // This is will contain the errors we return back to user
        errs := map[string]string{}
        // Errors found by the validator
        errsByField := govalidator.ErrorsByField(err.(govalidator.Errors))
        // Loop over our struct fields
        for i := 0; i < typ.NumField(); i++ {
            // Get the field
            f := typ.Field(i)
            // Do we have an error for the field
            e, ok := errsByField[f.Name]
            if ok {
                // Try and get the `json` struct tag
                name := strings.Split(f.Tag.Get("json"), ",")[0]
                // If the name is - we should ignore the field
                if name == "-" {
                    continue
                }
                // If the name is not blank we add it our error map
                if name != "" {
                    errs[name] = e
                    continue
                }
                // Finall if all else has failed just add the raw field name to the
                // error map
                errs[CamelCaseToSnakeCase(f.Name)] = e
            }
        }

        // Return the validation error
        return errs
	}
	
	return nil
}