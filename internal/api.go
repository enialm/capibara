package internal

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine, db *sql.DB) {
	r.POST("/event", func(c *gin.Context) {
		var req struct {
			Event string `json:"event" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid event"})
			return
		}
		if len(req.Event) == 0 || len(req.Event) > 128 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event name"})
			return
		}
		if err := InsertEvent(db, req.Event); err != nil {
			log.Printf("InsertEvent error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/stats", func(c *gin.Context) {
		var startPtr, endPtr *int64
		if startStr := c.Query("start"); startStr != "" {
			start, err := strconv.ParseInt(startStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start parameter value"})
				return
			}
			startPtr = &start
		}
		if endStr := c.Query("end"); endStr != "" {
			end, err := strconv.ParseInt(endStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end parameter value"})
				return
			}
			endPtr = &end
		}
		counts, err := GetEventCounts(db, startPtr, endPtr)
		if err != nil {
			log.Printf("GetEventCounts error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
			return
		}
		c.JSON(http.StatusOK, counts)
	})

	r.POST("/delete", func(c *gin.Context) {
		var req struct {
			Event string `json:"event" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid event"})
			return
		}
		n, err := DeleteEventsByName(db, req.Event)
		if err != nil {
			log.Printf("DeleteEventsByName error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": fmt.Sprintf("%d matching records deleted", n),
		})
	})

	r.POST("/truncate", func(c *gin.Context) {
		if err := TruncateEvents(db); err != nil {
			log.Printf("TruncateEvents error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "all records deleted"})
	})
}
