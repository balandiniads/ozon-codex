package main

import (
	"fmt"

	"github.com/diphantxm/ozon-api-client/ozon"
)

func main() {
	// Create a client with your Client-Id and Api-Key
	// [Documentation]: https://docs.ozon.ru/api/seller/en/#tag/Auth
	opts := []ozon.ClientOption{
		ozon.WithAPIKey("63ed3fb3-dab9-4e444f"),
		ozon.WithClientId("44459"),
	}
	client := ozon.NewClient(opts...)

	fmt.Printf("Barcode %s\n", client)

}
