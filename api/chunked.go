package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ChunkedParams struct {
	Query string `json:"query"`
}

// POST
func ChunkedTest(c *gin.Context) {
	// qrystr := c.Query("query")
	// if len(qrystr) == 0 {
	// 	fmt.Println("query not found")
	// } else {
	// 	fmt.Println("query found", qrystr)
	// }

	var json ChunkedParams
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("binding query succ")
		fmt.Println(json.Query)
	}

	// _ = qrystr

	// body := c.Request.Body
	// value, err := ioutil.ReadAll(body)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(string(value))

	c.JSON(http.StatusOK, "Chunked Test End")
}

func ChunkedResJson(c *gin.Context) {
	w := c.Writer
	header := w.Header()

	header.Set("Transfer-Encoding", "chunked")
	header.Set("Content-Type", "application:json")
	w.WriteHeader(http.StatusOK)

	dd := gin.H{
		"chunk_res": "hello",
	}
	buf, _ := json.Marshal(dd)

	w.Write(buf)
	w.(http.Flusher).Flush()

	for i := 0; i < 10; i++ {
		dd = gin.H{
			"chunk_res": "hello-" + strconv.Itoa(i),
		}
		buf, _ = json.Marshal(dd)

		w.Write(buf)
		w.(http.Flusher).Flush()

		time.Sleep(time.Duration(1) * time.Second)
	}

	dd = gin.H{
		"chunk_res": "hello-end",
	}
	buf, _ = json.Marshal(dd)
	w.Write(buf)
	w.(http.Flusher).Flush()
}

func ChunkedResText(c *gin.Context) {
	w := c.Writer

	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "chunked start")
	w.Flush()

	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "chunked %d", i)
		w.Flush()

		time.Sleep(time.Duration(1) * time.Second)
	}

	fmt.Fprintf(w, "chunked end")
	w.Flush()
}
