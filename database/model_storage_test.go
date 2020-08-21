package database

import (
	"testing"

	"github.com/Oringik/nyan404-libs/models"
	"github.com/stretchr/testify/assert"
)

func TestModelStorage_Get(t *testing.T) {
	arrayToRecord := []interface{}{
		&models.UserCase{
			UserInfo: models.UserInfo{
				PictureID: 5,
			},
		},
		&models.UserCase{
			UserInfo: models.UserInfo{
				PictureID: 2,
			},
		},
		&models.UserCase{
			UserInfo: models.UserInfo{
				PictureID: 11,
			},
		},
		&models.UserCase{
			UserInfo: models.UserInfo{
				PictureID: 114,
			},
		},
	}

	db := NewModelStorage()

	db.storage.WriteArray(arrayToRecord)

	value, err := db.Model(&models.UserCase{UserInfo: models.UserInfo{PictureID: 114}}).Field("PictureID").Get()
	assert.NoError(t, err)
	assert.Equal(t, value.(*models.UserCase).PictureID, uint(114))

}
