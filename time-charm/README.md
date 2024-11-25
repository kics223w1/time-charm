# numutils

A simple library to format time in milliseconds to a human readable string.
Port from [pretty-ms](https://github.com/sindresorhus/pretty-ms).
Checkout my other open source projects at [kics223w1](https://github.com/kics223w1).
Welcome to contribute to this project.

# Installation

```bash
go get github.com/kics223w1/time-charm
```

# Usage

```go
import "github.com/kics223w1/time-charm"

options := timecharm.Options{
    MillisecondsDecimalDigits: 3,
}
fmt.Printf("33.333:  %s\n", timecharm.PrettyMilliseconds(float64(33.333), options))
```
