package models

import (
	"errors"
	"testing"
)

func TestSettings_IsNICEnabledForDHCP(t *testing.T) {
	t.Parallel()

	// nil / empty は後方互換で全部true
	{
		var s Settings
		if !s.IsNICEnabledForDHCP("Ethernet") {
			t.Fatalf("expected true when EnabledDHCPNICs is empty")
		}
	}

	{
		s := Settings{EnabledDHCPNICs: []string{}}
		if !s.IsNICEnabledForDHCP("Wi-Fi") {
			t.Fatalf("expected true when EnabledDHCPNICs is empty slice")
		}
	}

	{
		s := Settings{EnabledDHCPNICs: []string{"Ethernet", "Wi-Fi"}}
		if !s.IsNICEnabledForDHCP("Wi-Fi") {
			t.Fatalf("expected true for included nic")
		}
		if s.IsNICEnabledForDHCP("VPN") {
			t.Fatalf("expected false for not included nic")
		}
	}
}

func TestIsValidNICName(t *testing.T) {
	t.Parallel()

	ok := []string{
		"Ethernet",
		"Wi-Fi",
		"イーサネット",
		"LAN 1",
	}
	for _, name := range ok {
		name := name
		t.Run("ok/"+name, func(t *testing.T) {
			t.Parallel()
			if !IsValidNICName(name) {
				t.Fatalf("expected valid: %q", name)
			}
		})
	}

	ng := []string{
		"",                 // empty
		"bad&name",         // shell meta
		"bad|name",         // shell meta
		"bad;name",         // shell meta
		"bad\nname",        // control
		"bad\tname",        // control
		"bad\"name",        // quote
		"bad'name",         // quote
		"bad\\name",        // escape
		"bad`name",         // shell
	}
	for _, name := range ng {
		name := name
		t.Run("ng/"+name, func(t *testing.T) {
			t.Parallel()
			if IsValidNICName(name) {
				t.Fatalf("expected invalid: %q", name)
			}
		})
	}
}

func TestProfile_Validate(t *testing.T) {
	t.Parallel()

	base := func() *Profile {
		return &Profile{
			ID:         "id",
			Name:       "home",
			IPAddress:  "192.168.1.10",
			SubnetMask: "255.255.255.0",
			Gateway:    "192.168.1.1",
			DNSPrimary: "8.8.8.8",
			NICName:    "Ethernet",
		}
	}

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		p := base()
		if err := p.Validate(); err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
	})

	t.Run("missing name", func(t *testing.T) {
		t.Parallel()
		p := base()
		p.Name = ""
		err := p.Validate()
		if !errors.Is(err, ErrInvalidProfileName) {
			t.Fatalf("expected ErrInvalidProfileName, got %v", err)
		}
	})

	t.Run("invalid ip", func(t *testing.T) {
		t.Parallel()
		p := base()
		p.IPAddress = "999.999.999.999"
		err := p.Validate()
		if !errors.Is(err, ErrInvalidIPAddress) {
			t.Fatalf("expected ErrInvalidIPAddress, got %v", err)
		}
	})

	t.Run("invalid subnet mask", func(t *testing.T) {
		t.Parallel()
		p := base()
		p.SubnetMask = "255.0.255.0"
		err := p.Validate()
		if !errors.Is(err, ErrInvalidSubnetMask) {
			t.Fatalf("expected ErrInvalidSubnetMask, got %v", err)
		}
	})

	t.Run("invalid gateway", func(t *testing.T) {
		t.Parallel()
		p := base()
		p.Gateway = "not-an-ip"
		err := p.Validate()
		if err == nil {
			t.Fatalf("expected error")
		}
	})

	t.Run("invalid dns", func(t *testing.T) {
		t.Parallel()
		p := base()
		p.DNSPrimary = "300.1.1.1"
		err := p.Validate()
		if err == nil {
			t.Fatalf("expected error")
		}
	})

	t.Run("invalid nic name chars", func(t *testing.T) {
		t.Parallel()
		p := base()
		p.NICName = "Ethernet & rm -rf"
		err := p.Validate()
		if !errors.Is(err, ErrInvalidNICName) {
			t.Fatalf("expected ErrInvalidNICName, got %v", err)
		}
	})
}

