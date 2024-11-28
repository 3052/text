# Overview

package `text`

## Index

- [Variables](#variables)
- [Functions](#functions)
  - [func Clean(s string) string](#func-clean)
  - [func CutBefore(s, sep []byte) ([]byte, []byte, bool)](#func-cutbefore)
  - [func Name(n Namer) (string, error)](#func-name)
- [Types](#types)
  - [type Cardinal](#type-cardinal)
    - [func (c Cardinal) String() string](#func-cardinal-string)
  - [type Namer](#type-namer)
  - [type Percent](#type-percent)
    - [func (p Percent) String() string](#func-percent-string)
  - [type ProgressMeter](#type-progressmeter)
    - [func (p \*ProgressMeter) Reader(resp \*http.Response) io.Reader](#func-progressmeter-reader)
    - [func (p \*ProgressMeter) Set(parts int)](#func-progressmeter-set)
    - [func (p \*ProgressMeter) Write(data []byte) (int, error)](#func-progressmeter-write)
  - [type Rate](#type-rate)
    - [func (r Rate) String() string](#func-rate-string)
  - [type Size](#type-size)
    - [func (s Size) String() string](#func-size-string)
  - [type Transport](#type-transport)
    - [func (Transport) RoundTrip(req \*http.Request) (\*http.Response, error)](#func-transport-roundtrip)
    - [func (Transport) Set(on bool)](#func-transport-set)
- [Source files](#source-files)

## Variables

```go
var DefaultName = "{{if .Show}}" +
  "{{if .Season}}" +
  "{{if .Title}}" +
  "{{.Show}} - {{.Season}} {{.Episode}} - {{.Title}}" +
  "{{else}}" +
  "{{.Show}} - {{.Season}} {{.Episode}}" +
  "{{end}}" +
  "{{else}}" +
  "{{.Show}} - {{.Title}}" +
  "{{end}}" +
  "{{else}}" +
  "{{if .Year}}" +
  "{{.Title}} - {{.Year}}" +
  "{{else}}" +
  "{{.Title}}" +
  "{{end}}" +
  "{{end}}"
```

```go
var DefaultTransport = http.DefaultTransport
```

## Functions

### func [Clean](./text.go#L34)

```go
func Clean(s string) string
```

### func [CutBefore](./text.go#L112)

```go
func CutBefore(s, sep []byte) ([]byte, []byte, bool)
```

### func [Name](./text.go#L44)

```go
func Name(n Namer) (string, error)
```

## Types

### type [Cardinal](./text.go#L138)

```go
type Cardinal float64
```

### func (Cardinal) [String](./text.go#L140)

```go
func (c Cardinal) String() string
```

### type [Namer](./text.go#L57)

```go
type Namer interface {
  Show() string
  Season() int
  Episode() int
  Title() string
  Year() int
}
```

### type [Percent](./text.go#L162)

```go
type Percent float64
```

### func (Percent) [String](./text.go#L164)

```go
func (p Percent) String() string
```

### type [ProgressMeter](./text.go#L70)

```go
type ProgressMeter struct {
  // contains filtered or unexported fields
}
```

### func (\*ProgressMeter) [Reader](./text.go#L94)

```go
func (p *ProgressMeter) Reader(resp *http.Response) io.Reader
```

### func (\*ProgressMeter) [Set](./text.go#L64)

```go
func (p *ProgressMeter) Set(parts int)
```

### func (\*ProgressMeter) [Write](./text.go#L101)

```go
func (p *ProgressMeter) Write(data []byte) (int, error)
```

### type [Rate](./text.go#L150)

```go
type Rate float64
```

### func (Rate) [String](./text.go#L152)

```go
func (r Rate) String() string
```

### type [Size](./text.go#L169)

```go
type Size float64
```

### func (Size) [String](./text.go#L171)

```go
func (s Size) String() string
```

### type [Transport](./text.go#L190)

```go
type Transport struct{}
```

### func (Transport) [RoundTrip](./text.go#L192)

```go
func (Transport) RoundTrip(req *http.Request) (*http.Response, error)
```

### func (Transport) [Set](./text.go#L181)

```go
func (Transport) Set(on bool)
```

## Source files

[text.go](./text.go)
