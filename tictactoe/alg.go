package main

import (
	"math/rand"
	"time"
)

func alg(game *map[int]string, placed *bool, turn *int, starter *int) {
	/* Possible Improvements
	- The algo detects when it can win in one move and executes that move
	  and does not execute a defensive move.
	- The algo detects when it cannot win in one move and executes defensive measures.
	- Basically better decision making.
	*/
	if (*turn) == 1 {
		c0, c1, c2, c3 := false, false, false, false
		if (*game)[0] != "" {
			c0 = true
		} else if (*game)[2] != "" {
			c1 = true
		} else if (*game)[6] != "" {
			c2 = true
		} else if (*game)[8] != "" {
			c3 = true
		}
		occupied := 0
		for i := 0; i < 8; i++ {
			if (*game)[i] != "" {
				occupied += 1
			}
		}
		if (*starter) == 0 && (c0 || c1 || c2 || c3) && occupied == 1 {
			(*game)[4] = "o"
			(*turn) = 0
			(*placed) = true
		}
		//horizontal check (offensive)
		if (*starter) == 1 {
			for i := 0; i < 9; i += 3 {
				if (*game)[i] == "o" && (*game)[i+1] == "o" && (*game)[i+2] == "" && (*turn) == 1 {
					(*game)[i+2] = "o"
					(*turn) = 0
					(*placed) = true
				} else if (*game)[i+1] == "o" && (*game)[i+2] == "o" && (*game)[i] == "" && (*turn) == 1 {
					(*game)[i] = "o"
					(*turn) = 0
					(*placed) = true
				} else if (*game)[i] == "o" && (*game)[i+2] == "o" && (*game)[i+1] == "" && (*turn) == 1 {
					(*game)[i+1] = "o"
					(*turn) = 0
					(*placed) = true
				}
			}
		}
		//vertical check (offensive)
		if (*starter) == 1 {
			for i := 0; i < 3; i++ {
				if (*game)[i] == "o" && (*game)[i+6] == "o" && (*game)[i+3] == "" && (*turn) == 1 {
					(*game)[i+3] = "o"
					(*turn) = 0
					(*placed) = true
				} else if (*game)[i] == "o" && (*game)[i+3] == "o" && (*game)[i+6] == "" && (*turn) == 1 {
					(*game)[i+6] = "o"
					(*turn) = 0
					(*placed) = true
				} else if (*game)[i+6] == "o" && (*game)[i+3] == "o" && (*game)[i] == "" && (*turn) == 1 {
					(*game)[i] = "o"
					(*turn) = 0
					(*placed) = true
				}
			}
		}
		//diagonal offensive
		if (*starter) == 1 {
			if (*game)[2] == "" && (*game)[4] == "o" && (*game)[6] == "o" && (*turn) == 1 {
				(*game)[2] = "o"
				(*turn) = 0
				(*placed) = true
			} else if (*game)[2] == "o" && (*game)[4] == "o" && (*game)[6] == "" && (*turn) == 1 {
				(*game)[6] = "o"
				(*turn) = 0
				(*placed) = true
			} else if (*game)[2] == "o" && (*game)[4] == "" && (*game)[6] == "o" && (*turn) == 1 {
				(*game)[4] = "o"
				(*turn) = 0
				(*placed) = true
			} else if (*game)[0] == "" && (*game)[4] == "o" && (*game)[8] == "o" && (*turn) == 1 {
				(*game)[0] = "o"
				(*turn) = 0
				(*placed) = true
			} else if (*game)[0] == "o" && (*game)[4] == "o" && (*game)[8] == "" && (*turn) == 1 {
				(*game)[8] = "o"
				(*turn) = 0
				(*placed) = true
			} else if (*game)[0] == "o" && (*game)[4] == "" && (*game)[8] == "o" && (*turn) == 1 {
				(*game)[4] = "o"
				(*turn) = 0
				(*placed) = true
			}
		}
		//horizontal check
		for i := 0; i < 9; i += 3 {
			if (*game)[i] == "x" && (*game)[i+1] == "x" && (*game)[i+2] == "" && (*turn) == 1 {
				(*game)[i+2] = "o"
				(*turn) = 0
				(*placed) = true
			} else if (*game)[i+1] == "x" && (*game)[i+2] == "x" && (*game)[i] == "" && (*turn) == 1 {
				(*game)[i] = "o"
				(*turn) = 0
				(*placed) = true
			} else if (*game)[i] == "x" && (*game)[i+2] == "x" && (*game)[i+1] == "" && (*turn) == 1 {
				(*game)[i+1] = "o"
				(*turn) = 0
				(*placed) = true
			}
		}

		//vertical check
		for i := 0; i < 2; i++ {
			if (*game)[i] == "x" && (*game)[i+6] == "x" && (*game)[i+3] == "" && (*turn) == 1 {
				// 1 2 3 4 5 6 7 8 9
				// 0 1 2 3 4 5 6 7 8
				// X . . O . . X . .
				// X . .
				// O . .
				// X . .
				(*game)[i+3] = "o"
				(*turn) = 0
				(*placed) = true
			} else if (*game)[i] == "x" && (*game)[i+3] == "x" && (*game)[i+6] == "" && (*turn) == 1 {
				// 1 2 3 4 5 6 7 8 9
				// 0 1 2 3 4 5 6 7 8
				// X . . X . . O . .
				// X . .
				// X . .
				// O . .
				(*game)[i+6] = "o"
				(*turn) = 0
				(*placed) = true
			} else if (*game)[i+6] == "x" && (*game)[i+3] == "x" && (*game)[i] == "" && (*turn) == 1 {
				// 1 2 3 4 5 6 7 8 9
				// 0 1 2 3 4 5 6 7 8
				// O . . X . . X . .
				// O . .
				// X . .
				// X . .
				(*game)[i] = "o"
				(*turn) = 0
				(*placed) = true
			}
		}

		//diagonal check
		if (*game)[2] == "" && (*game)[4] == "x" && (*game)[6] == "x" && (*turn) == 1 {
			// 1 2 3 4 5 6 7 8 9
			// 0 1 2 3 4 5 6 7 8
			// . . O . X . X . .
			// . . O
			// . X .
			// X . .
			(*game)[2] = "o"
			(*turn) = 0
			(*placed) = true
		} else if (*game)[2] == "x" && (*game)[4] == "x" && (*game)[6] == "" && (*turn) == 1 {
			// 1 2 3 4 5 6 7 8 9
			// 0 1 2 3 4 5 6 7 8
			// . . X . X . O . .
			// . . X
			// . X .
			// O . .
			(*game)[6] = "o"
			(*turn) = 0
			(*placed) = true
		} else if (*game)[2] == "x" && (*game)[4] == "" && (*game)[6] == "x" && (*turn) == 1 {
			// 1 2 3 4 5 6 7 8 9
			// 0 1 2 3 4 5 6 7 8
			// . . X . O . X . .
			// . . X
			// . O .
			// X . .
			(*game)[4] = "o"
			(*turn) = 0
			(*placed) = true
		} else if (*game)[0] == "" && (*game)[4] == "x" && (*game)[8] == "x" && (*turn) == 1 {
			// 1 2 3 4 5 6 7 8 9
			// 0 1 2 3 4 5 6 7 8
			// O . . . X . . . X
			// O . .
			// . X .
			// . . X
			(*game)[0] = "o"
			(*turn) = 0
			(*placed) = true
		} else if (*game)[0] == "x" && (*game)[4] == "x" && (*game)[8] == "" && (*turn) == 1 {
			// 1 2 3 4 5 6 7 8 9
			// 0 1 2 3 4 5 6 7 8
			// X . . . X . . . O
			// X . .
			// . X .
			// . . O
			(*game)[8] = "o"
			(*turn) = 0
			(*placed) = true
		} else if (*game)[0] == "x" && (*game)[4] == "" && (*game)[8] == "x" && (*turn) == 1 {
			// 1 2 3 4 5 6 7 8 9
			// 0 1 2 3 4 5 6 7 8
			// X . . . O . . . X
			// X . .
			// . O .
			// . . X
			(*game)[4] = "o"
			(*turn) = 0
			(*placed) = true
		}

		//horizontal check (offensive)
		for i := 0; i < 9; i += 3 {
			if (*game)[i] == "o" && (*game)[i+1] == "" && (*game)[i+2] == "" && (*turn) == 1 {
				(*game)[i+1] = "o"
				(*turn) = 0
				(*placed) = true
			} else if (*game)[i] == "" && (*game)[i+1] == "o" && (*game)[i+2] == "" && (*turn) == 1 {
				(*game)[i] = "o"
				(*turn) = 0
				(*placed) = true
			} else if (*game)[i] == "" && (*game)[i+1] == "" && (*game)[i+2] == "o" && (*turn) == 1 {
				(*game)[i+1] = "o"
				(*turn) = 0
				(*placed) = true
			}
		}

		//vertical check (offensive)
		for i := 0; i < 3; i++ {
			if (*game)[i] == "o" && (*game)[i+3] == "" && (*game)[i+6] == "" && (*turn) == 1 {
				(*game)[i+3] = "o"
				(*turn) = 0
				(*placed) = true
			} else if (*game)[i] == "" && (*game)[i+3] == "o" && (*game)[i+6] == "" && (*turn) == 1 {
				(*game)[i] = "o"
				(*turn) = 0
				(*placed) = true
			} else if (*game)[i] == "" && (*game)[i+3] == "" && (*game)[i+6] == "o" && (*turn) == 1 {
				(*game)[i+3] = "o"
				(*turn) = 0
				(*placed) = true
			}
		}
		//diagonal check (offensive)
		if (*game)[0] == "o" && (*game)[4] == "" && (*game)[8] == "o" && (*turn) == 1 {
			(*game)[4] = "o"
			(*turn) = 0
			(*placed) = true
		} else if (*game)[2] == "o" && (*game)[4] == "" && (*game)[6] == "" && (*turn) == 1 {
			(*game)[4] = "o"
			(*turn) = 0
			(*placed) = true
		} else if (*game)[2] == "" && (*game)[4] == "o" && (*game)[6] == "" && (*turn) == 1 {
			(*game)[2] = "o"
			(*turn) = 0
			(*placed) = true
		} else if (*game)[2] == "" && (*game)[4] == "" && (*game)[6] == "o" && (*turn) == 1 {
			(*game)[4] = "o"
			(*turn) = 0
			(*placed) = true
		} else if (*game)[0] == "o" && (*game)[4] == "" && (*game)[8] == "" && (*turn) == 1 {
			(*game)[4] = "o"
			(*turn) = 0
			(*placed) = true
		} else if (*game)[0] == "" && (*game)[4] == "o" && (*game)[8] == "" && (*turn) == 1 {
			(*game)[0] = "o"
			(*turn) = 0
			(*placed) = true
		} else if (*game)[0] == "" && (*game)[4] == "" && (*game)[8] == "o" && (*turn) == 1 {
			(*game)[4] = "o"
			(*turn) = 0
			(*placed) = true
		}
		//place
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		if occupied != 8 && (!c0 && !c1 && !c2 && !c3) {
			for (*placed) == false {
				num := r1.Intn(8)
				//professional tic tac toe players start in the corners
				if num == 0 || num == 2 || num == 6 || num == 8 {
					if (*game)[num] == "" && (*game)[num] != "x" && (*game)[num] != "o" && (*turn) == 1 {
						(*game)[num] = "o"
						(*placed) = true
						(*turn) = 0
					}
				}
			}
		}
		for (*placed) == false {
			num := r1.Intn(8)
			//professional tic tac toe players start in the corners
			if (*game)[num] == "" && (*game)[num] != "x" && (*game)[num] != "o" && (*turn) == 1 {
				(*game)[num] = "o"
				(*placed) = true
				(*turn) = 0
			}
		}
		(*placed) = false
		(*turn) = 0
	}
}
