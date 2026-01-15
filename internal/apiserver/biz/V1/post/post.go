package post

import "miniblog/internal/apiserver/store"

type PostBiz interface {
	PostExpansion
}

type PostExpansion interface {
}

type postBiz struct {
	store store.IStore
}
var _ PostBiz = (*postBiz)(nil)

func New(store store.IStore) *postBiz {
	return &postBiz{
		store: store,
	}
}