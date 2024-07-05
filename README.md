# sugoku

Simple CLI binary written in go to solve [sudoku](https://en.wikipedia.org/wiki/Sudoku) puzzles from `.txt` inputs.

The numbers should be space delimited with a missing number denoted with a `0`. 
See the `sudokus/` dir for examples.

I decided against using a brute force algorithm as I wanted to exercise automating the algorithms I use by hand.

## build the binary
```
go build -o bin
```

## solve a sudoku
```
./bin --file sudokus/s01a.txt
```
```
output:

file name:  sudokus/s01a.txt
+-------+-------+-------+
|   4   |       | 1 7 9 |
|     2 |     8 |   5 4 |
|     6 |     5 | 3   8 |
+-------+-------+-------+
|   8   |   7   | 9 1   |
|   5   |   9   |   3   |
|   1 9 |   6   |   4   |
+-------+-------+-------+
| 3     | 4     | 7     |
| 5 7   | 1     | 2     |
| 9 2 8 |       |   6   |
+-------+-------+-------+
Cells completed (76/81)
Cells completed (81/81)
Cells completed (81/81)
sudokus/s01a.txt 100 % Completed
+-------+-------+-------+
| 8 4 5 | 6 3 2 | 1 7 9 |
| 7 3 2 | 9 1 8 | 6 5 4 |
| 1 9 6 | 7 4 5 | 3 2 8 |
+-------+-------+-------+
| 6 8 3 | 5 7 4 | 9 1 2 |
| 4 5 7 | 2 9 1 | 8 3 6 |
| 2 1 9 | 8 6 3 | 5 4 7 |
+-------+-------+-------+
| 3 6 1 | 4 2 9 | 7 8 5 |
| 5 7 4 | 1 8 6 | 2 9 3 |
| 9 2 8 | 3 5 7 | 4 6 1 |
+-------+-------+-------+
```

## dirty one liner
Incase you wanted to test solving all the examples in a single sweep, here is a hacky bash command just for you.
```
go build -o bin && for file in ./sudokus/* ; do ./bin --file $file; done | grep %
```

```
output:

./sudokus/s01a.txt 100 % Completed
./sudokus/s01b.txt 100 % Completed
./sudokus/s01c.txt 100 % Completed
./sudokus/s02a.txt 100 % Completed
./sudokus/s02b.txt 100 % Completed
./sudokus/s02c.txt 100 % Completed
./sudokus/s03a.txt 100 % Completed
./sudokus/s03b.txt 100 % Completed
./sudokus/s03c.txt 61.728394 % Completed
./sudokus/s04a.txt 53.08642 % Completed
./sudokus/s04b.txt 43.209877 % Completed
./sudokus/s04c.txt 41.975307 % Completed
./sudokus/s05a.txt 49.382717 % Completed
./sudokus/s05b.txt 34.5679 % Completed
./sudokus/s05c.txt 54.320984 % Completed
./sudokus/s06a.txt 100 % Completed
./sudokus/s06b.txt 100 % Completed
./sudokus/s06c.txt 100 % Completed
./sudokus/s07a.txt 100 % Completed
./sudokus/s07b.txt 100 % Completed
./sudokus/s07c.txt 100 % Completed
./sudokus/s08a.txt 100 % Completed
./sudokus/s08b.txt 49.382717 % Completed
./sudokus/s08c.txt 67.90124 % Completed
./sudokus/s09a.txt 40.74074 % Completed
./sudokus/s09b.txt 41.975307 % Completed
./sudokus/s09c.txt 59.25926 % Completed
./sudokus/s10a.txt 100 % Completed
./sudokus/s10b.txt 100 % Completed
./sudokus/s10c.txt 100 % Completed
./sudokus/s11a.txt 100 % Completed
./sudokus/s11b.txt 100 % Completed
./sudokus/s11c.txt 100 % Completed
./sudokus/s12a.txt 30.864197 % Completed
./sudokus/s12b.txt 44.444447 % Completed
./sudokus/s12c.txt 38.271606 % Completed
./sudokus/s13a.txt 100 % Completed
./sudokus/s13b.txt 100 % Completed
./sudokus/s13c.txt 100 % Completed
./sudokus/s14a.txt 100 % Completed
./sudokus/s14b.txt 100 % Completed
./sudokus/s14c.txt 100 % Completed
./sudokus/s15a.txt 64.197525 % Completed
./sudokus/s15b.txt 44.444447 % Completed
./sudokus/s15c.txt 54.320984 % Completed
./sudokus/s16.txt 29.62963 % Completed
```
