package main

import (
	"flag"
	"log"
	"net/http"
	"projeto-ra-api-go/pkg/api/controller"
	"projeto-ra-api-go/pkg/api/provider/mongo/dao"
	"projeto-ra-api-go/pkg/api/router"
	"projeto-ra-api-go/pkg/utl/config"
	"projeto-ra-api-go/pkg/utl/mg"

	"projeto-ra-api-go/pkg/api/service"
)

// @title Ra Challenge API
// @version 1.0
// @description Ra Challenge API
// @termsOfService http://swagger.io/terms/
// @contact.name Thiago Menezes
// @contact.email thg.mnzs@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	cfgPath := flag.String("p", "./cmd/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatal("error get config file")
	}

	client, database, err := mg.New(cfg.DB.PSN, cfg.DB.DB)
	if err != nil {
		log.Fatal("error get database")
	}

	dao := dao.NewMongoComplain(client, database)
	raEpi := service.NewReclameAquiExternalApi()
	service := service.NewComplain(dao, raEpi)
	handler := controller.NewComplainController(service)
	healthHandler := controller.NewHealthHandler(dao)

	routers := router.NewRARouter(handler, healthHandler)

	log.Println("Servidor esta rodando na porta "+ cfg.Server.Port)
	log.Fatal(http.ListenAndServe(cfg.Server.Port, routers))
}