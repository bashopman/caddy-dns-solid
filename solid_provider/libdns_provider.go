package solid_provider

import (
	"context"
	"fmt"
	"github.com/libdns/libdns"
)

// Provider facilitates DNS record manipulation with Solid.
type Provider struct {
	SolidApiUrl string `json:"solid_api_url"`
}

// GetRecords lists all the records in the zone.
func (p *Provider) GetRecords(ctx context.Context, zone string) ([]libdns.Record, error) {
	return nil, fmt.Errorf("TODO: not implemented")
}

// AppendRecords adds records to the zone. It returns the records that were added.
func (p *Provider) AppendRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	return nil, fmt.Errorf("TODO: not implemented")
}

// SetRecords sets the records in the zone, either by updating existing records or creating new ones.
// It returns the updated records.
func (p *Provider) SetRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	fmt.Println("libdns_provider.go:SetRecords")

	for _, record := range records {
		fmt.Printf("SetRecords: %s\n", record)
	}
	//resp, err = http.Post("https://acme-gateway.network.intra.lighting.com/authRequest",
	//	"application/json", bytes.NewBuffer([]byte(`{"subjectAltName": `+zone+`}`)))
	return records, nil
}

// DeleteRecords deletes the records from the zone. It returns the records that were deleted.
func (p *Provider) DeleteRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	fmt.Println("libdns_provider.go:DeleteRecords")

	for _, record := range records {
		fmt.Printf("DeleteRecords: %s\n", record)
	}
	return records, nil
}

// Interface guards
var (
	_ libdns.RecordGetter   = (*Provider)(nil)
	_ libdns.RecordAppender = (*Provider)(nil)
	_ libdns.RecordSetter   = (*Provider)(nil)
	_ libdns.RecordDeleter  = (*Provider)(nil)
)
