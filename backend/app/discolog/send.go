package discolog

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func Send(data map[string]string, title string) {
	fields := make([]map[string]interface{}, 0)

	for key, value := range data {
		fields = append(fields, map[string]interface{}{
			"name":  key,
			"value": value,
		})
	}

	body := map[string]interface{}{
		"embeds": []map[string]interface{}{
			{
				"title":  title,
				"fields": fields,
			},
		},
	}

	bodyStr, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", os.Getenv("DISCORD_WEBHOOK_URL"), bytes.NewBuffer(bodyStr))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	http.DefaultClient.Do(req)
}
