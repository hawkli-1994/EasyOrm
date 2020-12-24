package EasyOrm

import "database/sql"

import "./log"
import "./session"

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	// 加载驱动
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{db}
	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	err := engine.db.Close()
	if err != nil {
		log.Error("failed to close database")
	}
	log.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}