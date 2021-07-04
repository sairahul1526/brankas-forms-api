package logger

import (
	"fmt"
	CONFIG "forms-api/config"
)

// Log - log based on test value
func Log(str ...interface{}) {
	if CONFIG.Log {
		fmt.Println(str...)
	}
}
