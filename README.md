# Example

```go
package main

import (
	"github.com/semichkin-gopkg/mapping"
	"log"
)

func main() {
	m := mapping.New(
		map[int]string{
			0: "0",
			1: "1",
			2: "2",
		},
		mapping.WithDefaultLeft[int, string](100),
		mapping.WithDefaultRight[int, string]("100"),
	)

	log.Println(m.ToRight(0))        // "0"
	log.Println(m.ToRight(1))        // "1"
	log.Println(m.ToRight(-5))       // "100"
	log.Println(m.ToLeft("0"))       // 0
	log.Println(m.ToLeft("unknown")) // 100
}
```