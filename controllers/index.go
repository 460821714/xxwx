package controllers

type IndexController struct {
	BaseControllers
}

func (index *IndexController) Get() {

	index.TplName = "index.html"

}
func (index *IndexController) post() {

}
