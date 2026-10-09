package main

import (
	"io"
	"testing"

	"github.com/maisem/tailmix/tunmux"
)

func TestTSNetConfigTun(t *testing.T) {
	t.Run("socks uses nil interface for userspace networking", func(t *testing.T) {
		cfg, err := tsnetConfig(runtimeProfile{}, io.Discard, false, false, "")
		if err != nil {
			t.Fatal(err)
		}
		if cfg.Tun != nil {
			t.Fatalf("SOCKS TUN = %T, want a nil interface so tsnet selects userspace networking", cfg.Tun)
		}
	})

	t.Run("tun retains the supplied device", func(t *testing.T) {
		device := tunmux.NewChanTUN("test-tun")
		t.Cleanup(func() { _ = device.Close() })
		cfg, err := tsnetConfig(runtimeProfile{Tun: device}, io.Discard, false, false, "")
		if err != nil {
			t.Fatal(err)
		}
		if cfg.Tun != device {
			t.Fatal("TUN mode did not retain its supplied device")
		}
	})
}
