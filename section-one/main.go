package main

import "log"

func main() {
	db, err := NewDBClient()
	if err != nil {
		log.Fatalf("Db Error: %s\n", err)
		return
	}

	err = db.RunMigration()
	if err != nil {
		log.Fatalf("Migration failed: %s\n", err)
	}

	service := NewServer(db)
	log.Fatal(service.Start())
}
