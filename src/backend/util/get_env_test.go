package util_test

import (
	"os"
	"testing"

	assert "github.com/go-playground/assert/v2"

	pomodork_error "pomodork-backend/message/error"
	"pomodork-backend/util"
)

type GetRequiredEnvInput struct {
	key     string
	envType string
	env     string
}

type GetRequiredEnvOutput struct {
	env interface{}
	err error
}

func TestGetRequiredEnv(t *testing.T) {
	cases := []struct {
		name     string
		input    GetRequiredEnvInput
		expected GetRequiredEnvOutput
	}{
		{
			name:     "【正常_001】入力「key:TEST_ENV_KEY_001、envType:string、env:TEST_ENV_001」 期待値「env:TEST_ENV_001、err:nil」",
			input:    GetRequiredEnvInput{key: "TEST_ENV_KEY_001", envType: "string", env: "TEST_ENV_001"},
			expected: GetRequiredEnvOutput{env: "TEST_ENV_001", err: nil},
		},
		{
			name:     "【正常_002】入力「key:TEST_ENV_KEY_002、envType:int、env:1000」 期待値「env:1000、err:nil」",
			input:    GetRequiredEnvInput{key: "TEST_ENV_KEY_002", envType: "int", env: "1000"},
			expected: GetRequiredEnvOutput{env: 1000, err: nil},
		},
		{
			name:     "【異常_001】入力「key:TEST_ENV_KEY_003、envType:string、env:空」 期待値「env:空、err:EnvNotFoundError」",
			input:    GetRequiredEnvInput{key: "TEST_ENV_KEY_003", envType: "string", env: ""},
			expected: GetRequiredEnvOutput{env: "", err: &pomodork_error.EnvNotFoundError{Key: "TEST_ENV_KEY_003"}},
		},
		{
			name:     "【異常_002】入力「key:TEST_ENV_KEY_004、envType:int、env:TEST_ENV_004」 期待値「env:空、err:EnvConvertError」",
			input:    GetRequiredEnvInput{key: "TEST_ENV_KEY_004", envType: "int", env: "TEST_ENV_004"},
			expected: GetRequiredEnvOutput{env: "", err: &pomodork_error.EnvConvertError{Key: "TEST_ENV_KEY_004", BeforeEnv: "TEST_ENV_004"}},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			os.Setenv(testCase.input.key, testCase.input.env)
			env, err := util.GetRequiredEnv(testCase.input.key, testCase.input.envType)
			assert.Equal(t, testCase.expected.env, env)
			assert.Equal(t, testCase.expected.err, err)
		})
	}
}

type GetNonRequiredEnvInput struct {
	key          string
	envType      string
	env          string
	defaultValue interface{}
}

type GetNonRequiredEnvOutput struct {
	env interface{}
	err error
}

func TestGetNonRequiredEnv(t *testing.T) {
	cases := []struct {
		name     string
		input    GetNonRequiredEnvInput
		expected GetNonRequiredEnvOutput
	}{
		{
			name:     "【正常_001】入力「key:TEST_ENV_KEY_001、envType:string、env:TEST_ENV_001、defaultValue:TEST_ENV_DEFAULT_001」 期待値「env:TEST_ENV_001、err:nil」",
			input:    GetNonRequiredEnvInput{key: "TEST_ENV_KEY_001", envType: "string", env: "TEST_ENV_001", defaultValue: "TEST_ENV_DEFAULT_001"},
			expected: GetNonRequiredEnvOutput{env: "TEST_ENV_001", err: nil},
		},
		{
			name:     "【正常_002】入力「key:TEST_ENV_KEY_002、envType:int、env:1000、defaultValue:2000」 期待値「env:1000、err:nil」",
			input:    GetNonRequiredEnvInput{key: "TEST_ENV_KEY_002", envType: "int", env: "1000", defaultValue: "2000"},
			expected: GetNonRequiredEnvOutput{env: 1000, err: nil},
		},
		{
			name:     "【異常_001】入力「key:TEST_ENV_KEY_003、envType:string、env:空、defaultValue:TEST_ENV_KEY_DEFAULT_003」 期待値「env:TEST_ENV_KEY_DEFAULT_003、err:nil」",
			input:    GetNonRequiredEnvInput{key: "TEST_ENV_KEY_003", envType: "string", env: "", defaultValue: "TEST_ENV_KEY_DEFAULT_003"},
			expected: GetNonRequiredEnvOutput{env: "TEST_ENV_KEY_DEFAULT_003", err: nil},
		},
		{
			name:     "【異常_002】入力「key:TEST_ENV_KEY_004、envType:int、env:TEST_ENV_004, defaultValue:4000」 期待値「env:4000、err:nil」",
			input:    GetNonRequiredEnvInput{key: "TEST_ENV_KEY_004", envType: "int", env: "TEST_ENV_004", defaultValue: 4000},
			expected: GetNonRequiredEnvOutput{env: 4000, err: nil},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			os.Setenv(testCase.input.key, testCase.input.env)
			env, err := util.GetNonRequiredEnv(testCase.input.key, testCase.input.envType, testCase.input.defaultValue)
			assert.Equal(t, testCase.expected.env, env)
			assert.Equal(t, testCase.expected.err, err)
		})
	}
}
