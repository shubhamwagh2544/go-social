package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/shubhamwagh2544/go-social/internal/store"
)

type createPostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	userId := 1 // TODO: get user id from auth token

	var payload createPostPayload
	err := readJSON(w, r, &payload)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	post := store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserId:  int64(userId),
		Tags:    payload.Tags,
	}

	ctx := r.Context()
	err = app.store.Posts.Create(ctx, &post)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	postId, err := strconv.Atoi(chi.URLParam(r, "postId"))
	// postId, err := strconv.ParseInt(chi.URLParam(r, "postId"), 10, 64)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	ctx := r.Context()
	post, err := app.store.Posts.GetById(ctx, int64(postId))
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundError(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	log.Printf("Post in PostHandler: %+v\n", post)

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
