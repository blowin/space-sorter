package main

import (
	"fmt"
	"os"
)

import (
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*walk.MainWindow
	edit      *walk.TextEdit
	choseButt PushButton
	startButt PushButton

	path string
}


func main() {
	mw := &MyMainWindow{}
	mw.Create()

	MainW := MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Space Sorter",
		MinSize:  Size{200, 150},
		Size:     Size{200, 200},
		Layout:   VBox{},
		Children: []Widget{
			TextEdit{
				AssignTo: &mw.edit,
				Enabled:  false,
			},
			mw.choseButt,
			mw.startButt,
		},
		Icon: "icon.ico",
	}

//////////////////////////RUN
	if _, err := MainW.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (mw *MyMainWindow) Create() {
	mw.startButt = PushButton{
		Text:      "Start",
		OnClicked: mw.startClick,
	}

	mw.choseButt = PushButton{
		Text:      "Chose Directory",
		OnClicked: mw.choseDirClick,
	}
}

func (mw *MyMainWindow) choseDirClick() {
	dlg := new(walk.FileDialog)
	dlg.FilePath = mw.path

	if ok, err := dlg.ShowBrowseFolder(mw); err != nil {
		return
	} else if !ok {
		return
	}

	mw.path = dlg.FilePath
	mw.edit.SetText(mw.path)
}

func (mw *MyMainWindow) startClick() {
	if strings.Compare(mw.edit.Text(), "") == 0 {
		walk.MsgBox(mw, "Warning", "Select folder", walk.MsgBoxIconWarning)
		return
	}
	waiter := Sort(mw.path)
	if waiter == nil {
		walk.MsgBox(mw, "Error", "Error of opening...", walk.MsgBoxIconError)
	} else {
		waiter.Wait()
		walk.MsgBox(mw, "Success", "Operation completed successfully", walk.MsgBoxIconInformation)
	}
}