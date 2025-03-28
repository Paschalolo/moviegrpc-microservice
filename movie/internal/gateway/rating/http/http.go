package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"movieexample.com/movie/internal/gateway"
	"movieexample.com/pkg/discovery"
	"movieexample.com/rating/pkg/model"
)

// Gateway defines a movie rating HTTP gateway
type Gateway struct {
	registry discovery.Registery
}

// New creates a new HTTP gateway for a rating service
func New(registry discovery.Registery) *Gateway {
	return &Gateway{registry}
}

// GetAggregatedRating returns the aggregated rating for a
// record or Errnotfound if there are no rating for it
func (g *Gateway) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error) {
	addrs, err := g.registry.ServiceAddresses(ctx, "rating")
	if err != nil {
		return 0, err
	}
	Url := "http://" + addrs[rand.Intn(len(addrs))] + "/rating"
	log.Println("calling rating services . Request:GET" + Url)
	req, err := http.NewRequest(http.MethodGet, Url, nil)
	if err != nil {
		return 0, nil
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", string(recordID))
	values.Add("type", fmt.Sprintf("%v", recordType))
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return 0, gateway.ErrNotFound
	} else if resp.StatusCode/100 != 2 {
		return 0, fmt.Errorf("non-2xx response: %v", resp)
	}
	var v float64
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return 0, nil
	}

	return v, nil

}

// Put Ratings writes a rating
func (g *Gateway) PutRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	addrs, err := g.registry.ServiceAddresses(ctx, "rating")
	if err != nil {
		return err
	}
	Url := "http://" + addrs[rand.Intn(len(addrs))] + "/rating"
	log.Println("calling rating services . Request:PUT" + Url)

	req, err := http.NewRequest(http.MethodPut, Url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", string(recordID))
	values.Add("type", fmt.Sprintf("%v", recordType))
	values.Add("userId", string(rating.UserID))
	values.Add("value", fmt.Sprintf("%v", rating.Value))
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("non -2xx response : %v ", resp)
	}
	return nil
}
