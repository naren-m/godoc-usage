# Using godoc

## Generating documentation using godoc

Run below command to see documentation in commandline

```shell
godoc github.com/naren-m/godoc-usage
```

Output:

```shell
PACKAGE DOCUMENTATION

package calculator
    import "github.com/naren-m/godoc-usage"

    Package calculator is tutorial that shows how to use godoc for creating
    documentation to you go code

FUNCTIONS

func Add(a int, b int) int
    Add calculates and returns sum of two integers passed as input
    parameters
```

## See documentation in web browser

```shell
godoc -http=:6060
```

In browser go to "Packages" and navigate from there to your package

## Generating documentaion in HTML format using godoc

Generating from the url flag

```shell
godoc -url http://localhost:6060/pkg/github.com/naren-m/godoc-usage/ > page.html
```

Generating without using html flag

```shell
godoc -html github.com/naren-m/godoc-usage > page.html
```

The HTML we generated above will look better after pointing the css to

- [style.css](https://golang.org/lib/godoc/style.css)
- [jquery.treeview.css](https://golang.org/lib/godoc/jquery.treeview.css)

and javascript to

- [jquery.min.js](https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js)
- [godocs.js](https://golang.org/lib/godoc/godocs.js)
- [playground.js](https://golang.org/lib/godoc/playground.js)
- [jquery.treeview.js](https://golang.org/lib/godoc/jquery.treeview.js)
- [jquery.treeview.edit.js](https://golang.org/lib/godoc/jquery.treeview.edit.js)


## References

- [GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc) doccumentation
- [GoDoc](https://blog.golang.org/godoc-documenting-go-code) for documenting your go code
- [Golang Documenting Package Examples](http://blog.el-chavez.me/2013/08/29/golang-documenting-package-examples/)
- [How to generate HTML documents using godoc](http://stackoverflow.com/questions/13530120/how-can-i-generate-html-documents-using-godoc)
- [How do you server simple documentation for go programs using godoc as a webpage?](http://stackoverflow.com/questions/24514885/how-do-you-server-simple-documentation-for-go-programs-using-godoc-as-a-webpage)