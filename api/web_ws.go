package api

import (
	"github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/tasselsd/gorum/pkg/session"
)

func prepareWs() {
	upgrader := websocket.Upgrader{}
	app.Get("/ws", func(ctx iris.Context) {
		ws, err := upgrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
		if err != nil {
			write_e500_page(err, ctx)
			return
		}
		s := sessionFromCtx(ctx)
		if s == nil {
			ws.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "Unauthenticated"))
			ws.Close()
			return
		}
		ctx.Values().Get("session").(*session.Session).SetWs(ws)
		err = ws.WriteMessage(websocket.TextMessage, []byte(`{"event": "connected"}`))
		if err != nil {
			write_e500_page(err, ctx)
		}
	})
}
