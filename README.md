# sudoku-go

build
```
go build -o bin
```

run
```
./bin --file sudokus/s01a.txt
```

oneliner
```
go build -o bin && for file in ./sudokus/* ; do ./bin --file $file; done
```