package main

import (
	"net/http"
)

func (app *application) postBagCreate(w http.ResponseWriter, r *http.Request) {}
func (app *application) deleteBag(w http.ResponseWriter, r *http.Request)     {}
func (app *application) getBagAll(w http.ResponseWriter, r *http.Request)     {}
func (app *application) getBagById(w http.ResponseWriter, r *http.Request)    {}
