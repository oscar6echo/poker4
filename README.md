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
go test -v

# benchmark
# from /poker-go
go test -run=XXX -bench=.
```

The package is quite fast as it goes through all (7 among 52) = 133.8m cases in 4s seconds.

```bash
❯ go test -v
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
curl http://localhost:5000/healthz
# "OK"

curl http://localhost:5000/config
# {"FACE":["2","3","4","5","6","7","8","9","T","J","Q","K","A"],"SUIT":["c","d","h","s"],"CARD_NO":{"2c":0,"2d":1,"2h":2,"2s":3,"3c":4,"3d":5,"3h":6,"3s":7,"4c":8,"4d":9,"4h":10,"4s":11,"5c":12,"5d":13,"5h":14,"5s":15,"6c":16,"6d":17,"6h":18,"6s":19,"7c":20,"7d":21,"7h":22,"7s":23,"8c":24,"8d":25,"8h":26,"8s":27,"9c":28,"9d":29,"9h":30,"9s":31,"Ac":48,"Ad":49,"Ah":50,"As":51,"Jc":36,"Jd":37,"Jh":38,"Js":39,"Kc":44,"Kd":45,"Kh":46,"Ks":47,"Qc":40,"Qd":41,"Qh":42,"Qs":43,"Tc":32,"Td":33,"Th":34,"Ts":35},"CARD_SY":{"0":"2c","1":"2d","10":"4h","11":"4s","12":"5c","13":"5d","14":"5h","15":"5s","16":"6c","17":"6d","18":"6h","19":"6s","2":"2h","20":"7c","21":"7d","22":"7h","23":"7s","24":"8c","25":"8d","26":"8h","27":"8s","28":"9c","29":"9d","3":"2s","30":"9h","31":"9s","32":"Tc","33":"Td","34":"Th","35":"Ts","36":"Jc","37":"Jd","38":"Jh","39":"Js","4":"3c","40":"Qc","41":"Qd","42":"Qh","43":"Qs","44":"Kc","45":"Kd","46":"Kh","47":"Ks","48":"Ac","49":"Ad","5":"3d","50":"Ah","51":"As","6":"3h","7":"3s","8":"4c","9":"4d"}}

curl http://localhost:5000/stats-five
# {"flush":{"NbHand":1277,"MinRank":5863,"MaxRank":7139,"NbOccur":5108},"four-of-a-kind":{"NbHand":156,"MinRank":7296,"MaxRank":7451,"NbOccur":624},"full-house":{"NbHand":156,"MinRank":7140,"MaxRank":7295,"NbOccur":3744},"high-card":{"NbHand":1277,"MinRank":0,"MaxRank":1276,"NbOccur":1302540},"one-pair":{"NbHand":2860,"MinRank":1277,"MaxRank":4136,"NbOccur":1098240},"straight":{"NbHand":10,"MinRank":5853,"MaxRank":5862,"NbOccur":10200},"straight-flush":{"NbHand":10,"MinRank":7452,"MaxRank":7461,"NbOccur":40},"three-of-a-kind":{"NbHand":858,"MinRank":4995,"MaxRank":5852,"NbOccur":54912},"two-pairs":{"NbHand":858,"MinRank":4137,"MaxRank":4994,"NbOccur":123552}}

curl http://localhost:5000/stats-seven
# response
# {"flush":{"NbHand":1277,"MinRank":5863,"MaxRank":7139,"NbOccur":4047644},"four-of-a-kind":{"NbHand":156,"MinRank":7296,"MaxRank":7451,"NbOccur":224848},"full-house":{"NbHand":156,"MinRank":7140,"MaxRank":7295,"NbOccur":3473184},"high-card":{"NbHand":407,"MinRank":48,"MaxRank":1276,"NbOccur":23294460},"one-pair":{"NbHand":1470,"MinRank":1295,"MaxRank":4136,"NbOccur":58627800},"straight":{"NbHand":10,"MinRank":5853,"MaxRank":5862,"NbOccur":6180020},"straight-flush":{"NbHand":10,"MinRank":7452,"MaxRank":7461,"NbOccur":41584},"three-of-a-kind":{"NbHand":575,"MinRank":5003,"MaxRank":5852,"NbOccur":6461620},"two-pairs":{"NbHand":763,"MinRank":4140,"MaxRank":4994,"NbOccur":31433400}}

curl -X POST  -d '{"players":[[8,29], [4,11]],"table":[]}' http://localhost:5000/calc
# [{"Win":0.64821345,"Tie":0.045598652},{"Win":0.26058924,"Tie":0.045598652}]

curl -X POST -d '{"players":[8,29],"table":[], "nb_player": 4, "nb_game": 100000}' http://localhost:5000/calc-mc
# {"Win":0.15905,"Tie":0}
```
