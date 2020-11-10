package main

import (
	"encoding/json"
	"fmt"
)

type Alarm struct {
	ID   int    `'json:"id"`
	Type string `'json:"type"`
}

func main() {
	m := Alarm{
		ID:   1,
		Type: "firing",
	}

	b, _ := json.Marshal(m)

	fmt.Println(string(b))

	var m_json Alarm
	err := json.Unmarshal(b, &m_json)

	if err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	Type := m_json.Type
	ID := m_json.ID
	fmt.Println(Type, ID)

}
