package common

type Paging struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Total      int64  `json:"total"`
	FakeCursor string `json:"fake_cursor"`
	NextCursor string `json:"next_cursor"`
}

func (p *Paging) Pagination()  {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 50
	}
}
