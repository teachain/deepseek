package log

import "testing"

func TestDefaultLogger(t *testing.T) {
	DefaultHandler(4)
	Trace("trace", "tip", "my trace")
	Debug("debug", "tip", "my debug")
	Info("info", "tip", "my info")
	Warn("warn", "tip", "my warn")
}
