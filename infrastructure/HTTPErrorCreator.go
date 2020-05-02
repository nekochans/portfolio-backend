package infrastructure

type HTTPErrorCreator struct{}

func (c *HTTPErrorCreator) CreateFromMsg(msg string) HTTPError {
	code := 500
	message := "Internal Server Error"

	m := map[string]int{
		"MySQLMemberRepository.Find: Member Not Found":             400,
		"MySQLMemberRepository.FindAll: Members Not Found":         400,
		"MySQLWebServiceRepository.FindAll: WebServices Not Found": 400,
	}

	if m[msg] != 0 {
		code = m[msg]
		message = msg
	}

	he := HTTPError{
		Code:    code,
		Message: message,
	}

	return he
}
