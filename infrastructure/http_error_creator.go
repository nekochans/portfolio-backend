package infrastructure

import Openapi "github.com/nekochans/portfolio-backend/openapi"

type HttpErrorCreator struct{}

func (h *HttpErrorCreator) CreateFromMsg(msg string) Openapi.Error {
	const notFoundErrorCode = 404
	const internalServerErrorCode = 500

	code := internalServerErrorCode
	message := "Internal Server Error"

	m := map[string]int{
		"MysqlMemberRepository.Find: Member Not Found":             notFoundErrorCode,
		"MysqlMemberRepository.FindAll: Members Not Found":         notFoundErrorCode,
		"MysqlWebServiceRepository.FindAll: WebServices Not Found": notFoundErrorCode,
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
