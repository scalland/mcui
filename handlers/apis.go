package handlers

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"

	mcum "mcui/memcache"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-gonic/gin"
)

type SetRequest struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
	TTL   int32  `json:"ttl"` // in seconds
}

func SetKey(c *gin.Context) {
	var req SetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item := &memcache.Item{
		Key:        req.Key,
		Value:      []byte(req.Value),
		Expiration: req.TTL,
	}
	if err := mcum.Client.Set(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Key set successfully"})
}

func GetKey(c *gin.Context) {
	key := c.Param("key")
	item, err := mcum.Client.Get(key)
	if err == memcache.ErrCacheMiss {
		c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"key": key, "value": string(item.Value)})
}

func DeleteKey(c *gin.Context) {
	key := c.Param("key")
	if err := mcum.Client.Delete(key); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Key deleted"})
}

func Stats(c *gin.Context) {
	conn, err := net.Dial("tcp", "localhost:11211")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var lines []string

	fmt.Fprintf(conn, "stats\r\n")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "END" {
			break
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	data := gin.H{
		"Stats": lines,
	}

	c.JSON(http.StatusOK, data)
}
