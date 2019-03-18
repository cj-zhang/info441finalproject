package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	tournamentAddr := os.Getenv("TOURNAMENTADDR")
	if len(tournamentAddr) == 0 {
		tournamentAddr = ":80"
	}
	// dsn := os.Getenv("DSN")
	// sqlDb, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// defer sqlDb.Close()
	// if err := sqlDb.Ping(); err != nil {
	// 	fmt.Printf("error pinging database: %v\n", err)
	// } else {
	// 	fmt.Printf("successfully connected!\n")
	// }
	// sqlStore := users.NewMySQLStore(sqlDb)
	// ctx := &TournamentContext{
	// 	UserStore:     sqlStore,
	// 	RabbitChannel: ch,
	// 	QueueName:     "notify",
	// }
	mux := http.NewServeMux()
	// mux.HandleFunc("/smashqq/tournaments", ctx.TourneyHandler)
	// mux.HandleFunc("/smashqq/tournaments/", ctx.TourneyHandler)
	// mux.HandleFunc("/smashqq/tournaments/{tournamentID}/players", ctx.PlayerHandler)
	// mux.HandleFunc("/smashqq/tournaments/{tournamentID}/organizers", ctx.OrganizerHandler)
	// mux.HandleFunc("/smashqq/tournaments/{tournamentID}/games", ctx.GamesHandler)
	// mux.HandleFunc("/smashqq/tournaments/{tournamentID}/standings", ctx.StandingsHandler)
	log.Printf("server is listening at %s...", tournamentAddr)
	log.Fatal(http.ListenAndServe(tournamentAddr, mux))
}
