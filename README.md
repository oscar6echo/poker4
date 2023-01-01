# Poker

## Overview

This repo contains:

- [poker](./poker): A **Go package** to evaluate Poker hands - Texas Hold'em
- [server](./server): A **Go package** to expose **Go package poker** as an API
- [keygen-go](./keygen-go): A **Go script** to search for keys used in package **poker**
- [keygen-c](./keygen-c): A **C script** to search for keys used in package poker - faster than **keygen-go**
- [Dockerfile](./Dockerfile) and commands to build and run server from within a container - image size 6.5 MB !

## Poker Go package

This [poker](./poker) Go package uses an algorithm introduced by [SpecialK](https://github.com/kennethshackleton/SKPokerEval).  
The underlying idea is to pre-calculate everything so that evaluating a hand is at most a couple of table lookups. Obviously this is very fast.

Cards are identified by their faces (2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K, A) only for _non flush hands_. So any combination of (7 among 13) should yield a different index. This could be done by mapping each face to one of the first 13 prime numbers, by example. But then the max index would be 2x3x5x7x11x..x41 ~ 3e14 which is way to large to fit in memory. So another mechanism must be found. SpecialK replaces multiplication by addition and looks for keys, such that no two distinct combinations (with max 4 identical keys) give the same sum, by brute force. Fortunately these keys (the "face keys") exist and can be found in a few hours. See [keygen-c](./keygen-c).

But first to determine whether a hand is a flush, apply the same mechanism to a much smaller problem: Each suit (spades, hearts, diamonds, clubs) is mapped to a key so that no two distinct combinations of 4 suits (with max 4 identical keys) add up to the same index. These keys (the "suit keys") are small and found instantly.

If the hand is not a flush, the keys described above are used. If it is, then again apply the same mechanism to a small problem: Each face is mapped to a key so that no two distinct combinations of 7 keys (all distinct) add up to the same index. These keys (the "flush keys") are small and found in a few seconds.

Besides it turns out these keys are small enough so that the face keys and the suit keys together fit into 32-bit integers.  
This is a surprising coincidence !

<img src="./img/roadrunner.jpeg" width="450" />

**SpecialK's algo**: _Speed and elegance combined... and luck !_

### Test Poker package

```bash
# test
# from /poker-go
go test ./poker -v

# benchmark
# from /poker-go
go test ./poker -run=XXX -bench=. -v
```

The package is quite fast as it goes through all (7 among 52) = 133.8m cases in 4s seconds.

```bash
❯ go test ./poker -v
=== RUN   TestGetRankFive
--- PASS: TestGetRankFive (0.03s)
=== RUN   TestGetRankSeven
--- PASS: TestGetRankSeven (0.02s)
=== RUN   TestGetRank
--- PASS: TestGetRank (0.02s)
=== RUN   TestBuildFiveHandStats
--- PASS: TestBuildFiveHandStats (0.10s)
=== RUN   TestBuildSevenHandStats
--- PASS: TestBuildSevenHandStats (3.83s)
PASS
ok      poker/poker     4.000s
```

## Build API server

```bash
# from /poker-go
# build
go build -o ../bin/poker

# from repo root
# run
./bin/poker
# folder certs must be in running dir with files:
#   ./certs/tls.crt
#   ./certs/tls.key
```

## Docker Build

```bash
# from repo /

# docker build
docker build -t go-poker-api:local .

# run
docker run -p 5000:5000 go-poker-api:local
```

## Request API

Test curls:

```bash
curl https://localhost:5000/healthz
# "OK"

curl https://localhost:5000/config
# {"FACE":["2","3","4","5","6","7","8","9","T","J","Q","K","A"],"SUIT":["C","D","H","S"],"CARD_NO":{"2C":0,"2D":1,"2H":2,"2S":3,"3C":4,"3D":5,"3H":6,"3S":7,"4C":8,"4D":9,"4H":10,"4S":11,"5C":12,"5D":13,"5H":14,"5S":15,"6C":16,"6D":17,"6H":18,"6S":19,"7C":20,"7D":21,"7H":22,"7S":23,"8C":24,"8D":25,"8H":26,"8S":27,"9C":28,"9D":29,"9H":30,"9S":31,"AC":48,"AD":49,"AH":50,"AS":51,"JC":36,"JD":37,"JH":38,"JS":39,"KC":44,"KD":45,"KH":46,"KS":47,"QC":40,"QD":41,"QH":42,"QS":43,"TC":32,"TD":33,"TH":34,"TS":35},"CARD_SY":{"0":"2C","1":"2D","10":"4H","11":"4S","12":"5C","13":"5D","14":"5H","15":"5S","16":"6C","17":"6D","18":"6H","19":"6S","2":"2H","20":"7C","21":"7D","22":"7H","23":"7S","24":"8C","25":"8D","26":"8H","27":"8S","28":"9C","29":"9D","3":"2S","30":"9H","31":"9S","32":"TC","33":"TD","34":"TH","35":"TS","36":"JC","37":"JD","38":"JH","39":"JS","4":"3C","40":"QC","41":"QD","42":"QH","43":"QS","44":"KC","45":"KD","46":"KH","47":"KS","48":"AC","49":"AD","5":"3D","50":"AH","51":"AS","6":"3H","7":"3S","8":"4C","9":"4D"}}

curl https://localhost:5000/stats-five
# {"flush":{"NbHand":1277,"MinRank":5863,"MaxRank":7139,"NbOccur":5108},"four-of-a-kind":{"NbHand":156,"MinRank":7296,"MaxRank":7451,"NbOccur":624},"full-house":{"NbHand":156,"MinRank":7140,"MaxRank":7295,"NbOccur":3744},"high-card":{"NbHand":1277,"MinRank":0,"MaxRank":1276,"NbOccur":1302540},"one-pair":{"NbHand":2860,"MinRank":1277,"MaxRank":4136,"NbOccur":1098240},"straight":{"NbHand":10,"MinRank":5853,"MaxRank":5862,"NbOccur":10200},"straight-flush":{"NbHand":10,"MinRank":7452,"MaxRank":7461,"NbOccur":40},"three-of-a-kind":{"NbHand":858,"MinRank":4995,"MaxRank":5852,"NbOccur":54912},"two-pairs":{"NbHand":858,"MinRank":4137,"MaxRank":4994,"NbOccur":123552}}

curl https://localhost:5000/stats-seven
# response
# {"flush":{"NbHand":1277,"MinRank":5863,"MaxRank":7139,"NbOccur":4047644},"four-of-a-kind":{"NbHand":156,"MinRank":7296,"MaxRank":7451,"NbOccur":224848},"full-house":{"NbHand":156,"MinRank":7140,"MaxRank":7295,"NbOccur":3473184},"high-card":{"NbHand":407,"MinRank":48,"MaxRank":1276,"NbOccur":23294460},"one-pair":{"NbHand":1470,"MinRank":1295,"MaxRank":4136,"NbOccur":58627800},"straight":{"NbHand":10,"MinRank":5853,"MaxRank":5862,"NbOccur":6180020},"straight-flush":{"NbHand":10,"MinRank":7452,"MaxRank":7461,"NbOccur":41584},"three-of-a-kind":{"NbHand":575,"MinRank":5003,"MaxRank":5852,"NbOccur":6461620},"two-pairs":{"NbHand":763,"MinRank":4140,"MaxRank":4994,"NbOccur":31433400}}

curl -X POST  -d '{"cards":[8,29,4,11,32]]}' https://localhost:5000/rank-five
# [{"Win":0.64821345,"Tie":0.045598652},{"Win":0.26058924,"Tie":0.045598652}]

curl -X POST  -d '{"players":[[8,29], [4,11]],"table":[]}' https://localhost:5000/calc
# [{"Win":0.64821345,"Tie":0.045598652},{"Win":0.26058924,"Tie":0.045598652}]

curl -X POST -d '{"players":[8,29],"table":[], "nb_player": 4, "nb_game": 100000}' https://localhost:5000/calc-mc
# {"Win":0.15905,"Tie":0}
```
