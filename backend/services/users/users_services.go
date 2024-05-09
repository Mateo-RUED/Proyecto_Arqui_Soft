package users

import "backend/domain/users"

func Login(request users.LoginRequest) users.LoginResponse {

	return users.LoginResponse{

		Token: "abcdef123456",
	}

}
