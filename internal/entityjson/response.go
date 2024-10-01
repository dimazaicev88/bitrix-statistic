package entityjson

type Response struct {
	Result interface{}
	Error  string
	Total  int
}
