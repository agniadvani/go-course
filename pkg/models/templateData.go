package models

//The struct that holds the format of data to be passed into the templates.
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFtoken string
	Flash     string
	Warning   string
	Error     string
}
