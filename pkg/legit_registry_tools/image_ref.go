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

func (i ImageRef) Ref() string {
	if i.Digest != "" {
		return fmt.Sprintf("%v@%v", i.Name, i.Digest)
	} else {
		return fmt.Sprintf("%v:%v", i.Name, i.Tag)
	}
}

func getImageWithDigest(ref string) (subRef, digest string) {
	if strings.Contains(ref, "@") {
		parts := strings.Split(ref, "@")
		subRef, digest = parts[0], parts[1]
	} else {
		subRef = ref
	}

	return
}

func getImageWithTag(ref string, hasDigest bool) (subRef, tag string) {
	if strings.Contains(ref, ":") {
		parts := strings.Split(ref, ":")
		subRef, tag = parts[0], parts[1]
	} else {
		subRef = ref
		if !hasDigest {
			tag = "latest" // default when there is no tag & digest
		}
	}

	return
}

func GetImageRef(ref string) (name, tag, digest string) {
	ref, digest = getImageWithDigest(ref)
	hasDigest := digest != ""
	name, tag = getImageWithTag(ref, hasDigest)

	if !hasDigest {
		var err error
		digest, err = crane.Digest(ref)
		if err != nil {
			digest = ""
		}
	}

	return
}
