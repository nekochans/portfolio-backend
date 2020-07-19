package infrastructure

type HTTPErrorCreator struct{}

func (c *HTTPErrorCreator) CreateFromMsg(msg string) HTTPError {
	code := 500
	message := "Internal Server Error"

	m := map[string]int{
		"MysqlMemberRepository.Find: Member Not Found":             404,
		"MysqlMemberRepository.FindAll: Members Not Found":         404,
		"MysqlWebServiceRepository.FindAll: WebServices Not Found": 404,
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
