package server

import (
	"poker/poker"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type configData struct {
	FACE    [13]string
	SUIT    [4]string
	CARD_NO map[string]int
	CARD_SY map[int]string
}

type rankFiveBody struct {
	Cards [][5]int `json:"cards" binding:"required"`
}

type rankSevenBody struct {
	Cards [][7]int `json:"cards" binding:"required"`
}

type calcBody struct {
	Players poker.PlayerCards `json:"players" binding:"required"`
	Table   poker.TableCards  `json:"table" binding:"required"`
}

type calcMonteCarloBody struct {
	Player   [2]int           `json:"players" binding:"required"`
	Table    poker.TableCards `json:"table" binding:"required"`
	NbPlayer int              `json:"nb_player" binding:"required"`
	NbGame   int              `json:"nb_game"`
}

func rankFiveHandler(c *gin.Context) {
	var body rankFiveBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(401, gin.H{"error": "invalid rankFive input"})
		return
	}
	CustomLog("/rank-five", "body", body)

	var ranks = make([]int, len(body.Cards))

	for i, cards := range body.Cards {
		r := poker.GetRankFive(cards)
		ranks[i] = r
	}

	c.JSON(200, ranks)
}

func rankSevenHandler(c *gin.Context) {
	var body rankSevenBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(401, gin.H{"error": "invalid rankSeven input"})
		return
	}
	CustomLog("/rank-seven", "body", body)

	var ranks = make([]int, len(body.Cards))

	for i, cards := range body.Cards {
		r := poker.GetRankSeven(cards)
		ranks[i] = r
	}

	c.JSON(200, ranks)
}

func calcHandler(c *gin.Context) {
	var body calcBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(401, gin.H{"error": "invalid calc input"})
		return
	}
	CustomLog("/calc", "body", body)

	var eqty = poker.CalcEquity(body.Players, body.Table)

	c.JSON(200, eqty)
}

func calcMonteCarloHandler(c *gin.Context) {
	var body calcMonteCarloBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(421, gin.H{"error": "invalid calcMonteCarlo input"})
		return
	}

	if body.NbPlayer < 1 || body.NbPlayer > 9 {
		c.JSON(422, gin.H{"error": "nb_player must be between 1 and 9"})
		return
	}
	T := len(body.Table)
	if T != 0 && T != 3 && T != 4 && T != 5 {
		c.JSON(423, gin.H{"error": "len(table) must be 0, 3, 4, 5"})
		return
	}

	if body.NbGame == 0 {
		body.NbGame = 1e5
		c.JSON(424, gin.H{"error": "set nb_game"})
		return
	}

	CustomLog("/calc-mc", "body", body)

	var eqty = poker.CalcEquityMonteCarlo(body.Player, body.Table, body.NbPlayer, body.NbGame)

	c.JSON(200, eqty)
}

func configHandler(c *gin.Context) {
	var config = configData{
		FACE:    poker.FACE,
		SUIT:    poker.SUIT,
		CARD_NO: poker.CARD_NO,
		CARD_SY: poker.CARD_SY,
	}
	c.JSON(200, config)
}

type statsHandler struct {
	statsFive  map[string]poker.HandTypeStatsStruct
	statsSeven map[string]poker.HandTypeStatsStruct
}

func (this *statsHandler) GetStatsFive(c *gin.Context) {
	c.JSON(200, this.statsFive)
}

func (this *statsHandler) GetStatsSeven(c *gin.Context) {
	c.JSON(200, this.statsSeven)
}

func Serve() {

	poker.Setup(false)

	var _statsHander = statsHandler{
		statsFive:  poker.BuildFiveHandStats(true),
		statsSeven: poker.BuildSevenHandStats(true),
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)

	// cors must be before endpoints
	corsConfig := cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	})
	router.Use(corsConfig)

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, "Ok")
	})

	router.GET("/config", configHandler)

	router.GET("/stats-five", _statsHander.GetStatsFive)
	router.GET("/stats-seven", _statsHander.GetStatsSeven)

	router.POST("/rank-five", rankFiveHandler)
	router.POST("/rank-seven", rankSevenHandler)

	router.POST("/calc", calcHandler)
	router.POST("/calc-mc", calcMonteCarloHandler)

	// router.Run("0.0.0.0:5000")
	certFile := "./certs/tls.crt"
	keyFile := "./certs/tls.key"
	router.RunTLS("0.0.0.0:5000", certFile, keyFile)
}
