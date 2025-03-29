package payload

import (
	"encoding/json"
	"fmt"
	"os"
)

type Payload map[string]interface{}

func (p *Payload) Load(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	if err := json.Unmarshal(b, p); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}
