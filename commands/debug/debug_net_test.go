package debug

import "testing"

func TestDebugNet(t *testing.T) {
	debugNet("ParseIp", []string{"localhost"})
}
