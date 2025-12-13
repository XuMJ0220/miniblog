// Copyright 2024 许铭杰 (1044011439@qq.com). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"miniblog/internal/pkg/log"
	"time"

	"github.com/spf13/viper"
)

// Config 运行时配置结构体, 用于存储应用相关的配置
// 不用 viper.Get, 因为这种方式能更加清晰知道应用提供了哪些配置项
type Config struct {
	ServerMode string
	JWTKey     string
	Expiration time.Duration
}

// UnionServer 定义一个联合服务器, 根据 ServerMode 决定要启动的服务器类型
type UnionServer struct {
	cfg *Config
}

func (cfg *Config) NewUnionServer() (*UnionServer, error) {
	return &UnionServer{cfg: cfg}, nil
}

// Run 运行应用.
func (s *UnionServer) Run() error {
	log.Infow("ServerMode from ServerOptions", "jwt-key", s.cfg.JWTKey)
	log.Infow("ServerMode from Viper", "jwt-key", viper.GetString("jwt-key"))

	select {}
}
