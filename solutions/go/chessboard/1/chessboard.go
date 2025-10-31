package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
    fileSlice := cb[file]
    
	for i, _ := range fileSlice {
        if fileSlice[i] == true {
            count++
        }
    }

    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0

    if rank >= 1 && rank <= 8 {
		for _, fileSlice := range cb {
            if fileSlice[rank - 1] == true {
                count++ 
            }
        }
    }

    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    count := 0

    for _, fileSlice := range cb {
		for range fileSlice {
            count++
        }
    }

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0

    for _, fileSlice := range cb {
		for i, _ := range fileSlice {
            if fileSlice[i] == true {
                count++ 
            }
        }
    }

	return count
}
