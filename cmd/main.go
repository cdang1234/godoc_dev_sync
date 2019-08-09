package main

import (
	"github.com/weconnect/playground/blob"
)

func main() {
	b := &blob.Blob{Content: "random content"}
	b.LogBlob()
	b.BlobType()
}
