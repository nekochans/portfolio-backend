// Package Openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package Openapi

// Defines values for ErrorCode.
const (
	ErrorCodeN400 ErrorCode = 400

	ErrorCodeN404 ErrorCode = 404

	ErrorCodeN500 ErrorCode = 500

	ErrorCodeN503 ErrorCode = 503
)

// エラーモデル
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

// ErrorCode defines model for Error.Code.
type ErrorCode int

// メンバーモデル
type Member struct {
	CvUrl          string `json:"cvUrl"`
	GithubPicture  string `json:"githubPicture"`
	GithubUserName string `json:"githubUserName"`
	Id             int64  `json:"id"`
}

// WebService defines model for WebService.
type WebService struct {
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Url         string `json:"url"`
}
