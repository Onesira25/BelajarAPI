package handler

type ToDoRequest struct {
	TaskName    string `json:"taskname" form:"taskname" validate:"required"`
	DueDate     string `json:"duedate" form:"duedate" validate:"required"`
	Description string `json:"desc" form:"desc"`
}

type ToDoUpdate struct {
	TaskName    string `json:"taskname" form:"taskname" validate:"required"`
	DueDate     string `json:"duedate" form:"duedate" validate:"required"`
	Description string `json:"desc" form:"desc"`
}
