package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/TechBowl-japan/go-stations/db"
	// TechBowlのコミットを取り込んだら削除されていた
	// "github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/router"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}
}

func realMain() error {
	// config values
	const (
		defaultPort   = ":8080"
		defaultDBPath = ".sqlite3/todo.db"
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultDBPath
	}

	// set time zone
	var err error
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	// set up sqlite3
	todoDB, err := db.NewDB(dbPath)
	if err != nil {
		return err
	}
	defer todoDB.Close()

	// NOTE: 新しいエンドポイントの登録はrouter.NewRouterの内部で行うようにする
	mux := router.NewRouter(todoDB)

	// TechBowlのコミットを取り込んだら削除されていた
	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(w, "Hello, world!\n")
	// }
	// mux.HandleFunc("/", helloHandler)
	// mux.HandleFunc("/healthz", handler.HealthzHandler)
	// var h handler.TODOHandler
	// mux.HandleFunc("/todos", h.CreateTODO)
	// log.Fatal(http.ListenAndServe(defaultPort, mux))

	// 以降TechBowlで追加された行
	// TODO: サーバーをlistenする

	return nil
}
