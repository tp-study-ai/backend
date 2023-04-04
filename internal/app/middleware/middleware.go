package middleware

import "github.com/tp-study-ai/backend/tools/authManager"

type CommonMiddleware struct {
	AuthManager authManager.AuthManager
}

func NewCommonMiddleware(authManager authManager.AuthManager) CommonMiddleware {
	return CommonMiddleware{
		AuthManager: authManager,
	}
}
