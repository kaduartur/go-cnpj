[![test](https://github.com/kaduartur/go-cnpj/actions/workflows/go.yml/badge.svg)](https://github.com/kaduartur/go-cnpj/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/kaduartur/go-cnpj/branch/main/graph/badge.svg?token=M0UZ2KNHV5)](https://codecov.io/gh/kaduartur/go-cnpj)
[![CodeFactor](https://www.codefactor.io/repository/github/kaduartur/go-cnpj/badge)](https://www.codefactor.io/repository/github/kaduartur/go-cnpj)

# go-cnpj

go-cnpj is a lightweight and simple lib for validating the Brazilian National Register of Legal Entities (CNPJ)

### :floppy_disk: How to install?

Run the `go get` command in terminal:
```sh
go get github.com/kaduartur/go-cnpj
```

### :computer: How to use this?

To use the lib you have to import the package into the file you want:
```go
package main

import "github.com/kaduartur/go-cnpj/cnpj"

func main() {
	if err := cnpj.IsValid("79.276.501/0001-55"); err != nil {
		panic(err)
	}

	if err := cnpj.IsValid("79276501000155"); err != nil {
		panic(err)
	}
}
```
