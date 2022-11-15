# go-simplex
Go binding for Simplex API 

## Installation

Use go get.

	go get github.com/sugarblock/go-simplex

Then import the go-simplex package into your own code.

	import (
        "github.com/sugarblock/go-simplex"
	    v2 "github.com/sugarblock/go-simplex/api/v2"
        )

## Usage

### GetAllEvent:

```go
import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/sugarblock/go-simplex"
)

func main() {
	cli, err := simplex.NewClient(nil, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	events, err := cli.GetAllEvent(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	json, _ := json.Marshal(events)
	fmt.Println(string(json))
}
```

## License

MIT