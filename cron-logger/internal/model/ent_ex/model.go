package ent_ex

type ListBaseReq struct {
	PN int `form:"page_num" json:"page_num"`
	PS int `form:"page_size" json:"page_size"`
}

type ListBaseResp struct {
	PS    int   `json:"page_size"`
	PN    int   `json:"page_num"`
	Total int64 `json:"total"`
}

type ListBaseOptions struct {
	PN int
	PS int
}

func (o *ListBaseOptions) Completed() {
	if o.PN <= 0 {
		o.PN = 1
	}
	if o.PS <= 0 {
		o.PS = 20
	}
}

func (o *ListBaseOptions) OffSet() int {
	return (o.PN - 1) * o.PS
}

func (o *ListBaseOptions) Limit() int {
	return o.PS
}
