package business

import (
	"context"
	"fmt"
	"time"
)

func (b biz) GetSports(ctx context.Context) error {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	return nil
}
