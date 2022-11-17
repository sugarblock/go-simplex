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

Set **apikey** environment variable

```
$ export SIMPLEX_APIKEY=''
```

### GetQuote:

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/sugarblock/go-simplex"
	v2 "github.com/sugarblock/go-simplex/api/v2"
)

func main() {
	cli, err := simplex.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}
	request := &v2.QuoteRequest{
		EndUserId:         "11b111d1-161e-32d9-6bda-8dd2b5c8af17",
		DigitalCurrency:   "BTC",
		FiatCurrency:      "USD",
		RequestedCurrency: "USD",
		RequestedAmount:   50,
		WalletId:          "WalletID",
		ClientIp:          "1.2.3.4",
	}

	quote, err := cli.GetQuote(context.TODO(), request)
	if err != nil {
		log.Fatal(err)
	}

	json, _ := json.Marshal(quote)
	fmt.Println(string(json))
}

```

### GetAllEvent:

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/sugarblock/go-simplex"
)

func main() {
	cli, err := simplex.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	events, err := cli.GetAllEvents(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	json, _ := json.Marshal(events)
	fmt.Println(string(json))
}
```

## License

Apache License 2.0