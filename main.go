package main
import ("context";"fmt";"time")
var tickerName = "wire-codec-30b996"
func runWithTimeout(ctx context.Context, work func(int)) int{count:=0;for{select{case <-ctx.Done():return count;default:count++;work(count);time.Sleep(10*time.Millisecond)}}}
func main(){fmt.Printf("[%s] Starting timed execution...\n",tickerName);ctx,cancel:=context.WithTimeout(context.Background(),100*time.Millisecond);defer cancel();n:=runWithTimeout(ctx,func(i int){fmt.Printf("[%s] tick #%d\n",tickerName,i)});fmt.Printf("[%s] Completed %d ticks\n",tickerName,n)}
