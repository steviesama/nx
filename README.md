# nx

A Go library that contains several packages to be used in all the various nx projects.

## Packages

### [github.com/steviesama/nx](https://github.com/steviesama/nx)

The `nx` is currently only a namespace that holds the rest of the packages.

### [github.com/steviesama/nx/analyze](https://github.com/steviesama/nx/tree/master/analyze)

The `nx/analyze` package provides a way to analyze data. 

The following function for instance takes an empty interface slice that could point to anything...but the function expects a slice of some type. It iterates over the slice in the empty interface and determines if there are any duplicate elements of equal value.

```go
func analyze.HasSliceDuplicates(
  slice interface{},
  compare analyze.CompareFunc,
) bool
```

Additionally, it takes an anonymous function the caller can provide as inversion of control to allow the caller to determine the equality of slice. So far a `StringCompare` type is provided as a stock comparer.

```go
type analyze.CompareFunc = func(this, that interface{}) analyze.Equality
```

Refer to `types.go` in the `nx/analyze` package for more details.

### [github.com/steviesama/nx/conv](https://github.com/steviesama/nx/tree/master/conv)

The `nx/conv` package handles conversions that the Go library doesn't, or doesn't without jumping through hoops.

`conv.InterfaceSlice` is currently the only function in the package.

```go
func conv.InterfaceSlice(slice interface{}) []interface{}
```

### [github.com/steviesama/nx/crypto](https://github.com/steviesama/nx/tree/master/crypto)

The `nx/crypto` package handles the encryption & decryption processes and keeps them abbreviated. The Go `crypto` package is a bit winded to use. The goal with this library is to isolate as much of that as possible.

### [github.com/steviesama/nx/crypto/jwt](https://github.com/steviesama/nx/tree/master/crypto/jwt)

The `nx/crypto/jwt` package will provide access to JSON Web Tokens. They are cryptographical secure and have payloads that can be cryptographically signed and send to the front end to manage user access.

### [github.com/steviesama/nx/database](https://github.com/steviesama/nx/tree/master/database)

The `nx/database` package currently provides a way to create database connection pools accessible by a key allowing all packages access to the resources.

Currently it uses MySQL; however, it needs to be refactored, possibly creating `nx/database/mysql` package that can work along side any database dialect.

MongoDB needs to be worked in there somewhere. The architecture will take some time.

### [github.com/steviesama/nx/database/model](https://github.com/steviesama/nx/tree/master/database/model)

The `nx/database/model` package is intended to be an extended version of what is in place in other nx projects now where once the `struct` data shape is created, and the `Init()` function is created...an entry can be made for the newly created data type in `_generate_datafuncs.go` which will allows it's various data access layer functions to be generated from a template leveraging the Go package `text/template` and reflection in order to create the functions based on the data shape of the given type.

`nx/database/model` however, will allow you to register new model definitions...and perform a series of `model.Builder` function calls in order to describe and annotate these new data types. And they get generated in a build step before the primary compilation.

This is only designed for MySQL so far. Would like to include MongoDB minimally afterward.

### [github.com/steviesama/nx/ioutil](https://github.com/steviesama/nx/tree/master/ioutil)

The `nx/ioutil` package is a high level version of the Go `io/ioutil` package.

Much like other packages, it provides ready access to file de/serialization that the Go `io/ioutil` can't achieve without hurdles and including other packages.

### [github.com/steviesama/nx/iter](https://github.com/steviesama/nx/tree/master/iter)

The `nx/iter` package is intended to be a package that provides a way for objects to be iterable.

Currently it has minimal code...and requires some refactoring and design consideration.

### [github.com/steviesama/nx/jsonutil](https://github.com/steviesama/nx/tree/master/jsonutil)

The `nx/jsonutil` package provides high level access to the Go `encoding/json` package's functionality without the hurdles.

It marshal and unmarshal json data as bytes or strings...with or without indentation.

`nx/jsonutil` is used in conjunction with other packages with complementary functionality to achieve their task.

It makes use of the Go `reflect` package if you are looking for some examples of that.

### [github.com/steviesama/nx/rand](https://github.com/steviesama/nx/tree/master/rand)

The `nx/rand` package provides various functions that have random elements. So far there are 3 functions.

```go
func rand.String(n int) string
func rand.Bytes(n int) []byte
func rand.Guid(removeHyphens bool) string
```

### [github.com/steviesama/nx/service](https://github.com/steviesama/nx/tree/master/service)

The `nx/service` package thus far is a namespace for `net` related packages. It refers to web services or micro-services. They aren't all defined yet but will include the following:

  - [github.com/steviesama/nx/service/api](https://github.com/steviesama/nx/tree/master/service/api)
    - This will contain an interface and model definitions for api interaction, including rate limited.
  - [github.com/steviesama/nx/service/dropbox](https://github.com/steviesama/nx/tree/master/service/dropbox)
    - A design pattern organized, consolidated version of the `/dropbox` endpoint currently in production in `nxEquip`.
  - [github.com/steviesama/nx/service/emailer](https://github.com/steviesama/nx/tree/master/service/emailer)
    - Again, refactored version of the `/emailer` endpoint in production in `nxEquip`. And `EmailAccount` model will be included with this. It had rate limiting...but essentially iun the form of throttling the emails sent per minute depending on the email provided configured with the Emailer.
  - [github.com/steviesama/nx/service/imgur](https://github.com/steviesama/nx/tree/master/service/imgur)
    - Hopefully a robust Imgur service. There are a couple of possibilities out there. But if it comes down to it...I'll just create one from scratch.

## Wrap up

The library is growing as a moderate pace, but it's not going to be overdeveloped. New functionality will only be aded as necessary. Though there is still quite a lot of functionality left to translate from current nx projects.
