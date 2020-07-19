package Openapi

import "fmt"

func (he *Error) Error() string {
	return fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
}
