package password_test

import (
	"testing"

	"rest-on-grpc-gateway/pkg/password"

	"github.com/stretchr/testify/require"
)

var pass = "pass"

func TestManager(t *testing.T) {
	t.Parallel()

	passwords := password.New()
	assert := require.New(t)
	hashPass, err := passwords.Hashing(pass)
	assert.NoError(err)
	compare := passwords.Compare(hashPass, []byte(pass))
	assert.Equal(true, compare)
}
