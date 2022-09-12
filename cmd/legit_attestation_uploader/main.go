package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/legit-labs/legit-registry-tools/pkg/legit_registry_tools"
)

var (
	image  string
	name   string
	path   string
	digest string
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "upload an attestation document to the designated location on the registry")
	}
	flag.StringVar(&image, "image", "", "The image (registry/image) to which the attestation refers")
	flag.StringVar(&name, "name", "", "The name of the attestation file (e.g. provenance, legit-score)")
	flag.StringVar(&path, "path", "", "The path to the attestation document")
	flag.StringVar(&digest, "digest", "", "The digest of the artifact to which the attestation refers")

	flag.Parse()

	if image == "" {
		log.Panicf("please provide an image")
	} else if name == "" {
		log.Panicf("please provide a name")
	} else if path == "" {
		log.Panicf("please provide an attestation path")
	} else if digest == "" {
		log.Panicf("please provide a digest")
	}

	err := legit_registry_tools.UploadAttestation(image, name, path, digest)
	if err != nil {
		log.Panicf("failed to upload attestation: %v", err)
	}

	log.Printf("provenance verified successfully.")
}
