package blob

import "reflect"

// BlobType logs the contents of blob
// Notice that BlobType is in a separate file yet it gets group with the Blob struct's documentation!
func (b *Blob) BlobType() reflect.Type {
	return reflect.TypeOf(b.Content)
}
