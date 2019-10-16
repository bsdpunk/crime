package main

import (
//	"github.com/facebookincubator/ent"
"github.com/bsdpunk/crime/ent"
"github.com/bsdpunk/crime/ent/migrate"
"github.com/facebookincubator/ent/dialect/sql"
"log"
)



func main() {
	client, err := sql.Open("mysql", "root:hxAWYa2YbW@tcp://127.0.0.1/crime_development")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx, 
		migrate.WithDropIndex(false),
		migrate.WithDropColumn(false), 
	)

	if err := client.Schema.Create(ctx); err != nil {
		    log.Fatalf("failed creating schema resources: %v", err)
	    }
}
