package repo

import (
	pbu "go-exam/user-service/genproto/user"
)

type UserStorageRedisI interface {
	Sign(user *pbu.UserDetail) (*pbu.ResponseMessage, error)
	Verification(req *pbu.VerificationUserRequest) (*pbu.User, error)
}
