# Beagle
![](https://i.imgur.com/l2aB3vG.jpg)

A super simple regexp wrapper to make things more digestible for the average
usecase.

## Example Usage
``` go
package main

import (
    "fmt"

    "github.com/DaiHasso/beagle"
)

func main() {
    regex := MustRegex(`(?P<thing>\S+) (are|is) (?P<difficulty>\S+)`)

    result := regex.Match(
        "I heard that go regexps are hard! But maybe it is easy.",
    )

    fmt.Printf(
        "Actually %s %s %s\n",
        result.NamedGroup("thing")[0],
        result.UnamedGroups()[0],
        result.NamedGroup("difficulty")[1],
    )
}
```

Will result in:

>Actually regexps are easy.
