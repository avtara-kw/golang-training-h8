package params

type CreateEmployee struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Departement string `json:"department"`
}
