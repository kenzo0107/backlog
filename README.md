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
	api := backlog.New("YOUR API KEY", "YOUR BASE URL")

	user, err := api.GetUserMySelf()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("user ID: %d, Name %s\n", user.ID, user.Name)
}
```

## License

[MIT License](https://github.com/kenzo0107/backlog/blob/master/LICENSE)
