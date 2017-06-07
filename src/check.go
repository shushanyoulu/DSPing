package main

import (
	"database/sql"
	"time"

	"github.com/shushanyoulu/DSPing/src/cmdping"

	_ "github.com/mattn/go-sqlite3"
)

func startPing(db *sql.DB, c Config, t Target, res chan TargetStatus) {
	go runPingTest(db, c, t, res)
}

func runPingTest(db *sql.DB, c Config, t Target, res chan TargetStatus) {
	for {
		// log.Println("starting runPingTest ", t.Name)
		var status TargetStatus
		pingres := cmdping.Ping(t.Addr, t.Addr, t.Interval)
		lastcheck := time.Now().Format("2006-01-02 15:04")
		logtime := time.Now().Format("02 15:04")
		if pingres.LossPk == "" {
			time.Sleep(30 * 10e8)
			pingres.LossPk = "100"
		}
		status = TargetStatus{Target: &t, SendPk: pingres.SendPk, RevcPk: pingres.RevcPk, LossPk: pingres.LossPk, MaxDelay: pingres.MaxDelay, MinDelay: pingres.MinDelay, AvgDelay: pingres.AvgDelay, LastCheck: lastcheck}
		lock.Lock()
		// log.Println("INSERT ", lastcheck, t.Addr, t.Name, pingres.MaxDelay, pingres.AvgDelay, pingres.MinDelay, pingres.SendPk, pingres.RevcPk, pingres.LossPk)
		stmt, _ := db.Prepare("REPLACE INTO pinglog(logtime, ip, name, maxdelay, mindelay, avgdelay, sendpk, revcpk, losspk, lastcheck) values(?,?,?,?,?,?,?,?,?,?)")
		stmt.Exec(logtime, t.Addr, t.Name, pingres.MaxDelay, pingres.AvgDelay, pingres.MinDelay, pingres.SendPk, pingres.RevcPk, pingres.LossPk, lastcheck)
		lock.Unlock()
		stmt.Close()
		res <- status
		// log.Printf("runPingTest on %s finish!", t.Name)
		time.Sleep(5 * 10e7)
	}

}
