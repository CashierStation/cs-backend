package global

import (
	"csbackend/authenticator"

	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

var Session *session.Store
var Authenticator *authenticator.Authenticator
var DB *gorm.DB
