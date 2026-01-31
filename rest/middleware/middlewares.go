package middleware

import "ecoommerce/config"

type Middlewares struct {
	cnf *config.Config
}

func NewMiddleWares(cnf *config.Config) *Middlewares {
	return &Middlewares{
		cnf: cnf,
	}
}
