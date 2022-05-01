package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

func main() {

	// fast
	genSuitKeys(40)
	genFiveFlushKeys(13)
	genSevenFlushKeys(13)

	// about 1mn
	genFiveFaceKeys(13)

	// about 2h
	genSevenFaceKeys(13)
}

// ---------------------------------
// ---------------------------------
// ---------------Key Gen Suit Keys

func genSuitKeys(bound int) {
	// generate keys for suits Spades, Hearts, Diamonds, Clubs
	// keys are such that the sums of any 2 combinations of 7 suits are distincts
	// (discarding all other card info)

	fmt.Println(" ")
	defer track(runningtime("genSuitKeys"))

	if bound == 0 {
		bound = 40
	}

	var solutions [][]int
	var keys [4]int
	var valid bool

	sumKeys := make([]int, bound*7)
	nbSol := 0

	for k0 := 0; k0 < bound; k0++ {
		for k1 := k0 + 1; k1 < bound; k1++ {
			for k2 := k1 + 1; k2 < bound; k2++ {
				for k3 := k2 + 1; k3 < bound; k3++ {
					keys[0] = k0
					keys[1] = k1
					keys[2] = k2
					keys[3] = k3

					c := 0
					for c1 := 0; c1 < 4; c1++ {
						for c2 := c1; c2 < 4; c2++ {
							for c3 := c2; c3 < 4; c3++ {
								for c4 := c3; c4 < 4; c4++ {
									for c5 := c4; c5 < 4; c5++ {
										for c6 := c5; c6 < 4; c6++ {
											for c7 := c6; c7 < 4; c7++ {
												sumKeys[c] = keys[c1] + keys[c2] + keys[c3] +
													keys[c4] + keys[c5] + keys[c6] + keys[c7]
												c++
											}
										}
									}
								}
							}
						}
					}

					valid = true
					i := 0
					for (valid) && (i < c-1) {
						j := i + 1
						for (valid) && (j < c) {
							if sumKeys[i] == sumKeys[j] {
								valid = false
							}
							j++
						}
						i++
					}

					if valid {
						store := make([]int, len(keys))
						copy(store, keys[:])
						fmt.Println(store)
						solutions = append(solutions, store)
						nbSol++
						if nbSol >= 10 {
							fmt.Printf("early terminate after %d solutions\n", nbSol)
							saveResult(solutions, "suit-keys")
							return
						}
					}

				}
			}
		}
	}
	saveResult(solutions, "suit-keys")
	fmt.Println("done")
}

// ---------------------------------
// ---------------------------------
// ---------------Key Gen Flush Keys - Five

func genFiveFlushKeys(nbKeys int) {
	// generate keys for faces 1, 2, 3,.., 9, T, J, Q, K, A
	// keys are such that the sums of any 2 combinations of 5 distinct faces are distincts
	// (discarding all other card info)

	fmt.Println(" ")
	defer track(runningtime("genFiveFlushKeys"))

	if nbKeys == 0 {
		nbKeys = 13
	}

	var solutions [][]int
	keys := make([]int, nbKeys)

	sumKeys := make([]int, 50000)
	var valid bool

	// bootstrapping
	keys[0] = 0
	keys[1] = 1
	keys[2] = 2
	keys[3] = 4
	keys[4] = 8
	keys[5] = 16
	keys[6] = 32
	k := 7
	fmt.Printf("bootstrapping -> k=%d: keys=%v\n", k, keys)

	t0 := time.Now()

	for k < nbKeys {
		t := keys[k-1] + 1
		valid = false

		for !valid {
			keys[k] = t
			c := 0
			for c1 := 0; c1 < k+1; c1++ {
				for c2 := c1 + 1; c2 < k+1; c2++ {
					for c3 := c2 + 1; c3 < k+1; c3++ {
						for c4 := c3 + 1; c4 < k+1; c4++ {
							for c5 := c4 + 1; c5 < k+1; c5++ {
								sumKeys[c] = keys[c1] + keys[c2] + keys[c3] + keys[c4] + keys[c5]
								c++
							}

						}
					}
				}
			}

			i := 0
			valid = true
			for (valid) && (i < c-1) {
				j := i + 1
				for (valid) && (j < c) {
					if sumKeys[i] == sumKeys[j] {
						valid = false
						t++
					}
					j++
				}
				i++
			}
		}
		t1 := time.Now()
		dt := t1.Sub(t0)
		fmt.Printf("k=%d: keys=%v - t=%.2f s\n", k, keys, dt.Seconds())
		solutions = append(solutions, keys[:])
		k++
	}

	saveResult(solutions[len(solutions)-1:], "five-flush-keys")
	fmt.Println("done")
}

// ---------------------------------
// ---------------------------------
// ---------------Key Gen Face Keys - Five

func genFiveFaceKeys(nbKeys int) {
	// generate keys for faces 1, 2, 3,.., 9, T, J, Q, K, A
	// keys are such that the sums of any 2 combinations of 5 faces (with max 4 repetition) are distincts
	// (discarding all other card info)

	fmt.Println(" ")
	defer track(runningtime("genFiveFaceKeys"))

	if nbKeys == 0 {
		nbKeys = 13
	}

	var solutions [][]int
	keys := make([]int, nbKeys)

	sumKeys := make([]int, 50000)
	var valid bool

	// bootstrapping
	keys[0] = 0
	keys[1] = 1
	keys[2] = 5
	k := 3
	fmt.Printf("bootstrapping -> k=%d: keys=%v\n", k, keys)

	t0 := time.Now()

	for k < nbKeys {
		t := keys[k-1] + 1
		valid = false

		for !valid {
			keys[k] = t
			c := 0
			for c1 := 0; c1 < k+1; c1++ {
				for c2 := c1; c2 < k+1; c2++ {
					for c3 := c2; c3 < k+1; c3++ {
						for c4 := c3; c4 < k+1; c4++ {
							for c5 := c4; c5 < k+1; c5++ {
								if c1 != c5 {
									sumKeys[c] = keys[c1] + keys[c2] + keys[c3] + keys[c4] + keys[c5]
									c++

								}
							}

						}
					}
				}
			}

			i := 0
			valid = true
			for (valid) && (i < c-1) {
				j := i + 1
				for (valid) && (j < c) {
					if sumKeys[i] == sumKeys[j] {
						valid = false
						t++
					}
					j++
				}
				i++
			}
		}
		t1 := time.Now()
		dt := t1.Sub(t0)
		fmt.Printf("k=%d: keys=%v - t=%.2f s\n", k, keys, dt.Seconds())
		solutions = append(solutions, keys[:])
		k++
	}

	saveResult(solutions[len(solutions)-1:], "five-face-keys")
	fmt.Println("done")
}

// ---------------------------------
// ---------------------------------
// ---------------Key Gen Flush Keys - Seven

func genSevenFlushKeys(nbKeys int) {
	// generate keys for faces 1, 2, 3,.., 9, T, J, Q, K, A
	// keys are such that the sums of any 2 combinations of 5 or 6 or 7 distinct faces are distincts
	// (discarding all other card info)

	fmt.Println(" ")
	defer track(runningtime("genFiveFlushKeys"))

	if nbKeys == 0 {
		nbKeys = 13
	}

	var solutions [][]int
	keys := make([]int, nbKeys)

	sumKeys := make([]int, 50000)
	var valid bool

	// bootstrapping
	keys[0] = 1
	keys[1] = 2
	keys[2] = 4
	keys[3] = 8
	keys[4] = 16
	keys[5] = 32
	keys[6] = 64
	keys[7] = 128
	k := 8
	fmt.Printf("bootstrapping -> k=%d: keys=%v\n", k, keys)

	t0 := time.Now()

	for k < nbKeys {
		t := keys[k-1] + 1
		valid = false

		for !valid {
			keys[k] = t
			c := 0
			// 7 suited cards
			for c1 := 0; c1 < k+1; c1++ {
				for c2 := c1 + 1; c2 < k+1; c2++ {
					for c3 := c2 + 1; c3 < k+1; c3++ {
						for c4 := c3 + 1; c4 < k+1; c4++ {
							for c5 := c4 + 1; c5 < k+1; c5++ {
								for c6 := c5 + 1; c6 < k+1; c6++ {
									for c7 := c6 + 1; c7 < k+1; c7++ {
										sumKeys[c] = keys[c1] + keys[c2] + keys[c3] + keys[c4] + keys[c5] + keys[c6] + keys[c7]
										c++
									}
								}
							}
						}
					}
				}
			}
			// 6 suited cards
			for c1 := 0; c1 < k+1; c1++ {
				for c2 := c1 + 1; c2 < k+1; c2++ {
					for c3 := c2 + 1; c3 < k+1; c3++ {
						for c4 := c3 + 1; c4 < k+1; c4++ {
							for c5 := c4 + 1; c5 < k+1; c5++ {
								for c6 := c5 + 1; c6 < k+1; c6++ {
									sumKeys[c] = keys[c1] + keys[c2] + keys[c3] + keys[c4] + keys[c5] + keys[c6]
									c++
								}
							}
						}
					}
				}
			}

			// 5 suited cards
			for c1 := 0; c1 < k+1; c1++ {
				for c2 := c1 + 1; c2 < k+1; c2++ {
					for c3 := c2 + 1; c3 < k+1; c3++ {
						for c4 := c3 + 1; c4 < k+1; c4++ {
							for c5 := c4 + 1; c5 < k+1; c5++ {
								sumKeys[c] = keys[c1] + keys[c2] + keys[c3] + keys[c4] + keys[c5]
								c++
							}

						}
					}
				}
			}

			i := 0
			valid = true
			for (valid) && (i < c-1) {
				j := i + 1
				for (valid) && (j < c) {
					if sumKeys[i] == sumKeys[j] {
						valid = false
						t++
					}
					j++
				}
				i++
			}

		}
		t1 := time.Now()
		dt := t1.Sub(t0)
		fmt.Printf("k=%d: keys=%v - t=%.2f s\n", k, keys, dt.Seconds())
		solutions = append(solutions, keys[:])
		k++
	}
	saveResult(solutions[len(solutions)-1:], "seven-flush-keys")
	fmt.Println("done")

}

// ---------------------------------
// ---------------------------------
// ---------------Key Gen Face Keys - Five

func genSevenFaceKeys(nbKeys int) {
	// generate keys for faces 1, 2, 3,.., 9, T, J, Q, K, A
	// keys are such that the sums of any 2 combinations of 7 faces (with max 4 repetition) are distincts
	// (discarding all other card info)

	fmt.Println(" ")
	defer track(runningtime("genSevenFaceKeys"))

	if nbKeys == 0 {
		nbKeys = 13
	}

	var solutions [][]int
	keys := make([]int, nbKeys)

	sumKeys := make([]int, 50000)
	var valid bool

	// bootstrapping - empirical
	keys[0] = 0
	keys[1] = 1
	keys[2] = 5
	k := 3
	fmt.Printf("bootstrapping -> k=%d: keys=%v\n", k, keys)

	t0 := time.Now()

	for k < nbKeys {
		t := keys[k-1] + 1
		valid = false

		for !valid {
			keys[k] = t
			c := 0
			for c1 := 0; c1 < k+1; c1++ {
				for c2 := c1; c2 < k+1; c2++ {
					for c3 := c2; c3 < k+1; c3++ {
						for c4 := c3; c4 < k+1; c4++ {
							for c5 := c4; c5 < k+1; c5++ {
								for c6 := c5; c6 < k+1; c6++ {
									for c7 := c6; c7 < k+1; c7++ {
										if (c1 != c5) && (c2 != c6) && (c3 != c7) {
											sumKeys[c] = keys[c1] + keys[c2] + keys[c3] + keys[c4] + keys[c5] + keys[c6] + keys[c7]
											c++

										}
									}
								}
							}

						}
					}
				}
			}

			i := 0
			valid = true
			for (valid) && (i < c-1) {
				j := i + 1
				for (valid) && (j < c) {
					if sumKeys[i] == sumKeys[j] {
						valid = false
						t++
					}
					j++
				}
				i++
			}
		}
		t1 := time.Now()
		dt := t1.Sub(t0)
		fmt.Printf("k=%d: keys=%v - t=%.3f s\n", k, keys, dt.Seconds())

		// fmt.Printf("k=%d: keys=%v", k, keys)
		solutions = append(solutions, keys[:])
		k++
	}

	saveResult(solutions[len(solutions)-1:], "seven-face-keys")
	fmt.Println("done")
}

// ---------------------------------
// ---------------------------------
// ---------------Utils

func runningtime(s string) (string, time.Time) {
	log.Println("Start:	", s)
	return s, time.Now()
}

func track(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("End:	", s, s, "took", endTime.Sub(startTime))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func saveResult(data [][]int, name string) {
	_, thisFile, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(thisFile), "output", name+".csv")

	f, err := os.Create(filePath)
	check(err)
	defer f.Close()

	wr := csv.NewWriter(f)
	for _, arr := range data {
		st := strings.Fields(strings.Trim(fmt.Sprint(arr), "[]"))
		err = wr.Write(st)
		check(err)
	}
	wr.Flush()
	fmt.Println("saved solutions to /ouput/" + name + ".csv")
}
