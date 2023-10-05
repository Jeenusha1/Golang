package contextss

import (
	"context"
	"fmt"
	"time"
)

func GetContext(ctx context.Context) context.Context{
	// withvalue in the context has the key and its pair
	return context.WithValue(ctx, "jee", "1234")
}
func GetContext1(ctx context.Context){
	for {
	 select {
		// done identified the context has finished executing
		// return a channel when context is closed or timeout
		case <- ctx.Done():
			fmt.Print("i am timed out")
			return
		default:
			fmt.Print("hereee")
		}
		time.Sleep(500 * time.Millisecond)
	} 
}


