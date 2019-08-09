package blob_test

import (
	"fmt"

	"github.com/cdang1234/godoc_dev_sync/blob"
	pegomock "github.com/petergtz/pegomock"
)

//go:generate pegomock generate --package blob github.com/cdang1234/godoc_dev_sync Anonymous

// This function is named Example(). godoc doesn't know what to associate it with
// so it is placed at the package level of the documentation
func Example() {
	fmt.Println("Hello")
}

func ExampleAnonymous() {
	anonymous = NewMockAnonymous()
	pegomock.When(anonymous.LogBlob()).ThenReturn("a") // Should return contents of Blob

	fmt.Println(anonymous.LogBlob())
	// Output: a
}

// This function is named ExampleBlob(), it is associated with
// Blob type.
func ExampleBlob() {
	b := &blob.Blob{Content: "a"}
	fmt.Print(b.Content == "a")
	// Output: true
}

// This function is named ExampleBlob_LogBlob_a and associated with the
// Blob type's LogBlob method. It has a label of a.
func ExampleBlob_LogBlob_a() {
	b := &blob.Blob{Content: "a"}
	b.LogBlob()
	// Output: a
}

// This function is named ExampleBlob_LogBlob_b and associated with the
// Blob type's LogBlob method. It has a label of b.
func ExampleBlob_LogBlob_b() {
	b := &blob.Blob{Content: "b"}
	b.LogBlob()
	// Output: b
}

// This function is named ExampleBlob_BlobType and associated with the
// Blob type's BlobType method. It returns the type of Blob's contents.
func ExampleBlob_BlobType() {
	b := &blob.Blob{Content: "b"}
	fmt.Println(b.BlobType())
	// Output: string
}
