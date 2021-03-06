// Copyright 2016 NetApp, Inc. All Rights Reserved.

package utils

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

func TestPow(t *testing.T) {
	log.Debug("Running TestPow...")

	if Pow(1024, 0) != 1 {
		t.Error("Expected 1024^0 == 1")
	}

	if Pow(1024, 1) != 1024 {
		t.Error("Expected 1024^1 == 1024")
	}

	if Pow(1024, 2) != 1048576 {
		t.Error("Expected 1024^2 == 1048576")
	}

	if Pow(1024, 3) != 1073741824 {
		t.Error("Expected 1024^3 == 1073741824")
	}
}

func TestConvertSizeToBytes(t *testing.T) {
	log.Debug("Running TestConvertSizeToBytes...")

	d := make(map[string]string)
	d["512"] = "512"
	d["1kB"] = "1024"
	d["4k"] = "4096"
	d["1GB"] = "1073741824"
	d["1gb"] = "1073741824"
	d["1g"] = "1073741824"
	d["2GB"] = "2147483648"

	for k, v := range d {
		s, err := ConvertSizeToBytes(k)
		if err != nil || s != v {
			t.Errorf("Expected convertSizeToBytes('%v') == '%v' but was %v", k, v, s)
		}
	}
}

func TestGetV(t *testing.T) {
	log.Debug("Running TestGetV...")

	d := make(map[string]string)
	d["key1"] = "value1"

	if val := GetV(d, "key1", "defaultValue"); val != "value1" {
		t.Errorf("Expected '%v' but was %v", "value1", val)
	}

	if val := GetV(d, "key2", "defaultValue"); val != "defaultValue" {
		t.Errorf("Expected '%v' but was %v", "defaultValue", val)
	}
}
