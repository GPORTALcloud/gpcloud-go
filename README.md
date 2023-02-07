### This repository is archived, please go to https://github.com/G-PORTAL/gpcloud-go for the latest version.

---

# GPCloud Golang Client

This is the official GPCloud Golang client. Please raise an issue if you have found any problems or having questions.

### Recommendations

- Golang 1.16 or higher

### Example usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GPORTALcloud/gpcloud-go/pkg/gpcloud/client"
	"github.com/GPORTALcloud/gpcloud-go/pkg/gpcloud/ptypes"
)

func main() {
	cl, err := client.NewClient("")
	if err != nil {
		log.Fatal(err)
	}
	keys, err := cl.PublicClient().ListJwtPublicKeys(context.Background(), &ptypes.EmptyRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for _, key := range keys.GetKeys() {
		fmt.Println(key.Kid)
	}
}
```
