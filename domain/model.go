package domain

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Offer struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	PhotoUrl    []string `json:"photoUrl"`
	Price       float32  `json:"price"`
}

//https://www.worldlink.com.cn/en/osdir/ozzo-validation.html
//https://segmentfault.com/a/1190000040214453/en
func (o *Offer) checkPhotoUrl(value interface{}) error {
	photoUrl, ok := value.([]string)
	if !ok {
		return errors.New("photo url value must be a array of strings")
	}
	if len(photoUrl) > 3 {
		return errors.New("not more then 3 photos")
	}
	if len(photoUrl) == 0 {
		return errors.New("offer must have 1 or more photo")
	}

	return nil
}

func (o *Offer) Validate() error {
	return validation.ValidateStruct(o,
		validation.Field(&o.Title, validation.Required, validation.Length(3, 200)),
		validation.Field(&o.Description, validation.Required, validation.Length(8, 1000)),
		validation.Field(&o.PhotoUrl, validation.Required, validation.By(o.checkPhotoUrl)),
		validation.Field(&o.Price, validation.Min(1.0)),
	)
}
