package images

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type RepositoryAuth struct {
	client client.Client
}

func NewRepositoryAuth(k8sClient client.Client) *RepositoryAuth {
	return &RepositoryAuth{client: k8sClient}
}

func (r *RepositoryAuth) GetAuthFileFromSecret(ctx context.Context, namespace, name string) (string, error) {
	secret := v1.Secret{}
	err := r.client.Get(ctx, client.ObjectKey{Namespace: namespace, Name: name}, &secret)
	if err != nil {
		return "", err
	}
	autFileBase64, found := secret.Data[".dockerconfigjson"]
	if !found {
		return "", errors.New(fmt.Sprintf(".dockerconfigjson not found in %s/%s Secret", namespace, name))
	}
	authFileBytes := make([]byte, base64.StdEncoding.DecodedLen(len(autFileBase64)))
	_, err = base64.StdEncoding.Decode(authFileBytes, autFileBase64)
	if err != nil {
		return "", err
	}

	return string(authFileBytes), nil
}
