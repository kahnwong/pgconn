package internal

import "testing"

func TestSetProxyPort(t *testing.T) {
	p := Pgconn{}
	port := p.SetProxyPort()

	// Test that port is within valid range
	if port < 5432 || port > 8000 {
		t.Errorf("SetProxyPort() returned %d, expected value between 5432 and 8000", port)
	}
}

func TestSetProxyPortMultipleCalls(t *testing.T) {
	p := Pgconn{}

	// Call multiple times to ensure it always returns valid ports
	for i := 0; i < 10; i++ {
		port := p.SetProxyPort()
		if port < 5432 || port > 8000 {
			t.Errorf("SetProxyPort() call %d returned %d, expected value between 5432 and 8000", i, port)
		}
	}
}
