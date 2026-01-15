package user

import (
	"context"
	"miniblog/internal/apiserver/model"
	"miniblog/internal/apiserver/store"
	apiv1 "miniblog/pkg/api/apiserver/v1"
	"miniblog/pkg/store/where"

	"github.com/jinzhu/copier"
)

type UserBiz interface {
	Create(ctx context.Context, rq *apiv1.CreateUserRequest) (*apiv1.CreateUserResponse, error)
	Update(ctx context.Context, rq *apiv1.UpdateUserRequest) (*apiv1.UpdateUserResponse, error)
	Delete(ctx context.Context, rq *apiv1.DeleteUserRequest) (*apiv1.DeleteUserResponse, error)
	Get(ctx context.Context, rq *apiv1.GetUserRequest) (*apiv1.GetUserResponse, error)
	List(ctx context.Context, rq *apiv1.ListUserRequest) (*apiv1.ListUserResponse, error)

	UserExpansion
}

type UserExpansion interface {
}

type userBiz struct {
	store store.IStore
}

var _ UserBiz = (*userBiz)(nil)

func New(store store.IStore) *userBiz {
	return &userBiz{
		store: store,
	}
}

// Create 实现 UserBiz 接口中的 Create 方法.
func (b *userBiz) Create(ctx context.Context, rq *apiv1.CreateUserRequest) (*apiv1.CreateUserResponse, error) {
	var userM model.UserM
	_ = copier.Copy(&userM, rq)
	if err := b.store.User().Create(ctx, &userM); err != nil {
		return nil, err
	}

	return &apiv1.CreateUserResponse{UserID: userM.UserID}, nil
}

// Update 实现 UserBiz 接口中的 Update 方法.
func (b *userBiz) Update(ctx context.Context, rq *apiv1.UpdateUserRequest) (*apiv1.UpdateUserResponse, error) {
	userM, err := b.store.User().Get(ctx, where.T(ctx))
	if err != nil {
		return nil, err
	}

	if rq.Username != nil {
		userM.Username = *rq.Username
	}

	if rq.Nickname != nil {
		userM.Nickname = *rq.Nickname
	}

	if rq.Email != nil {
		userM.Email = *rq.Email
	}

	if rq.Phone != nil {
		userM.Phone = *rq.Phone
	}

	err = b.store.User().Update(ctx, userM)
	if err != nil {
		return nil, err
	}

	return &apiv1.UpdateUserResponse{}, nil
}

// Delete 实现 UserBiz 接口中的 Delete 方法.
func (b *userBiz) Delete(ctx context.Context, rq *apiv1.DeleteUserRequest) (*apiv1.DeleteUserResponse, error) {
	// 只有 `root` 用户可以删除用户，并且可以删除其他用户
	// 这里不用 where.T()，因为 where.T() 会查询用户自己，而不是查询 rq.UserID 指定的用户
	err := b.store.User().Delete(ctx, where.F("userID",rq.UserID))
	if err != nil {
		return nil, err
	}

	return &apiv1.DeleteUserResponse{}, nil
}

