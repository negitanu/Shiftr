//go:build !windows
// +build !windows

package network

import (
	"fmt"

	"github.com/shiftr/shiftr/pkg/models"
)

// NetworkError はネットワーク関連のエラーを表します
type NetworkError struct {
	Code    string
	Message string
	Err     error
}

func (e *NetworkError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func unsupported() error {
	return &NetworkError{
		Code:    "UNSUPPORTED_PLATFORM",
		Message: "この機能はWindowsでのみ利用できます（GOOS=windows でビルドしてください）",
	}
}

// GetNICList は利用可能なNICのリストを取得します（Windows専用）
func GetNICList() ([]string, error) {
	return nil, unsupported()
}

// GetCurrentIPConfig は指定されたNICの現在のIP設定を取得します（Windows専用）
func GetCurrentIPConfig(nicName string) (*models.Profile, error) {
	return nil, unsupported()
}

// ApplyProfile はプロファイルの設定をNICに適用します（Windows専用）
func ApplyProfile(profile *models.Profile) error {
	return unsupported()
}

// ApplyDHCP は指定されたNICをDHCPに切り替えます（Windows専用）
func ApplyDHCP(nicName string) error {
	return unsupported()
}

