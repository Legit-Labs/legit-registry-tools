package legit_registry_tools

import (
	"fmt"
	"testing"
)

func TestImageRef(t *testing.T) {
	assert := func(ref, name, tag, digest, label string) {
		_name, _tag, _digest, err := GetImageRef(ref)
		if err != nil {
			t.Fatalf("failed to get image ref: %v", err)
		}
		if name != _name || tag != _tag || digest != _digest {
			t.Fatalf("failed to parse image with %v [%v,%v,%v]", label, _name, _tag, _digest)
		}
	}

	assert("image:tag@digest", "image", "tag", "digest", "both tag and digest")
	assert("image@digest", "image", "", "digest", "just digest")

	reservedImageName := "gallegit/helloc"
	reservedImageDigest := "sha256:2e5c54fa481719e4e7b9dd7bcaf62d8c805a84fb45775bed02a43068f3153351"

	assert(fmt.Sprintf("%v:latest", reservedImageName), reservedImageName, "latest", reservedImageDigest, "just tag")
	assert(reservedImageName, reservedImageName, "latest", reservedImageDigest, "none")
}
