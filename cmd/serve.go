package cmd

import (
	"github.com/turjoc120/ecom/config"
	"github.com/turjoc120/ecom/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Start(cnf)
}
