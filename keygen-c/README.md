# Search for Keys

We search 5 set of keys:

- suit keys for 7 cards
- flush keys for 5 cards
- face keys for 5 cards
- flush keys for 5 cards
- face keys for 7 cards

## Suit Keys

Each suit (Spades, Hearts, Diamonds, Clubs) is mapped to a **Suit Key**.
The Suit Keys are such that the sums of any 2 combinations of 7 cards are distincts (discarding all other card info).

```bash
gcc -Wall -g -O3 key-gen-suit-seven.c -o key-gen-suit-seven-exec
./key-gen-suit-seven-exec

# estimated time <1s
```

## Flush Keys for 5 Cards

Each face (1, 2, 3,.., 9, T, J, Q, K, A) is mapped to a **Flush-5 Key**.
The Flush-5 Keys are such that the sums of any 2 combinations of 5 distinct faces are distincts (discarding all other card info).

```bash
gcc -Wall -g -O3 key-gen-flush-five.c -o key-gen-flush-five-exec
./key-gen-flush-five-exec

# estimated time <1s
```

## Flush Keys for 7 Cards

Each face (1, 2, 3,.., 9, T, J, Q, K, A) is mapped to a **Flush-7 Key**.
The Flush-7 Keys are such that the sums of any 2 combinations of **5 or 6 or 7** distinct faces are distincts (discarding all other card info).

```bash
gcc -Wall -g -O3 key-gen-flush-seven.c -o key-gen-flush-seven-exec
./key-gen-flush-seven-exec

# estimated time <1s
```

## Face Keys for 5 Cards

Each face (1, 2, 3,.., 9, T, J, Q, K, A) is mapped to a **Face-5 Key**.
The Face-5 Keys are such that the sums of any 2 combinations of 5 faces **with max 4 repetition** are distincts (discarding all other card info).

```bash
gcc -Wall -g -O3 key-gen-face-five.c -o key-gen-face-five-exec
./key-gen-face-five-exec

# estimated time ~30s
```

## Face Keys for 7 Cards

Each face (1, 2, 3,.., 9, T, J, Q, K, A) is mapped to a **Face-7 Key**.
The Face-7 Keys are such that the sums of any 2 combinations of 7 faces **with max 4 repetition** are distincts (discarding all other card info).

```bash
gcc -Wall -g -O3 key-gen-face-seven.c -o key-gen-face-seven-exec
./key-gen-face-seven-exec

# estimated time ~4-5h
```
