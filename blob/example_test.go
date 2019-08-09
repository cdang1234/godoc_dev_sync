package blob_test

import (
	"fmt"
	"log"

	"github.com/weconnect/playground/blob"
)

// This function is named Example(). godoc doesn't know what to associate it with
// so it is placed at the package level of the documentation
func Example() {
	fmt.Println("Hello")
}

// This function is named ExampleBlob(), it is associated with
// Blob type.
func ExampleBlob() {
	b := &blob.Blob{Content: "a"}
	log.Print(b.Content == "a")
	// Output: true
}

// This function is named ExampleBlob_LogBlob_a and associated with the
// Blob type's LogBlob method. It has a label of a.
func ExampleBlob_LogBlob_a() {
	b := &blob.Blob{Content: "a"}
	b.LogBlob()
	// Output: "a"
}

// This function is named ExampleBlob_LogBlob_b and associated with the
// Blob type's LogBlob method. It has a label of b.
func ExampleBlob_LogBlob_b() {
	b := &blob.Blob{Content: "b"}
	b.LogBlob()
	// Output: "b"
}
