# godoc_dev_sync

## Godoc
It’s a tool to view your package’s documentation 

“It is a great exercise to use godoc in the early phases of your package’s API design to see how your concepts will be rendered on doc. Sometimes, the visualization also has an impact on the design. Godoc is the way your users will consume a package, so it is ok to tweak things to make them more accessible. Run godoc -http=<hostport> to start a godoc server locally.” 
  
Any Go packages installed inside $GOROOT/src/pkg and any GOPATH work spaces will already be accessible via godoc's command-line and HTTP interfaces. 
  
By default it runs on http://localhost:6060

Here is documentation on the different flags you can use for godoc: https://godoc.org/golang.org/x/tools/cmd/godoc 

## Documentation Formatting
The convention is simple: to document a type, variable, constant, function, or even a package, write a regular comment directly preceding its declaration, with no intervening blank line. Godoc will then present that comment as text alongside the item it documents. For example, this is the documentation for the fmt package's Fprint function:

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
```
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
```

## Code Blocks Insertion

## Writing Examples

## Bug Reporting

## Godoc Parse

## doc.go

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
The rise of go modules will cause some of us to develope outside the GOPATH. This is an issue as godoc only shows documentation inside the GOPATH. This an ongoing issue that I will continue to look into.

You can view the discussion of this issue on Github: https://github.com/golang/go/issues/26827

## HTML Generation
```
godoc -url=/pkg/your_package > your_page.html
```
