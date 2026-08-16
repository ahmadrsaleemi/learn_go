# Learn Go

Becoming production-ready in Go, building on 9 years of development experience.

## Goal

Learn Go well enough to work confidently in a production codebase — write idiomatic code, review PRs, debug issues, and ship features when needed.

## Learned so far

- **Hello World** — `package main`, `fmt`, variables, string concatenation ([`basics/hello_world.txt`](basics/hello_world.txt))
- **Imports** — multi-import blocks, `net/http`, basic error checks ([`basics/import.txt`](basics/import.txt))
- **Types & functions** — `string`/`int`, `:=`, custom functions ([`basics/data_types.txt`](basics/data_types.txt))
- **Structs** — defining types, fields, creating values, passing structs to functions ([`basics/data_types.txt`](basics/data_types.txt))
- **Control flow** — `if`/`else if`, `switch`, classic `for`, infinite `for` + `break`, nested loops, `continue` ([`basics/loops.txt`](basics/loops.txt))
- **Slices** — arrays vs slices, slicing, `copy`, `range`, 2D slices ([`basics/loops.txt`](basics/loops.txt))
- **Maps** — `make`, insert/delete, comma-ok check, `maps.Equal`, ranging over maps ([`basics/loops.txt`](basics/loops.txt))
- **Other concepts** — variadic functions (`...`), multiple return values, `strconv.Atoi`, `defer`, `panic` ([`basics/otherconcepts.txt`](basics/otherconcepts.txt))
- **Constants & arithmetic** — `const` groups and basic arithmetic ([`basics/const_and_arithemetic_ops.txt`](basics/const_and_arithemetic_ops.txt))
- **Closures** — functions that return functions and keep state across calls ([`intermediate/closure.txt`](intermediate/closure.txt))
- **Recursion** — factorial, fibonacci, sum of digits ([`intermediate/recursion.txt`](intermediate/recursion.txt))
- **Pointers** — declaring pointers, `&` address-of, `*` dereference ([`intermediate/pointers.txt`](intermediate/pointers.txt))
- **Embedded structs** — struct embedding and method overriding ([`intermediate/embedded_structs.txt`](intermediate/embedded_structs.txt))
- **Interfaces** — defining interfaces and implementing them on multiple types ([`intermediate/interfaces.txt`](intermediate/interfaces.txt))
- **Generics** — type parameters, generic functions, and a generic stack ([`intermediate/generics.txt`](intermediate/generics.txt))
- **String formatting** — `Printf` width/padding and raw vs interpreted strings ([`intermediate/string_formatting.txt`](intermediate/string_formatting.txt))
- **String manipulation** — escape sequences, indexing, ranging runes, `utf8` ([`intermediate/string_manipulation.txt`](intermediate/string_manipulation.txt))
- **String functions** — `strings`/`strconv` helpers like Split, Join, Itoa ([`intermediate/string_functions.txt`](intermediate/string_functions.txt))
- **Number parsing** — parsing strings to ints with `strconv.Atoi` ([`intermediate/number_parsing.txt`](intermediate/number_parsing.txt))
- **Time** — `time.Now`, Unix timestamps, parsing and formatting ([`intermediate/time.txt`](intermediate/time.txt))
- **Random numbers** — `math/rand` with seeded sources ([`intermediate/rand_nums.txt`](intermediate/rand_nums.txt))
- **Text templates** — `html/template` parse and execute ([`intermediate/text_templates.txt`](intermediate/text_templates.txt))
- **URL parsing** — `net/url` parse, query params, and building URLs ([`intermediate/url_parsing.txt`](intermediate/url_parsing.txt))
- **Base64 encoding** — encode and decode with `encoding/base64` ([`intermediate/encode.txt`](intermediate/encode.txt))
- **Hashing** — SHA-256/512, salts, and password hashing ([`intermediate/hashing.txt`](intermediate/hashing.txt))
- **JSON** — struct tags and `json.Marshal` ([`intermediate/json.txt`](intermediate/json.txt))
- **File operations** — create, write, read, and path helpers ([`intermediate/file_ops.txt`](intermediate/file_ops.txt))
- **CLI args** — `os.Args` and `flag` parsing ([`intermediate/cli_args.txt`](intermediate/cli_args.txt))
- **Goroutines** — launching concurrent work with `go` ([`advance/go_routines.go`](advance/go_routines.go))

## Course plan

### Basics & intermediate

- Comprehensive examples of basic Go concepts
- Intermediate concepts with detailed explanation and practice
- Pointers in Go
- Structs, maps, and slices — deep practice
- Git and GitHub
- Quizzes, slides, and downloadable PDF material

### Advanced Go

- Highly extensive advanced concepts section
- Go runtime — how it works and why it matters
- Reflect package — comprehensive use in gRPC & REST projects
- Algorithms applied to real-world cases
- Reading Go source code to solve problems
- Interview prep — 350+ questions and answers

### Concurrency

- How concurrency works in Go
- Goroutines — full coverage with many examples
- Channels — importance and real use cases

### APIs & production practice

- Plan before building an API
- Professional, industry-standard API design and folder structure
- REST API in Go
- gRPC API in Go
- Protocol Buffers and gRPC — extensive practice
- Middleware from scratch
- HTTP/2 and HTTPS APIs
- TLS/SSL in APIs
- SQL and NoSQL real use cases (MariaDB / MySQL, MongoDB)
- API benchmarking with tools like `wrk`, `h2load`, `ghz`

## Resources

- [Go Tour](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
