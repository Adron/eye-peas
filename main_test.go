package main

import "testing"

func TestIpsBetween(t *testing.T) {
	t.Run("should handle basic cases", func(t *testing.T) {
		assertEqual(t, IpsBetween("10.0.0.0", "10.0.0.50"), 50)
		assertEqual(t, IpsBetween("20.0.0.10", "20.0.1.0"), 246)
	})

	t.Run("should handle large ranges", func(t *testing.T) {
		assertEqual(t, IpsBetween("10.0.0.0", "10.0.1.0"), 256)
		assertEqual(t, IpsBetween("10.0.0.0", "10.1.0.0"), 65536)
	})

	t.Run("should handle different octets", func(t *testing.T) {
		assertEqual(t, IpsBetween("10.0.0.0", "11.0.0.0"), 16777216)
		assertEqual(t, IpsBetween("10.0.0.0", "12.0.0.0"), 33554432)
	})

	t.Run("should handle same start and end addresses", func(t *testing.T) {
		assertEqual(t, IpsBetween("10.0.0.0", "10.0.0.0"), 0)
		assertEqual(t, IpsBetween("192.168.1.1", "192.168.1.1"), 0)
	})
}

func assertEqual(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
