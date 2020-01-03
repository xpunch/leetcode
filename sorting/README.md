# Sorting

## Benchmark
```
go test -benchmem -benchtime=1s -bench=.
```
```
10
BenchmarkMergeSort-4                      140834              1280 ns/op             304 B/op          9 allocs/op
BenchmarkQuickSort-4                      322028               501 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSort-4                  299857               364 ns/op               0 B/op          0 allocs/op
BenchmarkCocktailShakerSort-4             400226               277 ns/op               0 B/op          0 allocs/op
BenchmarkGnomeSort-4                      453106               275 ns/op               0 B/op          0 allocs/op
BenchmarkBubbleSort-4                     639403               271 ns/op               0 B/op          0 allocs/op
BenchmarkCombSort-4                      1000000               236 ns/op               0 B/op          0 allocs/op

100
BenchmarkBubbleSort-4                     123136              9679 ns/op               0 B/op          0 allocs/op
BenchmarkCocktailShakerSort-4             121542              9602 ns/op               0 B/op          0 allocs/op
BenchmarkOddEvenSort-4                    117151              9615 ns/op               0 B/op          0 allocs/op
BenchmarkCombSort-4                       253772              4102 ns/op               0 B/op          0 allocs/op
BenchmarkGnomeSort-4                      146133              6984 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSort-4                  154164              7501 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort-4                      158782              8441 ns/op            5824 B/op         99 allocs/op
BenchmarkQuickSort-4                      234450              4625 ns/op               0 B/op          0 allocs/op
BenchmarkHeapSort-4                       138548              8372 ns/op               0 B/op          0 allocs/op

1000
BenchmarkBubbleSort-4           	    1465	    800408 ns/op	       0 B/op	       0 allocs/op
BenchmarkCocktailShakerSort-4   	    1444	    823468 ns/op	       0 B/op	       0 allocs/op
BenchmarkOddEvenSort-4          	    1794	    676167 ns/op	       0 B/op	       0 allocs/op
BenchmarkGnomeSort-4            	    2104	    545513 ns/op	       0 B/op	       0 allocs/op
BenchmarkHeapSort-4             	    2352	    514212 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectionSort-4        	    2499	    491969 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSort-4        	    3636	    346135 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountingSort-4         	    9996	    163017 ns/op	    8192 B/op	       1 allocs/op
BenchmarkMergeSort-4            	   12874	     94542 ns/op	   81536 B/op	     999 allocs/op
BenchmarkShellSort-4            	   20149	     60397 ns/op	       0 B/op	       0 allocs/op
BenchmarkCombSort-4             	   19312	     62997 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort-4            	   20307	     60566 ns/op	       0 B/op	       0 allocs/op

10000
BenchmarkBubbleSort-4                        100         110736128 ns/op               3 B/op          0 allocs/op
BenchmarkCocktailShakerSort-4                142          82933677 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSort-4                     218          60642253 ns/op               0 B/op          0 allocs/op
BenchmarkGnomeSort-4                         212          56417953 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort-4                        7542           1738135 ns/op         1127170 B/op       9999 allocs/op
BenchmarkQuickSort-4                       16023            880590 ns/op               0 B/op          0 allocs/op
BenchmarkCombSort-4                        17680            719814 ns/op               0 B/op          0 allocs/op
```