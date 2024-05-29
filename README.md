# sudoku-go

build
```
go build -o bin
```

solve a sudoku
```
./bin --file sudokus/s01a.txt
```

coverage 1 liner
```
go build -o bin && for file in ./sudokus/* ; do ./bin --file $file; done | grep %
```