package main

// entrypoint
func main() {

	s := ImportSudukoFromFile("s14a.txt")
	PrintSuduko(s)

	s.Solve()
	PrintSuduko(s)
}
