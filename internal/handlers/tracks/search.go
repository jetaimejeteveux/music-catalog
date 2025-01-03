package tracks

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) Search(c *gin.Context) {
	ctx := c.Request.Context()

	query := c.Query("query")
	pageSizeStr := c.Query("pageSize")
	pageIndexStr := c.Query("pageIndex")
	userId := c.GetUint("userId")

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		pageIndex = 1
	}

	response, err := h.service.Search(ctx, query, pageSize, pageIndex, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
