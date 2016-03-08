package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBuilds retrieves all available builds.
func GetBuilds(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// GetBuild retrieves a specific build.
func GetBuild(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// DeleteBuild removes a specific build.
func DeleteBuild(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PatchBuild updates an existing build.
func PatchBuild(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PostBuild creates a new build.
func PostBuild(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}
