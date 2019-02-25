package jsonw

import (
	"encoding/json"
	"fmt"
)

// UnMarshalError have default error and detail syntax info
type UnMarshalError struct {
	err    error
	syntax string
}

// Implement error interface
func (u UnMarshalError) Error() string {
	return fmt.Sprintf("error:%s\ncause:%s", u.err, u.syntax)
}

// Syntax return sytanx info
func (u UnMarshalError) Syntax() string {
	return u.syntax
}

// Unmarshal json
func Unmarshal(b []byte, i interface{}) error {
	if err := json.Unmarshal(b, i); err != nil {
		if err, ok := err.(*json.SyntaxError); ok {
			return UnMarshalError{
				err:    err,
				syntax: string(b[err.Offset-15 : len(b)]),
			}
		}
		return UnMarshalError{
			err: err,
		}
	}
	return nil
}
