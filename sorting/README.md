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
BenchmarkBubbleSort-4                     146353              8742 ns/op               0 B/op          0 allocs/op
BenchmarkCocktailShakerSort-4             134820              9766 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSort-4                  141255              8057 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort-4                      150060              7881 ns/op            5824 B/op         99 allocs/op
BenchmarkGnomeSort-4                      157920              6739 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSort-4                      282315              4637 ns/op               0 B/op          0 allocs/op
BenchmarkCombSort-4                       413973              4092 ns/op               0 B/op          0 allocs/op

1000
BenchmarkBubbleSort-4                       1304            770025 ns/op               0 B/op          0 allocs/op
BenchmarkCocktailShakerSort-4               1261            903892 ns/op               0 B/op          0 allocs/op
BenchmarkGnomeSort-4                        1713            707955 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSort-4                    2055            568802 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort-4                       12164            125558 ns/op           81536 B/op        999 allocs/op
BenchmarkQuickSort-4                       20682             60067 ns/op               0 B/op          0 allocs/op
BenchmarkCombSort-4                        21500             50499 ns/op               0 B/op          0 allocs/op

10000
BenchmarkBubbleSort-4                        100         110736128 ns/op               3 B/op          0 allocs/op
BenchmarkCocktailShakerSort-4                142          82933677 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSort-4                     218          60642253 ns/op               0 B/op          0 allocs/op
BenchmarkGnomeSort-4                         212          56417953 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort-4                        7542           1738135 ns/op         1127170 B/op       9999 allocs/op
BenchmarkQuickSort-4                       16023            880590 ns/op               0 B/op          0 allocs/op
BenchmarkCombSort-4                        17680            719814 ns/op               0 B/op          0 allocs/op
```