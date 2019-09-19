package util_test

import (
	"sooty-tern/pkg/util"
	"testing"
)

func TestMustUUID(t *testing.T) {
	uuid := util.MustUUID()
	t.Logf("uuid: " + uuid)
	t.Logf("uuid length: %d", len(uuid))
}
