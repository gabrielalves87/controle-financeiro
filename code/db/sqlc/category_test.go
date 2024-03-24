package db

import (
	"context"
	"testing"

	//"time"
	//"github.com/gabrielalves87/controle-financeiro/utils"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	ctx := context.Background()
	user := createRandomUser(t)
	arg := CreateCategoryParams{
		UserID:      user.ID,
		Title:       fake.Title(),
		Type:        "debit",
		Description: fake.Sentence(),
	}
	category, err := testQueries.CreateCategory(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.UserID, category.UserID)
	require.Equal(t, arg.Title, category.Title)
	require.Equal(t, arg.Type, category.Type)
	require.Equal(t, arg.Description, category.Description)
	require.NotEmpty(t, category.CreatedAt)
	return category
}

func TestCreatCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	ctx := context.Background()
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategory(ctx, category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Title, category2.Title)
	require.Equal(t, category1.Type, category2.Type)
	require.Equal(t, category1.Description, category2.Description)
	require.NotEmpty(t, category2.CreatedAt)

}

func TestDeleteCategory(t *testing.T) {
	ctx := context.Background()
	category := createRandomCategory(t)
	err := testQueries.DeleteCategory(ctx, category.ID)
	require.NoError(t, err)

}

func TestUpdateCategory(t *testing.T) {
	ctx := context.Background()
	category1 := createRandomCategory(t)
	arg := UpdateCategoriesParams{
		ID:          category1.ID,
		Title:       fake.Title(),
		Description: fake.Sentence(),
	}
	category2, err := testQueries.UpdateCategories(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, arg.ID, category1.ID)
	require.NotEqual(t, arg.Title, category1.Title)
	require.NotEqual(t, arg.Description, category1.Description)
	require.NotEmpty(t, category1.CreatedAt)
}

func TestListCategory(t *testing.T) {
	ctx := context.Background()
	// var lastcategory Category
	// for i := 0 ; i < 5 ; i ++ {
	lastcategory := createRandomCategory(t)
	// }

	arg := GetCategoriesParams{
		UserID:      lastcategory.UserID,
		Type:        lastcategory.Type,
		Title:       lastcategory.Title,
		Description: lastcategory.Description,
	}
	categorys, err := testQueries.GetCategories(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, categorys)
	for _, category := range categorys{ 
		require.Equal(t, lastcategory.ID, category.ID)
		require.Equal(t, lastcategory.UserId, category.UserId)
		require.Equal(t, lastcategory.Title, category.Title)
		require.Equal(t, lastcategory.Description, category.Description)
		require.NotEmpty(t, category.CreatedAt)
	}
}
