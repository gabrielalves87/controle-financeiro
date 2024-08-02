package db

import (
	"context"
	"testing"

	//"time"
	//"github.com/gabrielalves87/controle-financeiro/utils"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
)

const (
	
)

func createRandomUser(t *testing.T) User {
	ctx := context.Background()
	arg := CreateUserParams{
		Username: fake.FullName() ,
		Password: fake.SimplePassword(),
		Email: fake.EmailAddress(),
	}
	user, err := testQueries.CreateUser(ctx ,arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)
	require.NotEmpty(t, user.CreatedAt)
	return user
}


func TestCreatUser(t *testing.T) {
	createRandomUser(t)
}


func TestGetUser(t *testing.T) {
	ctx := context.Background()
	user1 := createRandomUser(t)
	user2,err := testQueries.GetUser(ctx, user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.NotEmpty(t, user2.CreatedAt)

}

func TestGetbyID(t *testing.T) {
	ctx := context.Background()
	user1 := createRandomUser(t)
	user2,err := testQueries.GetUserById(ctx, user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.NotEmpty(t, user2.CreatedAt)

	}
