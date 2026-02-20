package main

import (
	"net/http"

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
		writeJSONError(w, http.StatusInternalServerError, err.Error())
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
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
