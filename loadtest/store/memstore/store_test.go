package memstore

import (
	"testing"

	"github.com/mattermost/mattermost-server/model"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("NewMemStore", func(t *testing.T) {
		s := New()
		require.NotNil(t, s)
	})
}

func TestUser(t *testing.T) {
	s := New()

	t.Run("NilUser", func(t *testing.T) {
		u := s.User()
		require.Nil(t, u)
	})

	t.Run("SetUser", func(t *testing.T) {
		u := &model.User{}
		err := s.SetUser(u)
		require.Nil(t, err)
		uu := s.User()
		require.Equal(t, u, uu)
	})
}

func TestId(t *testing.T) {
	s := New()

	t.Run("EmptyId", func(t *testing.T) {
		id := s.Id()
		require.Empty(t, id)
	})

	t.Run("ExpectedId", func(t *testing.T) {
		expected := model.NewId()
		s.SetUser(&model.User{
			Id: expected,
		})
		id := s.Id()
		require.Equal(t, expected, id)
	})
}