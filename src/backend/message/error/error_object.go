package error

import "fmt"

// EnvNotFoundError 環境変数取得失敗
type EnvNotFoundError struct {
	Key string
}

// Error EnvNotFoundErrorのエラーメッセージ
func (e *EnvNotFoundError) Error() string {
	return fmt.Sprintf(MSG_E00001, e.Key)
}

// Code EnvNotFoundErrorのエラーコード
func (e *EnvNotFoundError) Code() string {
	return CODE_E00001
}

// EnvConvertError 環境変数変換失敗
type EnvConvertError struct {
	Key       string
	BeforeEnv interface{}
}

// Error EnvConvertErrorのエラーメッセージ
func (e *EnvConvertError) Error() string {
	return fmt.Sprintf(MSG_E00002, e.Key, e.BeforeEnv)
}

// Code EnvConvertErrorのエラーコード
func (e *EnvConvertError) Code() string {
	return CODE_E00002
}

// SQLExecError SQL実行エラー
type SQLExecError struct {
	TableName string
	Err       error
}

// Error SQLExecErrorのエラーメッセージ
func (e *SQLExecError) Error() string {
	return fmt.Sprintf(MSG_E00003, e.TableName, e.Err)
}

// Code SQLExecErrorのエラーコード
func (e *SQLExecError) Code() string {
	return CODE_E00003
}

// ExistsError 存在チェックエラー
type ExistsError struct {
	TableName  string
	Conditions string
}

// Error ExistsErrorのエラーメッセージ
func (e *ExistsError) Error() string {
	return fmt.Sprintf(MSG_E00004, e.TableName, e.Conditions)
}

// Code ExistsErrorのエラーコード
func (e *ExistsError) Code() string {
	return CODE_E00004
}

// NotFoundError レコードなしエラー
type NotFoundError struct {
	TableName  string
	Conditions string
}

// Error NotFoundErrorのエラーメッセージ
func (e *NotFoundError) Error() string {
	return fmt.Sprintf(MSG_E00005, e.TableName, e.Conditions)
}

// Code NotFoundErrorのエラーコード
func (e *NotFoundError) Code() string {
	return CODE_E00005
}

// CreateError レコードなしエラー
type CreateError struct {
	TableName string
	Id        string
}

// Error CreateErrorのエラーメッセージ
func (e *CreateError) Error() string {
	return fmt.Sprintf(MSG_E00006, e.TableName, e.Id)
}

// Code CreateErrorのエラーコード
func (e *CreateError) Code() string {
	return CODE_E00006
}
