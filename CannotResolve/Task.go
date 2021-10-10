package CannotResolve

type Task struct {
	Header       string `json:"header"`
	Text         string `json:"text"`
	User         string `json:"user"`
	Priority     string `json:"priority"`
	CreationTime string `json:"creation_time"`
	EndTime      string `json:"end_time"`
	UserId       int    `json:"-"`
	PriorityId   int    `json:"-"`
}

func (t *Task) IdFromValue() {
	if t.User == "" || t.Priority == "" {
		panic("User or Priority is empty")
	}
	t.UserId = GetUsers()[t.User]
	t.PriorityId = GetPriority()[t.Priority]
}

func (t *Task) ValueFromID() {
	if t.UserId == 0 || t.PriorityId == 0 {
		panic("UserID or PriorityID is empty")
	}
	for key, val := range GetUsers() {
		if val == t.UserId {
			t.User = key
		}
	}
	for key, val := range GetPriority() {
		if val == t.PriorityId {
			t.Priority = key
		}
	}
}
