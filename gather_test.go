package main

import (
	"testing"

	"github.com/go-test/deep"
)

func TestGoodPath(t *testing.T) {
	keys := getAllKeys("./src/lang")
	exp := map[string]string{"dashboard_actions_referrals_dialog_trigger_cta_description": "en/lang3.json", "key_1": "lang.json", "key_2": "lang.json", "key_3": "lang.json", "key_4": "lang2.json", "key_5": "lang2.json"}

	if deep.Equal(keys, exp) != nil {
		t.Errorf("getAllKeys('./src/lang') = %v; want %v", keys, exp)
	}
}

func TestBadPath(t *testing.T) {
	keys := getAllKeys("./foo")
	exp := map[string]string{}

	if deep.Equal(keys, exp) != nil {
		t.Errorf("getAllKeys('./foo') = %v; want %v", keys, exp)
	}
}
