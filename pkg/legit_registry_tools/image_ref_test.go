package legit_registry_tools

import (
	"fmt"
	"strings"
	"testing"
)

func TestImageRef(t *testing.T) {
	assert := func(ref, name, tag, digest, label string) {
		imageRef, err := NewImageRef(ref)
		if err != nil {
			t.Fatalf("failed to get image ref: %v", err)
		}
		expected := ImageRef{Name: name, Tag: tag, Digest: digest}
		if *imageRef != expected {
			t.Fatalf("failed to parse image with %v [%v]", label, imageRef)
		}

		if HasDigest(ref) {
			if ref != imageRef.Ref() {
				t.Fatalf("exact ref mismatch %v [%v != %v]", label, ref, imageRef.Ref())
			}
		} else {
			if !strings.HasPrefix(imageRef.Ref(), ref) {
				t.Fatalf("partial ref mismatch %v [%v doesn't start with %v]", label, imageRef.Ref(), ref)
			}
		}

	}

	assert("image:tag@digest", "image", "tag", "digest", "both tag and digest")
	assert("image@digest", "image", "", "digest", "just digest")

	reservedImageName := "gallegit/helloc"
	reservedImageDigest := "sha256:2e5c54fa481719e4e7b9dd7bcaf62d8c805a84fb45775bed02a43068f3153351"

	assert(fmt.Sprintf("%v:latest", reservedImageName), reservedImageName, "latest", reservedImageDigest, "just tag")
	assert(reservedImageName, reservedImageName, "latest", reservedImageDigest, "none")
}
