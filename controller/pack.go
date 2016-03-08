package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPacks retrieves all available packs.
func GetPacks(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// GetPack retrieves a specific pack.
func GetPack(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// DeletePack removes a specific pack.
func DeletePack(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PatchPack updates an existing pack.
func PatchPack(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PostPack creates a new pack.
func PostPack(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}
