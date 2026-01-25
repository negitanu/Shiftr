//go:build !windows
// +build !windows

package systray

import "fmt"

// Run はシステムトレイアプリケーションを起動します（Windows専用）
func Run() error {
	return fmt.Errorf("Shiftr はWindows専用です。GOOS=windows でビルドしてください")
}

