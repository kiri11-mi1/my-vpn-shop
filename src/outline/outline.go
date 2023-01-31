package outline

type AccessKey struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Port      string `json:"port"`
	Method    string `json:"method"`
	AccessUrl string `json:"accessUrl"`
}

type AccessKeys []AccessKey

type KeysResponse struct {
	Keys AccessKeys `json:"accessKeys"`
}
