# suduko-go

build
```
go build -o bin
```

run
```
./bin --file sudukos/s01a.txt
```

oneliner
'''
go build -o bin && for file in ./sudukos/* ; do ./bin --file $file; done
```