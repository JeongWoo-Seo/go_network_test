package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type GetUserResponse struct {
	*ApiResponse
	Users []*User `json:"result"`
}

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

type CreateUserResponse struct {
	*ApiResponse
}

func (c *CreateUserRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.Age,
	}
}

type UpdateUserRequest struct {
	Name      string `json:"name" binding:"required"`
	UpdateAge int64  `json:"updateAge" binding:"required"`
}

type UpdateUserResponse struct {
	*ApiResponse
}

type DeleteUserRequest struct {
	Name string `json:"name" binding:"required"`
}

type DeleteUserResponse struct {
	*ApiResponse
}

func (c *DeleteUserRequest) ToUser() *User {
	return &User{
		Name: c.Name,
	}
}
