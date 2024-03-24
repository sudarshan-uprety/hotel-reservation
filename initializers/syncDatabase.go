package initializers

import (
	"github.com/sudarshan-uprety/hotel-reservation/db"
	"github.com/sudarshan-uprety/hotel-reservation/types"
)

func SyncDatabase() {
	db.DB.AutoMigrate(&types.User{})
}
