# godoc_dev_sync

This README covers practices to maximize the usage of the godoc tool. Download this repository and run ``godoc`` to view the examples.

## Godoc
It’s a tool to view your package’s documentation 

“It is a great exercise to use godoc in the early phases of your package’s API design to see how your concepts will be rendered on doc. Sometimes, the visualization also has an impact on the design. Godoc is the way your users will consume a package, so it is ok to tweak things to make them more accessible. Run godoc -http=<hostport> to start a godoc server locally.” -- https://rakyll.org/style-packages/
  
Any Go packages installed inside $GOROOT/src/pkg and any GOPATH work spaces will already be accessible via godoc's command-line and HTTP interfaces. 
  
By default it runs on http://localhost:6060

Here is documentation on the different flags you can use for godoc: https://godoc.org/golang.org/x/tools/cmd/godoc 

Godoc compiles your code as it builds the HTML. If your code godoc's web server will still work but it doesn't render the HTML for the package whose code is failing to compile.

Here is documentation on the different flags you can use for godoc: https://godoc.org/golang.org/x/tools/cmd/godoc

## Documentation Formatting
The convention is simple: to document a type, variable, constant, function, or even a package, write a regular comment directly preceding its declaration, with no intervening blank line. Godoc will then present that comment as text alongside the item it documents. For example, this is the documentation for the fmt package's Fprint function:

```
// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
```

## Code Blocks Insertion
You can embed blocks of code in your godoc with the addition of an extra indent to your comment's text:

https://github.com/cdang1234/godoc_dev_sync/blob/master/blob/blob.go#L18

## Writing Examples
You can insert examples by creating a file with the suffix ``_test.go`` within the same package as the code you would like to write examples for. I suggest writing a file called "example_test.go" whenever you want to write examples rather than appending examples to the package's testing suites.

The naming convention to declare examples for the package, a function F, a type T and method M on type T are:

```
func Example() { ... }
func ExampleF() { ... }
func ExampleT() { ... }
func ExampleT_M() { ... }
```
Multiple example functions for a package/type/function/method may be provided by appending a distinct suffix to the name. The suffix must start with a lower-case letter.

```
func Example_suffix() { ... }
func ExampleF_suffix() { ... }
func ExampleT_suffix() { ... }
```

Please note that go's testing package also runs and verifies example code. This means if that the concluding line of your example has a line that begins with `` // Output: `` whatever is that defined by this output line will have its value compared with the standard output of the function. 

Let's take an example from this repository

```
func ExampleBlob_LogBlob_b() {
	b := &blob.Blob{Content: "b"}
	b.LogBlob()
	// Output: b
}
```

The above code has a concluding line `` // Output: b`` which is interpreted by Go's testing package as asserting that the standard output of ExampleBlob_LogBlob_b() will be `b`. If this function doesn't write to standard out or writes the wrong value to standard out your tests will fail!

More examples within this repository:
https://github.com/cdang1234/godoc_dev_sync/blob/master/blob/example_test.go

Additional documentation on examples:
https://golang.org/pkg/testing/

## Writing Examples for Interfaces

Let's say you want examples to show how you want your interfaces to behave. Since we can't instantiate interfaces we must instead generate mocks of the interfaces via pegomock. 

Once this is done you leverage pegomock to demontrate what the expected behavior is for types that implement the interface.



## Bug Reporting
Top-level comments that begin with the word "BUG(who)” are recognized as known bugs, and included in the "Bugs” section of the package documentation. 

Example: https://github.com/cdang1234/godoc_dev_sync/blob/master/blob/blob.go#L1

## Documentation Binding

Notice that documentation for a function will be grouped with its receiver. 

The Blob struct is defined in blob.go yet its function BlobType() is defined in random_file.go. If you run godoc on this repository you will find that BlobType()'s documentation is grouped with the Blob type's documentation.

## doc.go

If you want to add documentation to the package-level overview section then write a doc.go file. Here is an example of what the doc.go file generates for https://github.com/go-mgo/mgo:

![doc.go example](https://golang-for-python-programmers.readthedocs.io/en/latest/_images/package.png)

When writing package comments of any size, keep in mind that their first sentence will appear in godoc's package list:
https://github.com/cdang1234/godoc_dev_sync/blob/master/doc.go#L2

doc.go's are defined per package which means you can have multiple doc.go files in your project.

You can insert code into your doc.go file as well. Simply indent the text you would like to display as code.

## Titles for doc.go
You can make titles in your godoc. A title is a line that is separated from its following line by an empty line, begins with a capital letter and doesn't end with punctuation.
For example, the code:
```
// Sentence 1
//
// Sentence 2
```

will yield:
Sentence 1
Sentence 2

While this code:
```
// Sentence 1.
//
// Sentence 2.
```

will yield:
Sentence 1.
Sentence 2.

## Paragraphs for doc.go
To start a new paragraph, add an empty line in the comment between the 2 paragraphs.
For example:
```
// Paragraph 1.
// Still paragraph 1.
//
// Paragraph 2.
// Still Paragraph 2.
```

will yield:
Paragraph 1. Still paragraph 1.
Paragraph 2. Still Paragraph 2.


## Moving Outside of GOPATH
The rise of go modules will cause some of us to develope outside the GOPATH. This is an issue as godoc only shows documentation inside the GOPATH.

You can view the discussion of this ongoing issue on Github: https://github.com/golang/go/issues/26827

Some workarounds include writing a script to copy your project onto the GOPATH or copying your project into a the GOPATH of a Docker container and running godoc from inside the container.

### Articles Use for this Dev Sync
https://blog.golang.org/godoc-documenting-go-code

https://rakyll.org/style-packages/

https://godoc.org/github.com/fluhus/godoc-tricks

### Supplement Article

https://noti.st/kenigbolomeyastephen/gFreMs/continuous-documentation-the-best-time-is-now#sUftrMl

The article above was given during a lightning talk at Gophercon 2019 and has a lot of helpful notes on how to improve documentation for your repository

