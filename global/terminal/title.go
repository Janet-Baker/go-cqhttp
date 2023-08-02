//go:build !windows

package terminal

import (
	"fmt"
	"time"
)

// SetTitle 设置标题为 go-cqhttp `版本` `版权`
func SetTitle() {
	fmt.Printf("\033]0;go-cqhttp "+" © 2020 - %d Mrs4s"+"\007", time.Now().Year())
}
