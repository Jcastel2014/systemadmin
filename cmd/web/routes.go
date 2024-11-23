package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *appDependencies) routes() http.Handler {

	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/submit", a.Submit)

	router.HandlerFunc(http.MethodGet, "/login", a.Login)
	router.HandlerFunc(http.MethodPost, "/login", a.Login)
	router.HandlerFunc(http.MethodGet, "/register", a.Register)
	router.HandlerFunc(http.MethodPost, "/register", a.Register)
	router.HandlerFunc(http.MethodGet, "/", a.Authenticate(a.SimpleServer)) 
	router.HandlerFunc(http.MethodGet, "/logout", a.Logout)

	return router

}
