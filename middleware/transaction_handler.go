package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func TransactionHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start a transaction
		tx := db.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		// Set the transaction in the context
		c.Set("Tx", tx)

		// Call the next handler
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			// Rollback the transaction
			tx.Rollback()

			// Log the error
			for _, err := range c.Errors {
				log.Error().Err(err).Msg("Transaction rollback")
			}
			return
		}

		// Commit the transaction
		err := tx.Commit().Error
		if err != nil {
			log.Error().Err(err).Msg("Transaction commit error")
			return
		}
	}
}
