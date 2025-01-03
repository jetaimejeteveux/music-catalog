package tracks

import (
	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/music-catalog/internal/models/trackactivites"
	"net/http"
)

func (h *Handler) UpsertTrackActivities(c *gin.Context) {
	ctx := c.Request.Context()

	var req trackactivites.TrackActivityReqest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userId := c.GetUint("userId")
	err := h.service.UpsertTrackActivities(ctx, userId, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}
