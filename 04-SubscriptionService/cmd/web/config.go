package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type Config struct {
	Sessions *scs.SessionManager
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Wait     *sync.WaitGroup
}
