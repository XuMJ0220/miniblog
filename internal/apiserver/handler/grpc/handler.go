package grpc

import (
	"context"
	apiv1 "miniblog/pkg/api/apiserver/v1"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Handler 负责处理模块的请求.
type Handler struct {
	apiv1.UnimplementedMiniBlogServer // 提供默认实现
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Healthz(ctx context.Context, rq *emptypb.Empty) (*apiv1.HealthzResponse, error) {
	return &apiv1.HealthzResponse{
		Status:    apiv1.ServiceStatus_Healthy,
		Timestamp: time.Now().Format(time.DateTime),
	}, nil
}
