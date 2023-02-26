package view

import "github.com/rivo/tview"

type PagesView struct {
	app   *tview.Application
	pages *tview.Pages
}

func NewPagesView() *PagesView {
	app := tview.NewApplication()
	pages := tview.NewPages()
	app.SetRoot(pages, true)
	return &PagesView{app: app, pages: pages}
}

func (v *PagesView) ResultPage(resultMsg string) {
	resultView := tview.NewTextView().SetText(resultMsg)
	v.pages.AddPage("Result", resultView, true, false)
	v.pages.SwitchToPage("Result")
}
