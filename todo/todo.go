package todo

import (
	"cmp"
	"encoding/json"
	"os"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
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
	var items []Item
	dat, err := os.ReadFile(filename)
	if err != nil {
		return items, err
	}
	err = json.Unmarshal(dat, &items)
	if err != nil {
		return items, err
	}
	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyP() string {
	switch i.Priority {
	case 1:
		return "(1)"
	case 3:
		return "(3)"
	default:
		return "(2)"
	}
}

func (i *Item) PrettyDone() string {
	switch i.Done {
	case false:
		return "done(false)"
	default:
		return "done(true)"
	}
}

func (i *Item) Lable() string {
	return strconv.Itoa(i.position) + "."
}

func SortItems(a, b Item) int {
	if a.Done != b.Done {
		return -1
	}
	var c = cmp.Compare(a.Priority, b.Priority)
	if c == 0 {
		return cmp.Compare(a.position, b.position)
	}
	return c
}
