package triggers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/atropos112/atrka/types"
	"github.com/gin-gonic/gin"
)

// WebhookHandler startrs a webhook server that listens on a given port and path and sends the received json payload as a message via producer
func WebhookHandler(port int, path string) func(ctx context.Context, ch chan<- types.Message) error {
	return func(_ context.Context, ch chan<- types.Message) error {
		// gin.SetMode(gin.ReleaseMode)
		r := gin.New()

		// r.Use(ginzap.Ginzap(l, time.RFC3339, true))
		r.POST(path, func(c *gin.Context) {
			var json map[string]interface{}

			if err := c.BindJSON(&json); err == nil {
				// fmt.Println(json)
				msg := types.JSONMessage{JSON: json}
				ch <- msg
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
			}
		})

		if err := r.Run(":" + strconv.Itoa(port)); err != nil {
			return err
		}

		return nil
	}
}

// MakeWebhookTrigger is a helpful function to create a webhook trigger using the webhook handler
func MakeWebhookTrigger[P types.Producer](producer P, port int, path string) types.Trigger[P] {
	return types.NewTrigger(WebhookHandler(port, path), producer)
}
