package entities

func PieceTiles(index int) [][]bool {
	switch index {
	// 1 : 1
	case 0:
		return [][]bool{
			{true, true, false},
			{false, true, false},
			{false, true, true},
		}

	// 1 : 2
	case 1:
		return [][]bool{
			{true, true, false},
			{true, false, false},
			{true, false, false},
		}

	// 1 : 3
	case 2:
		return [][]bool{
			{true, true},
			{true, true},
		}

	// 1 : 4
	case 3:
		return [][]bool{
			{true, false},
			{true, false},
		}

	// 2 : 1
	case 4:
		return [][]bool{
			{true, true, true, true},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		}

	// 2 : 2
	case 5:
		return [][]bool{
			{true, true, true, true},
			{false, false, true, false},
			{false, false, false, false},
			{false, false, false, false},
		}

	// 3 : 1
	case 6:
		return [][]bool{
			{true, true, true},
			{false, true, true},
			{false, false, false},
		}

	// 3 : 2
	case 7:
		return [][]bool{
			{false, true, true},
			{true, true, false},
			{false, false, false},
		}

	// 3 : 3
	case 8:
		return [][]bool{
			{true, true, true},
			{false, false, false},
			{false, false, false},
		}

	// 4 : 1
	case 9:
		return [][]bool{
			{true, true, true, true},
			{false, false, false, true},
			{false, false, false, false},
			{false, false, false, false},
		}

	// 4 : 2
	case 10:
		return [][]bool{
			{true, true, true},
			{false, true, false},
			{false, false, false},
		}

	// 4 : 3
	case 11:
		return [][]bool{
			{true, true},
			{true, false},
		}

	default:
		return [][]bool{}
	}
}
