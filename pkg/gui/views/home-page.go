package views

import (
	"fmt"

	"testing/internal/entity"

	"github.com/lxn/walk"
	dcl "github.com/lxn/walk/declarative"
)

// const gopher = `iVBORw0KGgoAAAANSUhEUgAAAEsAAAA8CAAAAAALAhhPAAAFfUlEQVRYw62XeWwUVRzHf2+OPbo9d7tsWyiyaZti6eWGAhISoIGKECEKCAiJJkYTiUgTMYSIosYYBBIUIxoSPIINEBDi2VhwkQrVsj1ESgu9doHWdrul7ba73WNm3vOPtsseM9MdwvvrzTs+8/t95ze/33sI5BqiabU6m9En8oNjduLnAEDLUsQXFF8tQ5oxK3vmnNmDSMtrncks9Hhtt/qeWZapHb1ha3UqYSWVl2ZmpWgaXMXGohQAvmeop3bjTRtv6SgaK/Pb9/bFzUrYslbFAmHPp+3WhAYdr+7GN/YnpN46Opv55VDsJkoEpMrY/vO2BIYQ6LLvm0ThY3MzDzzeSJeeWNyTkgnIE5ePKsvKlcg/0T9QMzXalwXMlj54z4c0rh/mzEfr+FgWEz2w6uk8dkzFAgcARAgNp1ZYef8bH2AgvuStbc2/i6CiWGj98y2tw2l4FAXKkQBIf+exyRnteY83LfEwDQAYCoK+P6bxkZm/0966LxcAAILHB56kgD95PPxltuYcMtFTWw/FKkY/6Opf3GGd9ZF+Qp6mzJxzuRSractOmJrH1u8XTvWFHINNkLQLMR+XHXvfPPHw967raE1xxwtA36IMRfkAAG29/7mLuQcb2WOnsJReZGfpiHsSBX81cvMKywYZHhX5hFPtOqPGWZCXnhWGAu6lX91ElKXSalcLXu3UaOXVay57ZSe5f6Gpx7J2MXAsi7EqSp09b/MirKSyJfnfEEgeDjl8FgDAfvewP03zZ+AJ0m9aFRM8eEHBDRKjfcreDXnZdQuAxXpT2NRJ7xl3UkLBhuVGU16gZiGOgZmrSbRdqkILuL/yYoSXHHkl9KXgqNu3PB8oRg0geC5vFmLjad6mUyTKLmF3OtraWDIfACyXqmephaDABawfpi6tqqBZytfQMqOz6S09iWXhktrRaB8Xz4Yi/8gyABDm5NVe6qq/3VzPrcjELWrebVuyY2T7ar4zQyybUCtsQ5Es1FGaZVrRVQwAgHGW2ZCRZshI5bGQi7HesyE972pOSeMM0dSktlzxRdrlqb3Osa6CCS8IJoQQQgBAbTAa5l5epO34rJszibJI8rxLfGzcp1dRosutGeb2VDNgqYrwTiPNsLxXiPi3dz7LiS1WBRBDBOnqEjyy3aQb+/bLiJzz9dIkscVBBLxMfSEac7kO4Fpkngi0ruNBeSOal+u8jgOuqPz12nryMLCniEjtOOOmpt+KEIqsEdocJjYXwrh9OZqWJQyPCTo67LNS/TdxLAv6R5ZNK9npEjbYdT33gRo4o5oTqR34R+OmaSzDBWsAIPhuRcgyoteNi9gF0KzNYWVItPf2TLoXEg+7isNC7uJkgo1iQWOfRSP9NR11RtbZZ3OMG/VhL6jvx+J1m87+RCfJChAtEBQkSBX2PnSiihc/Twh3j0h7qdYQAoRVsRGmq7HU2QRbaxVGa1D6nIOqaIWRjyRZpHMQKWKpZM5feA+lzC4ZFultV8S6T0mzQGhQohi5I8iw+CsqBSxhFMuwyLgSwbghGb0AiIKkSDmGZVmJSiKihsiyOAUs70UkywooYP0bii9GdH4sfr1UNysd3fUyLLMQN+rsmo3grHl9VNJHbbwxoa47Vw5gupIqrZcjPh9R4Nye3nRDk199V+aetmvVtDRE8/+cbgAAgMIWGb3UA0MGLE9SCbWX670TDy1y98c3D27eppUjsZ6fql3jcd5rUe7+ZIlLNQny3Rd+E5Tct3WVhTM5RBCEdiEK0b6B+/ca2gYU393nFj/n1AygRQxPIUA043M42u85+z2SnssKrPl8Mx76NL3E6eXc3be7OD+H4WHbJkKI8AU8irbITQjZ+0hQcPEgId/Fn/pl9crKH02+5o2b9T/eMx7pKoskYgAAAABJRU5ErkJggg==`

// func gopherPNG() io.Reader { return base64.NewDecoder(base64.StdEncoding, strings.NewReader(gopher)) }

type HomePage struct {
	*walk.Composite
	app entity.App
	// openPB    *walk.PushButton
	db        *walk.DataBinder
	browserCB *walk.ComboBox
	// homeData  *formHomePage
	export     *walk.ComboBox
	limit      *walk.NumberEdit
	useperiod  *walk.CheckBox
	utmhost    *walk.LineEdit
	utmport    *walk.LineEdit
	dbfilename *walk.LineEdit
	saveconf   *walk.PushButton
	debug      *walk.CheckBox
}

type formHomePage struct {
	Browser    string
	Pwd        string
	Output     string
	Export     string
	Limit      int
	UsePeriod  bool
	UtmHost    string
	UtmPort    string
	DbFilename string
	Debug      bool
}

var Hd *formHomePage

func newHomePage(parent walk.Container, a entity.App) (entity.Page, error) {
	defer a.GetRecovery().RecoverLog("views:[home-page.go]")
	p := new(HomePage)
	p.app = a
	Hd = new(formHomePage)

	Hd.Browser = a.GetConfiguration().Browser
	Hd.Pwd = a.GetPwd()
	Hd.Output = a.GetOutput()
	Hd.Export = a.GetExport()
	Hd.DbFilename = a.GetConfiguration().Database.DbName
	Hd.UtmHost = a.GetConfiguration().UtmHost
	Hd.UtmPort = a.GetConfiguration().UtmPort
	Hd.Debug = a.GetConfiguration().Debug
	if Hd.Browser == "" {
		Hd.Browser = "по умолчанию"
	}
	// p.saveconf.SetEnabled(false)

	// img, err := png.Decode(gopherPNG())
	// if err != nil {
	// 	core.App().ErrorLog().AnErr("guiService:[gui\\multipagemainwindow.go]", err).Send()
	// }
	// pngImg, err := walk.NewBitmapFromImage(img)
	// if err != nil {
	// 	core.App().ErrorLog().AnErr("guiService:[gui\\multipagemainwindow.go]", err).Send()
	// }
	if err := (dcl.Composite{
		AssignTo: &p.Composite,
		DataBinder: dcl.DataBinder{
			AssignTo:            &p.db,
			Name:                "Hd",
			DataSource:          Hd,
			ErrorPresenter:      dcl.ToolTipErrorPresenter{},
			OnDataSourceChanged: p.changeData,
		},
		Name: "homePage",
		// Background: dcl.SolidColorBrush{walk.RGB(255, 255, 255)},
		Layout: dcl.VBox{MarginsZero: false, SpacingZero: true, Margins: dcl.Margins{0, 0, 5, 0}},
		// Layout: dcl.Grid{
		// 	Columns: 2,
		// },
		// Layout:    dcl.Flow{MarginsZero: true, SpacingZero: true},
		Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
		Children: []dcl.Widget{
			dcl.Composite{
				Border: true,
				Layout: dcl.Grid{Columns: 2},
				Children: []dcl.Widget{
					dcl.Label{
						Text:          "Выберите браузер:",
						TextAlignment: dcl.AlignNear,
					},
					dcl.ComboBox{
						AssignTo:              &p.browserCB,
						Alignment:             dcl.AlignHNearVNear,
						Editable:              false,
						Value:                 dcl.Bind("Browser"),
						Model:                 []string{"", "Chrome", "Firefox", "Yandex", "MSEdge"},
						OnCurrentIndexChanged: p.changeIndexBrowser,
					},
					// dcl.HSpacer{},
					dcl.PushButton{
						// ColumnSpan: 2,
						Text:      "Открыть Веб Приложение",
						OnClicked: p.openWebApp,
					},
					dcl.HSpacer{},
					dcl.Label{
						Text: "Хост УТМ:",
					},
					dcl.LineEdit{
						AssignTo:      &p.utmhost,
						Text:          dcl.Bind("UtmHost"),
						OnTextChanged: p.enablePB,
					},
					dcl.Label{
						Text: "Порт УТМ:",
					},
					dcl.LineEdit{
						AssignTo:      &p.utmport,
						Text:          dcl.Bind("UtmPort"),
						OnTextChanged: p.enablePB,
					},
					dcl.Label{
						Text: "Файл БД:",
					},
					dcl.LineEdit{
						AssignTo:      &p.dbfilename,
						Text:          dcl.Bind("DbFilename"),
						OnTextChanged: p.enablePB,
					},
					dcl.Label{
						Text: "Отладка:",
					},
					dcl.CheckBox{
						AssignTo: &p.debug,
						Checked:  dcl.Bind("Debug"),
					},
					dcl.PushButton{
						AssignTo: &p.saveconf,
						// ColumnSpan: 2,
						Text:      "Обновить конфигурацию",
						OnClicked: p.saveConfig,
					},
					dcl.HSpacer{},
					// dcl.HSpacer{},
					// dcl.Label{
					// 	Text: "Экспорт формат:",
					// },
					// dcl.ComboBox{
					// 	AssignTo:              &p.export,
					// 	Editable:              false,
					// 	Value:                 dcl.Bind("Export"),
					// 	OnCurrentIndexChanged: p.changeExport,
					// 	Model: []string{
					// 		"csv",
					// 		"xlsx",
					// 	},
					// },
					// dcl.HSpacer{},
					// dcl.Label{
					// 	Text: "Лимит:",
					// },
					// dcl.HSpacer{},
					dcl.HSpacer{},
					// dcl.PushButton{
					// 	ColumnSpan: 2,
					// 	Text:       "Очистить БД Декларации",
					// 	OnClicked:  p.clearDeclaracia,
					// },
					// dcl.HSpacer{},
				},
			},
		},
	}).Create(dcl.NewBuilder(parent)); err != nil {
		return nil, fmt.Errorf("Home_Page.Create()%w", err)
	}
	// p.browserCB.SetText("firfox")
	// p.browserCB.TextChanged().Attach(p.changeBrowser)

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, fmt.Errorf("walk.InitWrapperWindow(p) %w", err)
	}

	return p, nil
}

func (p *HomePage) changeIndexBrowser() {
	txt := p.browserCB.Text()
	// txt := p.browser
	p.app.DebugLog().Msgf("changeIndexBrowser() %s", txt)
	if err := p.app.SetBrowser(txt); err != nil {
		p.app.ErrorLog().AnErr("SetExport", err).Send()
	}
}

func (p *HomePage) changeData() {
	p.app.DebugLog().Msgf("changeData() %+v", Hd)
}

func (p *HomePage) openWebApp() {
	uri := p.app.GetBaseUrl()
	p.app.Open(uri)
}

func (p *HomePage) openOutput() {
	p.app.OpenDir()
}

func (p *HomePage) Clear() {

}

func (p *HomePage) changeExport() {
	txt := p.export.Text()
	p.app.DebugLog().Str("Format", txt).Msg("Change Export")

	if err := p.app.SetExport(txt); err != nil {
		p.app.ErrorLog().AnErr("SetExport", err).Send()
	}
}

func (p *HomePage) Update() {

}

func (p *HomePage) enablePB() {
	p.saveconf.SetEnabled(true)
}

func (p *HomePage) saveConfig() {
	host := p.utmhost.Text()
	port := p.utmport.Text()
	dbname := p.dbfilename.Text()
	debug := p.debug.Checked()
	if err := p.app.GetConfig().Set("utmhost", host, true); err != nil {
		p.app.ErrorLog().AnErr("set utmhost in config", err).Send()
	}
	if err := p.app.GetConfig().Set("utmport", port, true); err != nil {
		p.app.ErrorLog().AnErr("set utmhost in config", err).Send()
	}
	if err := p.app.GetConfig().Set("database.dbname", dbname, true); err != nil {
		p.app.ErrorLog().AnErr("set database.dbname in config", err).Send()
	}
	if err := p.app.GetConfig().Set("debug", debug, true); err != nil {
		p.app.ErrorLog().AnErr("set debug in config", err).Send()
	}
	// if err := p.app.InitUtm(); err != nil {
	// 	p.app.ErrorLog().AnErr("(p *HomePage) saveConfig() p.app.InitUtm()", err).Send()
	// }
	// if err := p.app.InitDb(); err != nil {
	// 	p.app.ErrorLog().AnErr("(p *HomePage) saveConfig() p.app.InitDb()", err).Send()
	// }
	// if err := p.app.GetRepo().Start(); err != nil {
	// 	p.app.ErrorLog().AnErr("(p *HomePage) saveConfig() p.app.GetRepo().Start()", err).Send()
	// }
	p.app.Restart()
}
