package domain

// Table describe table in system
type Table struct {
	Model
	Name string `json:"name"`
	/* Status description
	0: empty
	1: in progress
	2: ordered
	*/
	Status int     `json:"status"`
	Orders []Order `json:"orders"`
}
