package server

import (
	"net/http"
	"poker/poker"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	Players  poker.PlayerCards `json:"players" binding:"required"`
	Table    poker.TableCards  `json:"table" binding:"required"`
	NbPlayer int               `json:"nb_player" binding:"required"`
	NbGame   int               `json:"nb_game"`
}

// @Tags		Status
// @Summary	server status
// @Produce	plain
// @Router		/healthz [get]
// @Success	200
// @Failure	500
func healthzHandler(c *gin.Context) {
	c.JSON(200, "Ok")
}

// @Tags		Hand
// @Summary	evaluate 5-card hand
// @Accept		json
// @Produce	json
// @Param		cards	body	rankFiveBody	true	"5-card hands"
// @Router		/rank-five [post]
// @Success	200	{object}	[]int	"ranks"
// @Failure	401	"invalid rank-five input"
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

// @Tags		Hand
// @Summary	evaluate 7-card hand
// @Accept		json
// @Produce	json
// @Param		cards	body	rankSevenBody	true	"5-card hands"
// @Router		/rank-seven [post]
// @Success	200	{object}	[]int	"ranks"
// @Failure	401	"invalid rank-seven input"
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

// @Tags		Calculate
// @Summary	exhaustive calculator
// @Accept		json
// @Produce	json
// @Param		cards	body	calcBody	true	"game cards: table and players"
// @Router		/calc [post]
// @Success	200	{object}	[]poker.handEquity	"players hands equity"
// @Failure	401	"invalid calc input"
// @Failure	422	"nb_player must be between 2 and 10"
// @Failure	423	"len(table) must be 0, 3, 4, 5"
func calcHandler(c *gin.Context) {
	var body calcBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(401, gin.H{"error": "invalid calc input"})
		return
	}
	CustomLog("/calc", "body", body)

	P := len(body.Players)
	// fmt.Println("P", P)
	if P < 2 || P > 10 {
		c.JSON(422, gin.H{"error": "nb_player must be between 2 and 10"})
		return
	}
	T := len(body.Table)
	// fmt.Println("T", T)
	if T != 0 && T != 3 && T != 4 && T != 5 {
		c.JSON(423, gin.H{"error": "len(table) must be 0, 3, 4, 5"})
		return
	}

	var eqty = poker.CalcEquity(body.Players, body.Table)

	c.JSON(200, eqty)
}

// @Tags		Calculate
// @Summary	monte carlo calculator
// @Accept		json
// @Produce	json
// @Param		cards	body	calcMonteCarloBody	true	"game cards: table and players"
// @Router		/calc-mc [post]
// @Success	200	{object}	[]poker.handEquity	"players hands equity"
// @Failure	401	"invalid calc-mc input"
// @Failure	422	"nb_player must be between 2 and 10"
// @Failure	423	"len(table) must be 0, 3, 4, 5"
// @Failure	423	"nb_game must be set"
func calcMonteCarloHandler(c *gin.Context) {
	var body calcMonteCarloBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(421, gin.H{"error": "invalid calcMonteCarlo input"})
		return
	}

	P := body.NbPlayer
	if P < 2 || P > 10 {
		c.JSON(422, gin.H{"error": "nb_player must be between 2 and 10"})
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

	var eqty = poker.CalcEquityMonteCarlo(body.Players, body.Table, body.NbPlayer, body.NbGame)

	c.JSON(200, eqty)
}

// @Tags		Static
// @Summary	static data
// @Produce	json
// @Router		/config [get]
// @Success	200	{object}	configData	"poker API static data"
// @Failure	500
func configHandler(c *gin.Context) {
	var config = configData{
		FACE:    poker.FACE,
		SUIT:    poker.SUIT,
		CARD_NO: poker.CARD_NO,
		CARD_SY: poker.CARD_SY,
	}
	c.JSON(200, config)
}

// ------------------------------------
// BOGUS HANDLER NECESSARY FOR SWAGGER
// REAL HANDLER DEFINED IN Serve
// ------------------------------------
//
//	@Tags		Stats
//	@Summary	5-card hands stats
//	@Produce	json
//	@Router		/stats-five [get]
//	@Success	200	{object}	map[string]poker.HandTypeStatsStruct	"5-card hands stats"
//	@Failure	500
func statsFiveHandler(c *gin.Context) {
	c.JSON(200, "bogus")
}

// ------------------------------------
// BOGUS HANDLER NECESSARY FOR SWAGGER
// REAL HANDLER DEFINED IN Serve
// ------------------------------------
//
//	@Tags		Stats
//	@Summary	7-card hands stats
//	@Produce	json
//	@Router		/stats-seven [get]
//	@Success	200	{object}	map[string]poker.HandTypeStatsStruct	"7-card hands stats"
//	@Failure	500
func statsSevenHandler(c *gin.Context) {
	c.JSON(200, "bogus")
}

func swaggerRedirectHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
}

// type statsHandler struct {
// 	statsFive  map[string]poker.HandTypeStatsStruct
// 	statsSeven map[string]poker.HandTypeStatsStruct
// }

// func (this *statsHandler) GetStatsFive(c *gin.Context) {
// 	c.JSON(200, this.statsFive)
// }

// func (this *statsHandler) GetStatsSeven(c *gin.Context) {
// 	c.JSON(200, this.statsSeven)
// }

func Serve() {

	poker.Setup(false)

	statsFive := poker.BuildFiveHandStats(true)
	statsSeven := poker.BuildSevenHandStats(true)

	//	@Tags		Stats
	//	@Summary	5-card hands stats
	//	@Produce	json
	//	@Router		/stats-five [get]
	//	@Success	200	{object}	map[string]poker.HandTypeStatsStruct	"5-card hands stats"
	//	@Failure	500
	var statsFiveHandler2 = func(c *gin.Context) {
		c.JSON(200, statsFive)
	}

	//	@Tags		Stats
	//	@Summary	7-card hands stats
	//	@Produce	json
	//	@Router		/stats-seven [get]
	//	@Success	200	{object}	map[string]poker.HandTypeStatsStruct	"7-card hands stats"
	//	@Failure	500
	var statsSevenHandler2 = func(c *gin.Context) {
		c.JSON(200, statsSeven)
	}

	// var _statsHander = statsHandler{
	// 	statsFive:  poker.BuildFiveHandStats(true),
	// 	statsSeven: poker.BuildSevenHandStats(true),
	// }

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

	router.GET("/healthz", healthzHandler)

	router.GET("/config", configHandler)

	router.GET("/stats-five", statsFiveHandler2)
	router.GET("/stats-seven", statsSevenHandler2)
	// router.GET("/stats-five", _statsHander.GetStatsFive)
	// router.GET("/stats-seven", _statsHander.GetStatsSeven)

	router.POST("/rank-five", rankFiveHandler)
	router.POST("/rank-seven", rankSevenHandler)

	router.POST("/calc", calcHandler)
	router.POST("/calc-mc", calcMonteCarloHandler)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/docs", swaggerRedirectHandler)

	// router.Run("0.0.0.0:5000")
	certFile := "./certs/tls.crt"
	keyFile := "./certs/tls.key"
	router.RunTLS("0.0.0.0:5000", certFile, keyFile)
}
