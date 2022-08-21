package error

import "fmt"

/* エラーメッセージ */

// EnvNotFoundError 環境変数存在なし
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
	BeforeEnv string
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

// DuplicateError レコードなしエラー
type DuplicateError struct {
	TableName string
	Key       string
}

// Error DuplicateErrorのエラーメッセージ
func (e *DuplicateError) Error() string {
	return fmt.Sprintf(MSG_E00006, e.TableName, e.Key)
}

// Code DuplicateErrorのエラーコード
func (e *DuplicateError) Code() string {
	return CODE_E00006
}

// DBConnectionOpenError DB接続エラー
type DBConnectionOpenError struct {
	Db      string
	Address string
	Port    string
	Err     error
}

// Error DBConnectionOpenErrorのエラーメッセージ
func (e *DBConnectionOpenError) Error() string {
	return fmt.Sprintf(MSG_E00007, e.Db, e.Address, e.Port, e.Err)
}

// Code DBConnectionOpenErrorのエラーコード
func (e *DBConnectionOpenError) Code() string {
	return CODE_E00007
}

// ServeFatalError サーバー起動失敗
type ServeFatalError struct {
}

// Error ServeFatalErrorのエラーメッセージ
func (e *ServeFatalError) Error() string {
	return MSG_E00008
}

// Code ServeFatalErrorのエラーコード
func (e *ServeFatalError) Code() string {
	return CODE_E00008
}

/* 警告メッセージ */
// EnvNotFoundWarn 環境変数取得失敗
type EnvNotFoundWarn struct {
	Key          string
	DefaultValue interface{}
}

// Error EnvNotFoundWarnのエラーメッセージ
func (e *EnvNotFoundWarn) Error() string {
	return fmt.Sprintf(MSG_W00001, e.Key, e.DefaultValue)
}

// Code EnvNotFoundWarnのエラーコード
func (e *EnvNotFoundWarn) Code() string {
	return CODE_W00002
}

// EnvConvertWarn 環境変数変換失敗
type EnvConvertWarn struct {
	Key          string
	BeforeEnv    string
	DefaultValue interface{}
}

// Error EnvConvertWarnのエラーメッセージ
func (e *EnvConvertWarn) Error() string {
	return fmt.Sprintf(MSG_W00002, e.Key, e.BeforeEnv, e.DefaultValue)
}

// Code EnvConvertWarnのエラーコード
func (e *EnvConvertWarn) Code() string {
	return CODE_W00002
}
