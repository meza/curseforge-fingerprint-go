package curseforgeFingerprint

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestHashing(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	fixturePath := filepath.Join(currentDir, "__fixtures__", "test1.md")
	fixturePath2 := filepath.Join(currentDir, "__fixtures__", "test2.md")
	fixturePath3 := filepath.Join(currentDir, "__fixtures__", "fabric-carpet-24w33a-1.4.148+v240818.jar")
	fixturePath4 := filepath.Join(currentDir, "__fixtures__", "ServerRedstoneBlock-fabric-1.19.3-1.0.0.jar")

	actualHash1 := GetFingerprintFor(fixturePath)
	actualHash2 := GetFingerprintFor(fixturePath2)
	actualHash3 := GetFingerprintFor(fixturePath3)
	actualHash4 := GetFingerprintFor(fixturePath4)

	assert.Equal(t, uint32(3608199863), actualHash1, "Hashes should match")
	assert.Equal(t, uint32(3493718775), actualHash2, "Hashes should match")
	assert.Equal(t, uint32(419608402), actualHash3, "Hashes should match")
	assert.Equal(t, uint32(2490932876), actualHash4, "Hashes should match")
}
