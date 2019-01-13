package judge

import "strings"

// Group represents a judge group
type Group struct {
	Judges []Judge
}

func (group *Group) String() string {
	strs := make([]string, len(group.Judges))
	for _, judge := range group.Judges {
		strs = append(strs, judge.String())
	}
	return strings.Join(strs, "\n")
}

// Judge represents a judge
type Judge struct {
	Name string
}

func (judge *Judge) String() string {
	return judge.Name
}

// Get a judge group
func Get() (*Group, error) {
	return &Group{}, nil
}
