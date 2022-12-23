package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSomething(t *testing.T) {

	OkOffer := Offer{
		Title:       "Iphone 100 Pro max",
		Description: "Top for this budget!",
		PhotoUrl:    []string{"Abc", "qwe", "ajklsd"},
		Price:       2000.8,
	}

	require.NoError(t, OkOffer.Validate())

	NotOkOffer := Offer{
		Title:       "Iphone 10 pro max",
		Description: "some description",
		PhotoUrl:    []string{"description1", "description2", "description3", "description4"},
		Price:       100,
	}
	require.Error(t, NotOkOffer.Validate())
}
