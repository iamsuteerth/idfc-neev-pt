package rangeloops

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
// Basically LHS is key or index whereas RHS is the value associated with that index or key.
func CountInFile(cb Chessboard, file string) int {
	var squareOccupied int = 0
	for _, square := range cb[file] {
		if square {
			squareOccupied += 1
		}
	}
	return squareOccupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	var squareOccupied int = 0
	for _, file := range cb {
		if file[rank-1] {
			squareOccupied += 1
		}
	}
	return squareOccupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for range file {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for _, square := range file {
			if square {
				count++
			}
		}
	}
	return count
}
