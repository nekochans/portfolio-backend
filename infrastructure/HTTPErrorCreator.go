package infrastructure

type HTTPErrorCreator struct{}

func (c *HTTPErrorCreator) CreateFromMsg(msg string) HTTPError {
	code := 500
	message := "Internal Server Error"

	m := map[string]int{
		"MySQLMemberRepository.Find: DB.Prepare Error":     500,
		"MySQLMemberRepository.Find: stmt.Query Error":     500,
		"MySQLMemberRepository.Find: rows.Scan Error":      500,
		"MySQLMemberRepository.Find: Member Not Found":     400,
		"MySQLMemberRepository.FindAll: DB.Prepare Error":  500,
		"MySQLMemberRepository.FindAll: stmt.Query Error":  500,
		"MySQLMemberRepository.FindAll: rows.Scan Error":   500,
		"MySQLMemberRepository.FindAll: Members Not Found": 400,
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
