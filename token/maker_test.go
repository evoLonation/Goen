package token

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker("asq235ggi3aa53ys897g9a8w43asq235ggi3aa53ys897g9a8w43asq235ggi3aa53ys897g9a8w43")
	require.NoError(t, err)
	SubTestMaker(maker, t)

}
func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker("asq235ggi3asq235ggi3asq235ggi322")
	require.NoError(t, err)
	SubTestMaker(maker, t)
}

func SubTestMaker(maker Maker, t *testing.T) {
	username := "evoLonation"
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.Id)
	require.Equal(t, payload.Username, username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}
