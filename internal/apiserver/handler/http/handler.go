// Copyright 2024 许铭杰 (1044011439@qq.com). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package http

import "miniblog/internal/apiserver/biz"

// Handler 处理博客模块的请求.
type Handler struct {
	biz biz.IBiz
}

// NewHandler 创建新的 Handler 实例.
func NewHandler(biz biz.IBiz) *Handler {
	return &Handler{
		biz: biz,
	}
}
