package poker

import "fmt"

const NB_FACE = 13
const NB_SUIT = 4
const DECK_SIZE = NB_SUIT * NB_FACE
const SUIT_MASK = 511
const SUIT_BIT_SHIFT = 9

var CARD_NO = make(map[string]int)
var CARD_SY = make(map[int]string)

//  (c)lubs, (d)iamonds, (h)earts, (s)pades
var SUIT = [4]string{"c", "d", "h", "s"}

var SUIT_KEY = [4]uint32{0, 1, 29, 37}

//  faces: 2, 3, 4, 5, 6, 7, 8, 9, T(10), (J)ack, (Q)ueen, (K)ing, (A)ce
var FACE = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

var FLUSH_FIVE_KEY = [NB_FACE]uint32{0, 1, 2, 4, 8, 16, 32, 56, 104, 192, 352, 672, 1288}
var FLUSH_SEVEN_KEY = [NB_FACE]uint32{1, 2, 4, 8, 16, 32, 64, 128, 240, 464, 896, 1728, 3328}

var FACE_FIVE_KEY = [NB_FACE]uint32{0, 1, 5, 22, 94, 312, 992, 2422, 5624, 12522, 19998, 43258, 79415}
var FACE_SEVEN_KEY = [NB_FACE]uint32{0, 1, 5, 22, 98, 453, 2031, 8698, 22854, 83661, 262349, 636345, 1479181}

var MAX_SUIT_KEY = SUIT_KEY[3] * 7

var MAX_FLUSH_FIVE_KEY = subSum(FLUSH_FIVE_KEY[:], 5)
var MAX_FLUSH_SEVEN_KEY = subSum(FLUSH_SEVEN_KEY[:], 7)

var MAX_FACE_FIVE_KEY = FACE_FIVE_KEY[NB_FACE-1]*4 + FACE_FIVE_KEY[NB_FACE-2]*1
var MAX_FACE_SEVEN_KEY = FACE_SEVEN_KEY[NB_FACE-1]*4 + FACE_SEVEN_KEY[NB_FACE-2]*3

var CARD_FACE [DECK_SIZE]int
var CARD_SUIT [DECK_SIZE]int

var CARD_FLUSH_KEY [DECK_SIZE]uint32
var CARD_FACE_KEY [DECK_SIZE]uint32

func InitKeys(verbose bool) {

	for k := range CARD_NO {
		delete(CARD_NO, k)
	}
	for k := range CARD_SY {
		delete(CARD_SY, k)
	}

	if MAX_SUIT_KEY >= 2^(1<<SUIT_BIT_SHIFT) {
		panic("suit keys are too large to be stored in SUIT_BIT_SHIFT bits")
	}

	s := SUIT_BIT_SHIFT
	if MAX_FACE_SEVEN_KEY >= 2^(32-(1<<s)) {
		panic("face keys are too large to be stored in 32-SUIT_BIT_SHIFT bits")
	}

	for f := 0; f < NB_FACE; f++ {
		for s := 0; s < NB_SUIT; s++ {
			n := NB_SUIT*f + s
			CARD_FACE[n] = f
			CARD_SUIT[n] = s

			CARD_FLUSH_KEY[n] = FLUSH_SEVEN_KEY[f]
			CARD_FACE_KEY[n] = (FACE_SEVEN_KEY[f] << SUIT_BIT_SHIFT) + SUIT_KEY[s]

			symbol := FACE[f] + SUIT[s]
			CARD_SY[n] = symbol
			CARD_NO[symbol] = n
		}
	}

	if verbose {
		showKeys()
	}

}

func showKeys() {
	fmt.Printf("checks\n")
	s := SUIT_BIT_SHIFT
	fmt.Printf("\tMAX_SUIT_KEY=%d < 2^SUIT_BIT_SHIFT=2^%d=%d ? %t\n", MAX_SUIT_KEY, s, 1<<s, MAX_SUIT_KEY < 2^(1<<s))
	fmt.Printf("\tMAX_FACE_SEVEN_KEY=%d < 2^(32-SUIT_BIT_SHIFT)=2^%d=%d ? %t\n", MAX_FACE_SEVEN_KEY, 32-s, 1<<(32-s), MAX_FACE_SEVEN_KEY < 2^(32-1<<s))

	fmt.Printf("\ncards\n")
	fmt.Printf("\tFACE = %v\n", FACE)
	fmt.Printf("\tSUIT = %v\n", SUIT)
	fmt.Printf("\tCARD_NO = %v\n", CARD_NO)
	fmt.Printf("\tCARD_SY = %v\n", CARD_SY)

	fmt.Printf("\neval keys\n")
	fmt.Printf("\tNB_FACE = %d\n", NB_FACE)
	fmt.Printf("\tNB_SUIT = %d\n", NB_SUIT)
	fmt.Printf("\tDECK_SIZE = %d\n", DECK_SIZE)
	fmt.Printf("\tSUIT_MASK = %d\n", SUIT_MASK)
	fmt.Printf("\tSUIT_BIT_SHIFT = %d\n", SUIT_BIT_SHIFT)

	fmt.Printf("\tSUIT_KEY = %v\n", SUIT_KEY)
	fmt.Printf("\tFLUSH_FIVE_KEY = %v\n", FLUSH_FIVE_KEY)
	fmt.Printf("\tFLUSH_SEVEN_KEY = %v\n", FLUSH_SEVEN_KEY)
	fmt.Printf("\tFACE_FIVE_KEY = %v\n", FACE_FIVE_KEY)
	fmt.Printf("\tFACE_SEVEN_KEY = %v\n", FACE_SEVEN_KEY)

	fmt.Printf("\tMAX_SUIT_KEY = %d\n", MAX_SUIT_KEY)
	fmt.Printf("\tMAX_FLUSH_FIVE_KEY = %d\n", MAX_FLUSH_FIVE_KEY)
	fmt.Printf("\tMAX_FLUSH_SEVEN_KEY = %d\n", MAX_FLUSH_SEVEN_KEY)
	fmt.Printf("\tMAX_FACE_FIVE_KEY = %d\n", MAX_FACE_FIVE_KEY)
	fmt.Printf("\tMAX_FACE_SEVEN_KEY = %d\n", MAX_FACE_SEVEN_KEY)

	fmt.Printf("\tCARD_FACE = %v\n", CARD_FACE)
	fmt.Printf("\tCARD_SUIT = %v\n", CARD_SUIT)

	fmt.Printf("\tCARD_FLUSH_KEY = %v\n", CARD_FLUSH_KEY)
	fmt.Printf("\tCARD_FACE_KEY = %v\n", CARD_FACE_KEY)

}

func subSum(arr []uint32, nLast int) uint32 {

	var sum uint32 = 0
	for i := len(arr) - nLast; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum

}
