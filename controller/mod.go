package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMods retrieves all available mods.
func GetMods(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// GetMod retrieves a specific mod.
func GetMod(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// DeleteMod removes a specific mod.
func DeleteMod(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PatchMod updates an existing mod.
func PatchMod(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PostMod creates a new mod.
func PostMod(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}
