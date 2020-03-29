package infrastructure

type HTTPErrorCreator struct{}

func (c *HTTPErrorCreator) CreateFromMsg(msg string) HTTPError {
	code := 500
	message := "Internal Server Error"

	m := map[string]int{
		"MySQLMemberRepository.FindAll: DB.Prepare Error": 500,
		"MySQLMemberRepository.FindAll: stmt.Query Error": 500,
		"MySQLMemberRepository.FindAll: rows.Scan Error":  500,
		"MySQLMemberRepository.FindAll: Not Found Error":  400,
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
