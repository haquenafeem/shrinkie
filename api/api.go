package api

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/haquenafeem/shrinkie/internal/random"
	"github.com/haquenafeem/shrinkie/model"
	"github.com/haquenafeem/shrinkie/repository"
)

type Api struct {
	repo   *repository.Repository
	router *gin.Engine
}

func (api *Api) isValidUrl(str string) bool {
	url, err := url.ParseRequestURI(str)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	address := net.ParseIP(url.Host)

	if address == nil {
		return strings.Contains(url.Host, ".")
	}

	return true
}

func (api *Api) setupRoutes() {
	api.router.Static("/assets", "./assets")
	api.router.LoadHTMLGlob("templates/*")

	api.router.POST("/shrink", api.shrinkHandler)
	api.router.GET("/", api.indexHandler)
	api.router.GET("/:id", api.redirect)
	api.router.GET("/list", api.listAll)
}

func (api *Api) indexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func (api *Api) redirect(ctx *gin.Context) {
	hexValue := ctx.Param("id")
	urlModel, err := api.repo.GetUrl(hexValue)
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{"err": err.Error()})
		return
	}

	if urlModel.RedirectTo == "" {
		ctx.HTML(http.StatusOK, "error.html", gin.H{"err": "not found"})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, urlModel.RedirectTo)
}

func (api *Api) shrinkHandler(ctx *gin.Context) {
	var req model.ShrinkRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &model.ShrinkResponse{
			IsSuccess: false,
			Err:       "json binding error",
		})

		return
	}

	isValid := api.isValidUrl(req.Url)
	if !isValid {
		ctx.JSON(http.StatusBadRequest, &model.ShrinkResponse{
			IsSuccess: false,
			Err:       "invalid url",
		})

		return
	}

	randomString := random.RandomStringDefualt()
	url := &model.URL{
		RedirectTo:   req.Url,
		RandomString: randomString,
	}

	if err := api.repo.CreateURL(url); err != nil {
		ctx.JSON(http.StatusInternalServerError, &model.ShrinkResponse{
			IsSuccess: false,
			Err:       "failed to create entry",
		})

		return
	}

	ctx.JSON(http.StatusOK, &model.ShrinkResponse{
		IsSuccess:  true,
		Err:        "",
		RedirectTo: req.Url,
		HexValue:   randomString,
		FullURL:    ctx.Request.Host + "/" + randomString,
	})
}

func (api *Api) listAll(ctx *gin.Context) {
	urls, err := api.repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"urls": urls,
	})
}

func (api *Api) Run(port int) {
	if err := api.router.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}

func New(repo *repository.Repository, router *gin.Engine) *Api {
	api := &Api{
		repo:   repo,
		router: router,
	}

	api.setupRoutes()

	return api
}
