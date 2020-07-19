package infrastructure

import Openapi "github.com/nekochans/portfolio-backend/openapi"

type OpenApiErrorCreator struct{}

func (c *OpenApiErrorCreator) CreateFromMsg(msg string) Openapi.Error {
	code := 500
	message := "Internal Server Error"

	m := map[string]int{
		"MysqlMemberRepository.Find: Member Not Found":     404,
		"MysqlMemberRepository.FindAll: Members Not Found": 404,
		"MysqlWebServiceRepository.FindAll: WebServices Not Found": 404,
	}

	if m[msg] != 0 {
		code = m[msg]
		message = msg
	}

	he := Openapi.Error{
		Code:    code,
		Message: message,
	}

	return he
}
