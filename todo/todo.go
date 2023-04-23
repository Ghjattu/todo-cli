package todo

import (
	"encoding/json"
	"os"
	"strconv"
)

type Item struct {
	Text      string
	Priority  int
	Done      bool
	Timestamp string
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if len(b) == 0 {
		return []Item{}, nil
	}
	if err != nil {
		return []Item{}, err
	}
	items := make([]Item, 0)
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	return items, nil
}

func (item *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		item.Priority = 1
	case 3:
		item.Priority = 3
	default:
		item.Priority = 2
	}
}

func (item *Item) PrettyP() string {
	if item.Priority == 1 {
		return "(1)"
	}
	if item.Priority == 3 {
		return "(3)"
	}
	return " "
}

func (item *Item) PrettyDone() string {
	if item.Done {
		return "[X]"
	}
	return "[ ]"
}

func (item *Item) Label(i int) string {
	return strconv.Itoa(i) + "."
}

type ByPri []Item

func (s ByPri) Len() int {
	return len(s)
}

func (s ByPri) Less(i, j int) bool {
	if s[i].Done == s[j].Done {
		if s[i].Priority == s[j].Priority {
			return s[i].Timestamp > s[j].Timestamp
		}
		return s[i].Priority < s[j].Priority
	}
	return !s[i].Done
}

func (s ByPri) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
