package view

import (
	"context"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"os/signal"
	"syscall"
)

// PagesView - view with multiple pages.
type PagesView struct {
	app   *tview.Application
	pages *tview.Pages
}

// NewPagesView creates new PagesView object.
func NewPagesView() *PagesView {
	app := tview.NewApplication()
	pages := tview.NewPages()
	app.SetRoot(pages, true)
	return &PagesView{app: app, pages: pages}
}

// ResultPage shows page with description of certain operation result.
func (v *PagesView) ResultPage(resultMsg string) {
	resultView := tview.NewTextView().SetText(resultMsg)
	v.pages.AddPage("Result", resultView, true, false)
	v.pages.SwitchToPage("Result")
	err := v.app.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func (v *PagesView) GetContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	notifyCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-notifyCtx.Done()
		stop()
		cancel()
	}()
	return ctx
}
