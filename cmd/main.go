package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trichardsonjr78/terraform-operator-api/pkg/api"
	"github.com/trichardsonjr78/terraform-operator-api/pkg/common/db"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	api.RegisterRoutes(r, h)

	r.Run(port)

}
