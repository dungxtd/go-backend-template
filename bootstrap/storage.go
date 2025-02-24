package bootstrap

import (
	"github.com/sportgo-app/sportgo-go/storage"
)

func NewStorage(env *Env) storage.MinioClient {
	return storage.NewMinioRestClient(env.MinioEndpoint, env.MinioAccessKeyID, env.MinioSecretKey, env.MinioUseSSL)
}
