<img src="assets/gologo.webp" style="width: 110px">

# jsonlite

A lightweight and simple JSON manipulation library for Go.

`jsonlite` provides a convenient way to load JSON files, decode their contents into Go structures, modify the decoded data, and write the changes back to the original file.

The library is designed to keep JSON file manipulation simple while still providing the flexibility of Go's standard `encoding/json` package.

---

## Features

* Load JSON files directly from disk
* Decode JSON into any Go value
* Work with dynamic JSON using `map[string]any`
* Modify decoded JSON data
* Write modified data back to the original JSON file
* Configurable indentation
* Configurable JSON prefix
* Configurable file permissions
* Built entirely on Go's standard library

---

## Installation

Install the package using `go get`:

```bash
go get github.com/Reollinos/jsonlite-go
```

Or, if the project is hosted under a repository such as GitHub:

```bash
go get github.com/Reollinos/jsonlite-go
```

Then import it in your Go program:

```go
import "github.com/Reollinos/jsonlite-go"
```

---

# Quick Start

Suppose you have a file called `example.json`:

```json
{
    "name": "Renan",
    "age": 16,
    "active": true
}
```

You can load and manipulate it with:

```go
package main

import (
    "fmt"
    "github.com/Reollinos/jsonlite-go"
)

func main() {
    jsonFile, err := jsonlite.Load("example.json")
    if err != nil {
        panic(err)
    }

    var data map[string]any

    err = jsonFile.Structure(&data)
    if err != nil {
        panic(err)
    }

    fmt.Println(data["name"])
}
```

Output:

```text
Renan
```

---

# Loading a JSON File

Use `jsonlite.Load()` to load a JSON file:

```go
file, err := jsonlite.Load("config.json")
if err != nil {
    panic(err)
}
```

`Load()` reads the file and keeps its contents internally.

The returned object can then be used to decode and manipulate the JSON.

---

# Decoding JSON

The `Structure()` method uses Go's `encoding/json` internally.

You can decode the JSON into any compatible Go value.

### Using a struct

```go
type User struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    Active bool   `json:"active"`
}

var user User

err := file.Structure(&user)
if err != nil {
    panic(err)
}

fmt.Println(user.Name)
```

This is useful when you already know the structure of your JSON.

---

## Using `map[string]any`

If the JSON structure is unknown or dynamic, use a map:

```go
var data map[string]any

err := file.Structure(&data)
if err != nil {
    panic(err)
}
```

For example:

```json
{
    "name": "Renan",
    "age": 16,
    "active": true
}
```

Can be accessed with:

```go
fmt.Println(data["name"])
fmt.Println(data["age"])
fmt.Println(data["active"])
```

Output:

```text
Renan
16
true
```

This means you don't need to define a struct beforehand.

---

# Working With Nested JSON

Nested objects are represented as another `map[string]any`.

For example:

```json
{
    "users": {
        "reollinos": {
            "id": 123,
            "active": true
        }
    }
}
```

You can access the nested values like this:

```go
users := data["users"].(map[string]any)

reollinos := users["reollinos"].(map[string]any)

fmt.Println(reollinos["id"])
fmt.Println(reollinos["active"])
```

---

# Adding Data

Because `map[string]any` is a normal Go map, you can add new values normally.

```go
data["version"] = 1
data["author"] = "Renan"
data["active"] = true
```

You can also add entire objects:

```go
data["config"] = map[string]any{
    "debug": true,
    "theme": "dark",
}
```

---

# Removing Data

Use Go's built-in `delete()` function:

```go
delete(data, "author")
```

For nested objects:

```go
users := data["users"].(map[string]any)

delete(users, "reollinos")
```

---

# Saving Changes

After modifying your data, use `ReloadJson()` to write the current structure back to the original JSON file.

```go
err := file.ReloadJson()
if err != nil {
    panic(err)
}
```

For example:

```go
var data map[string]any

err := file.Structure(&data)
if err != nil {
    panic(err)
}

data["version"] = 2
delete(data, "oldValue")

err = file.ReloadJson()
if err != nil {
    panic(err)
}
```

The JSON file will then contain the updated data.

---

# Formatting

`jsonlite` provides global preferences for JSON formatting.

The default configuration is:

```go
jsonlite.Preferences
```

with:

```go
Indent: "    "
Prefix: ""
Perm:   0644
```

## Indentation

You can change the indentation:

```go
jsonlite.Preferences.Indent = "\t"
```

Or use two spaces:

```go
jsonlite.Preferences.Indent = "  "
```

---

## Prefix

You can configure the prefix used by `json.MarshalIndent`:

```go
jsonlite.Preferences.Prefix = ""
```

For example:

```go
jsonlite.Preferences.Prefix = "    "
```

---

## File Permissions

The permissions used when rewriting the JSON file can be configured through:

```go
jsonlite.Preferences.Perm
```

Example:

```go
jsonlite.Preferences.Perm = 0644
```

---

# Complete Example

The following example loads a JSON file, reads its contents, adds a new value, removes another value, and saves the result.

### `example.json`

```json
{
    "name": "Renan",
    "version": 1,
    "debug": false
}
```

### `main.go`

```go
package main

import (
    "fmt"
    "github.com/Reollinos/jsonlite-go"
)

func main() {
    file, err := jsonlite.Load("example.json")
    if err != nil {
        panic(err)
    }

    var data map[string]any

    err = file.Structure(&data)
    if err != nil {
        panic(err)
    }

    fmt.Println("Name:", data["name"])

    // Add data
    data["author"] = "Renan"
    data["version"] = 2

    // Remove data
    delete(data, "debug")

    // Save changes
    err = file.ReloadJson()
    if err != nil {
        panic(err)
    }

    fmt.Println("JSON updated successfully.")
}
```

Result:

```json
{
    "author": "Renan",
    "name": "Renan",
    "version": 2
}
```

---

# API

## `Load`

```go
func Load(path string) (*header, error)
```

Loads a JSON file from disk.

### Parameters

| Parameter | Type     | Description           |
| --------- | -------- | --------------------- |
| `path`    | `string` | Path to the JSON file |

### Returns

* `*header` containing the loaded JSON
* `error` if the file cannot be read

Example:

```go
file, err := jsonlite.Load("config.json")
```

---

## `Structure`

```go
func (hd *header) Structure(structAddress any) error
```

Decodes the loaded JSON into the provided Go value.

Example:

```go
var config map[string]any

err := file.Structure(&config)
```

Or:

```go
var config Config

err := file.Structure(&config)
```

The destination must generally be passed as a pointer.

---

## `ReloadJson`

```go
func (hd *header) ReloadJson() error
```

Serializes the current structure and writes it back to the JSON file originally loaded with `Load()`.

Example:

```go
err := file.ReloadJson()
```

---

# How It Works

The basic workflow is:

```text
             JSON FILE
                 │
                 ▼
          jsonlite.Load()
                 │
                 ▼
          JSON file loaded
                 │
                 ▼
        file.Structure(&data)
                 │
                 ▼
        Go structure / map
                 │
          ┌──────┴──────┐
          │             │
       Modify        Read
          │             │
          └──────┬──────┘
                 ▼
        file.ReloadJson()
                 │
                 ▼
          JSON FILE UPDATED
```

`jsonlite` does not replace Go's JSON system. Instead, it provides a small abstraction around `os.ReadFile`, `os.WriteFile`, and `encoding/json`.

---

# Error Handling

All file and JSON operations return an `error`.

Always check the returned error:

```go
file, err := jsonlite.Load("config.json")
if err != nil {
    panic(err)
}
```

For production applications, you may want to handle the error instead of using `panic`:

```go
file, err := jsonlite.Load("config.json")
if err != nil {
    fmt.Println("Unable to load JSON:", err)
    return
}
```

---

# Requirements

* Go 1.18 or newer
* No external runtime dependencies

The library uses Go's standard packages:

```text
encoding/json
fmt
os
```

---

# Design Philosophy

`jsonlite` is intentionally small.

The goal is not to replace `encoding/json`, but to provide a convenient layer for applications that frequently need to:

1. Open a JSON file
2. Decode it
3. Manipulate its contents
4. Save it again

For applications that require advanced JSON manipulation, schema validation, streaming, or high-performance serialization, specialized JSON libraries may be more appropriate.

---

# Project Status

`jsonlite` is currently under development.

The API may change between versions as new features are introduced and existing functionality is improved.

---

# Contributing

Contributions, bug reports, feature requests, and improvements are welcome.

Before submitting a change, make sure your code:

* Follows standard Go formatting
* Uses clear and idiomatic Go
* Includes appropriate error handling
* Does not introduce unnecessary dependencies

Format the project with:

```bash
go fmt ./...
```

And run the tests with:

```bash
go test ./...
```

---

# License

This project is distributed under the license included in the repository.

If you use or redistribute `jsonlite`, please preserve the original copyright and attribution notices required by the project's license.

---

# Example Project Structure

A project using `jsonlite` could look like:

```text
my-project/
├── go.mod
├── main.go
├── config.json
└── ...
```

With:

```go
import "github.com/Reollinos/jsonlite-go"
```

The library itself can be organized as:

```text
jsonlite/
├── go.mod
├── jsonlite.go
├── README.md
└── ...
```

---

## Made with Go

`jsonlite` is built with Go and uses the standard library whenever possible.
