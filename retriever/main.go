package 

type Retriever interface{
	Get(url string) string
}