Backlog API in Go [![GoDoc](https://godoc.org/github.com/kenzo0107/backlog?status.svg)](https://godoc.org/github.com/kenzo0107/backlog) [![test](https://github.com/kenzo0107/backlog/actions/workflows/test.yml/badge.svg)](https://github.com/kenzo0107/backlog/actions/workflows/test.yml) [![lint](https://github.com/kenzo0107/backlog/actions/workflows/lint.yml/badge.svg)](https://github.com/kenzo0107/backlog/actions/workflows/lint.yml)
[![codecov](https://codecov.io/gh/kenzo0107/backlog/branch/master/graph/badge.svg)](https://codecov.io/gh/kenzo0107/backlog)
===============

This library supports most if not all of the `backlog` REST calls.


## Installing

### *go get*

    $ go get -u github.com/kenzo0107/backlog

## Example

### Get my user information

```go
package main

import (
	"fmt"
	"os"

	"github.com/kenzo0107/backlog"
)

func main() {
	c := backlog.New("YOUR API KEY", "YOUR BASE URL")

	user, err := c.GetUserMySelf()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("user ID: %d, Name %s\n", user.ID, user.Name)
}
```

### Download space icon

```go
func main() {
	file, err := os.Create("icon.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	c := backlog.New("YOUR API KEY", "YOUR BASE URL")

	if err := c.GetSpaceIcon(file); err != nil {
		fmt.Println(err)
		return
	}
}
```

## Contributing

You are more than welcome to contribute to this project. Fork and make a Pull Request, or create an Issue if you see any problem.

Before making any Pull Request please run the following:

```
make pr-prep
```

This will check/update code formatting, linting and then run all tests

## License

[MIT License](https://github.com/kenzo0107/backlog/blob/master/LICENSE)
