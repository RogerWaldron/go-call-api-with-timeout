package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "username", "fred" )

	userID, err := fetchUserID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response took %v and returned: %+v\n", time.Since(start), userID)
}

func fetchUserID(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 200)
	defer cancel()

	userName := ctx.Value("username").(string)

	type result struct {
		userId string
		err 	error
	}

	resultCh := make(chan result, 1)

	go func() {
		res, err := callThirdPartyAPI(userName)
		resultCh <- result{
			userId: res,
			err: err,
		}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultCh:
		return res.userId, res.err
	}
}

func callThirdPartyAPI(username string) (string, error){
	// time.Sleep(time.Millisecond * 500) // context deadline exceeded
	time.Sleep(time.Millisecond * 50)
	return  "1", nil
}