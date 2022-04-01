package errors

type Error struct {
	Code   string `json:"code"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func GenerateError(code string, status int, title string, detail string) Error {
	return Error{code, status, title, detail}
}
