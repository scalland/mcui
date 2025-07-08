package handlers

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	mcum "mcui/memcache"

	"github.com/bradfitz/gomemcache/memcache"
)

func RenderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func HandleSetHTML(c *gin.Context) {
	key := c.PostForm("key")
	value := c.PostForm("value")
	ttlStr := c.PostForm("ttl")

	var ttl int32 = 0
	if ttlStr != "" {
		if parsed, err := strconv.Atoi(ttlStr); err == nil {
			ttl = int32(parsed)
		}
	}

	err := mcum.Client.Set(&memcache.Item{
		Key:        key,
		Value:      []byte(value),
		Expiration: ttl,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "Set failed: %v", err)
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}

func HandleGetHTML(c *gin.Context) {
	key := c.Query("key")
	item, err := mcum.Client.Get(key)
	data := gin.H{}
	if err == nil {
		data["Value"] = string(item.Value)
	}
	c.HTML(http.StatusOK, "index.html", data)
}

func HandleDeleteHTML(c *gin.Context) {
	key := c.PostForm("key")
	mcum.Client.Delete(key)
	c.Redirect(http.StatusSeeOther, "/")
}

func HandleStatsHTML(c *gin.Context) {
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

	c.HTML(http.StatusOK, "index.html", data)
}
