//go:build !windows
// +build !windows

package utils

// IsAdmin は現在のプロセスが管理者権限で実行されているかどうかを確認します（Windows専用）
func IsAdmin() bool {
	return false
}

