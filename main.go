package main

import (
	"fmt"

	"github.com/diphantxm/ozon-api-client/ozon"
)

func main() {
	// Create a client with your Client-Id and Api-Key
	// [Documentation]: https://docs.ozon.ru/api/seller/en/#tag/Auth
	opts := []ozon.ClientOption{
		ozon.WithAPIKey("63ed3fb3-dab9-4e97-82e4-604c446ce1ef"),
		ozon.WithClientId("1344259"),
	}
	client := ozon.NewClient(opts...)

	fmt.Printf("Barcode %s\n", client)

}
