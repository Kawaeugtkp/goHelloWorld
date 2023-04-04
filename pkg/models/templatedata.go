package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	Intmap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{} // interfaceはタイプがはっきりしないものとして
	// 使えるみたいなことを言っています
	CSRToken string
	Flash string
	Warning string
	Error string
}