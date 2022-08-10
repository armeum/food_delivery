package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Translation(c *gin.Context) {

	locale := c.Request.URL.Query().Get(":locale")

	var lang language.Tag

	switch locale {
	case "en":
		lang = language.MustParse("en")
	case "ru":
		lang = language.MustParse("ru")
	case "uz":
		lang = language.MustParse("uz")
	default:
		http.NotFound(c.Writer, c.Request)
		return
	}

	p := message.NewPrinter(lang)

	p.Fprintf(c.Writer, "Welcome!\n")

	// new gin engine
	// gin.SetMode(gin.ReleaseMode)
	// r := gin.New()

	// apply i18n middleware
	// r.Use(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
	// 	RootPath:         "./translation",
	// 	AcceptLanguage:   []language.Tag{language.Russian, language.English, language.Uzbek},
	// 	DefaultLanguage:  language.English,
	// 	UnmarshalFunc:    json.Unmarshal,
	// 	FormatBundleFile: "json",
	// })))

	// r.Use(bundle)

	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, ginI18n.MustGetMessage("welcome"))
	// })

	// r.GET("/:name", func(c *gin.Context) {
	// 	c.String(http.StatusOK, ginI18n.MustGetMessage(&i18n.LocalizeConfig{
	// 		MessageID: "welcomeWithName",
	// 		TemplateData: map[string]string{
	// 			"name": c.Param("name"),
	// 		},
	// 	}))
	// })
}
