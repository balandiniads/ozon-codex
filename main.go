package main

import (
	"fmt"

	"github.com/diphantxm/ozon-api-client/ozon"
)

func main() {
	// Create a client with your Client-Id and Api-Key
	// [Documentation]: https://docs.ozon.ru/api/seller/en/#tag/Auth
	opts := []ozon.ClientOption{
		ozon.WithAPIKey("PUT DOWN UR API KEY HERE"),
		ozon.WithClientId("PUT DOWN UR CLIENT KEY HERE"),
	}
	client := ozon.NewClient(opts...)

	fmt.Printf("Barcode %s\n", client)

}
