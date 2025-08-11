package network

import (
	"CRUD_SERVER/service"
	"CRUD_SERVER/types"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	userRouterInit     sync.Once //sync.Once 여러번 호출이 되어도 한 번만 호출이 되도록함
	userRouterInstance *userRouter
)

type userRouter struct {
	router      *Network
	userService *service.User
}

func NewUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}
	})

	router.registerGET("/", userRouterInstance.get)
	router.registerPOST("/", userRouterInstance.create)
	router.registerDELETE("/", userRouterInstance.delete)
	router.registerUPDATE("/", userRouterInstance.update)

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	var req types.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse(-1, "바인딩 오류 입니다", err.Error()),
		})
	} else if err := u.userService.Create(req.ToUser()); err != nil {
		u.router.failResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse(-1, "create 오류 입니다", err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse(1, "성공", nil),
		})
	}

}

func (u *userRouter) get(c *gin.Context) {
	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse(1, "성공 확인", nil),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	var req types.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse(-1, "바인딩 오류 입니다", err.Error()),
		})
	} else if err := u.userService.Update(req.Name, req.UpdateAge); err != nil {
		u.router.failResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse(-1, "update 오류 입니다", err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse(1, "성공", nil),
		})
	}

}

func (u *userRouter) delete(c *gin.Context) {
	var req types.DeleteUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse(-1, "바인딩 오류 입니다", err.Error()),
		})
	} else if err := u.userService.Delete(req.ToUser()); err != nil {
		u.router.failResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse(-1, "updata 오류 입니다", err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse(1, "성공", nil),
		})
	}

}
