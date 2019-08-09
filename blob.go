/* BUG(r): Top-level comments that begin with the word "BUG(who)” are recognized as known bugs,
and included in the "Bugs” section of the package documentation. The "who” part should be
the user name of someone who could provide more information.
*/

package blob

import (
	"fmt"
	"reflect"
)

// Anonymous is an interface that defines functions for structs with anonymous types
type Anonymous interface {
	LogBlob()
	BlobType() reflect.Type
}

// Blob is a class that contains random information
type Blob struct {
	Content interface{}
}

// LogBlob logs the contents of blob
// You can embed blocks of code in your godoc, such as this:
//  b.LogBlob()
// To do that, simply add an extra indent to your comment's text.
// This works for doc.go as well
func (b *Blob) LogBlob() {
	fmt.Printf("%v", b.Content)
}
