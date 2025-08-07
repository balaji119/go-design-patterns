package singleton

import "testing"

func TestSingleton(t *testing.T) {
    s1 := GetInstance()
    s2 := GetInstance()

    if s1 != s2 {
        t.Error("Expected same instance, got different")
    }
}
