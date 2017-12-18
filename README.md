# GoLang

## Installation

- Head to www.golang.org
- You will figure out the rest

## Compiling & Running

> go run filename.go

## Code formatting

> gofmt filename.go

_This will output the output to console only._

## Documentations on the go

> godoc pkgName
> godoc pkgName funcName

## Code Commenting [C-style]

    // For single line comments use double front slashes

    /*
    Multi line comments can be added like this.
    */`

## Caveats 🔥

- Use Double Quotes only, avoid single quotes
- Semicolons are optional
- You will see `imported and not used: "math"` warnings if you are not using any imported packages

## Importing External Packages

> go get golang.org/x/tools/cmd/goimports