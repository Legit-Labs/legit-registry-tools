package legit_registry_tools

import (
	"fmt"
	"strings"

	"github.com/google/go-containerregistry/pkg/crane"
)

type ImageRef struct {
	Name   string
	Tag    string
	Digest string
}

func NewImageRef(ref string) (*ImageRef, error) {
	ref, digest := getImageWithDigest(ref)
	hasDigest := digest != ""
	name, tag := getImageWithTag(ref, hasDigest)

	if !hasDigest {
		var err error
		digest, err = crane.Digest(ref)
		if err != nil {
			return nil, err
		}
	}

	return &ImageRef{
		Name:   name,
		Tag:    tag,
		Digest: digest,
	}, nil
}

func (i ImageRef) Ref() string {
	ref := i.Name
	if i.Tag != "" {
		ref = fmt.Sprintf("%v:%v", ref, i.Tag)
	}
	if i.Digest != "" {
		ref = fmt.Sprintf("%v@%v", ref, i.Digest)
	}

	return ref
}

const (
	digestSeparator = "@"
	tagSeparator    = ":"
)

func HasDigest(ref string) bool {
	return strings.Contains(ref, digestSeparator)
}
func HasTag(ref string) bool {
	return strings.Contains(ref, tagSeparator)
}
func SplitByDigest(ref string) (string, string) {
	parts := strings.Split(ref, digestSeparator)
	return parts[0], parts[1]
}
func SplitByTag(ref string) (string, string) {
	parts := strings.Split(ref, tagSeparator)
	return parts[0], parts[1]
}

func getImageWithDigest(ref string) (subRef, digest string) {
	if HasDigest(ref) {
		subRef, digest = SplitByDigest(ref)
	} else {
		subRef = ref
	}

	return
}

func getImageWithTag(ref string, hasDigest bool) (subRef, tag string) {
	if HasTag(ref) {
		subRef, tag = SplitByTag(ref)
	} else {
		subRef = ref
		if !hasDigest {
			tag = "latest" // default when there is no tag & digest
		}
	}

	return
}
