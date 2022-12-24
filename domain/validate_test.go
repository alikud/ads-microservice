package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidationOfferModel(t *testing.T) {

	OkOffer := Offer{
		Title:       "Iphone 100 Pro max",
		Description: "Top for this budget!",
		PhotoUrl:    []string{"Abc", "qwe", "ajklsd"},
		Price:       2000.8,
	}

	require.NoError(t, OkOffer.Validate())

	//Not more 3 links
	offerManyPhotoUrl := Offer{
		Title:       "Iphone 10 pro max",
		Description: "some description",
		PhotoUrl:    []string{"description1", "description2", "description3", "description4"},
		Price:       100,
	}

	require.Error(t, offerManyPhotoUrl.Validate())

	offerShorTitle := Offer{
		Title:       "pc",
		Description: "2 gb 2 cpu best for study",
		PhotoUrl:    []string{},
		Price:       500,
	}

	require.Error(t, offerShorTitle.Validate())

	offerNoPhotos := Offer{
		Title:       "macbook",
		Description: "2gb 2cpu best for study",
		PhotoUrl:    nil,
		Price:       1233,
	}

	require.Error(t, offerNoPhotos.Validate())

}
