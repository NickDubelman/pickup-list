package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly/v2"

	"github.com/NickDubelman/pickup-list/db"
)

func main() {
	// Establish connection to our db
	dsn := "root@tcp(localhost:3381)/pickup_list?parseTime=true"
	client, err := db.Open(dialect.MySQL, dsn)
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	ctx := context.Background()

	// We will accumulate a bulk insert
	bulk := []*db.NBAPlayerCreate{}

	c := colly.NewCollector(colly.AllowedDomains("www.basketball-reference.com"))

	c.OnHTML("ul.page_index li > a", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		c.OnHTML("th[data-stat='player'] a", func(e2 *colly.HTMLElement) {
			bulk = append(bulk, client.NBAPlayer.Create().SetName(e2.Text))
		})

		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Visit("https://www.basketball-reference.com/players/")

	err = client.NBAPlayer.
		CreateBulk(bulk...).
		OnConflict().
		Ignore(). // ignore unique constraint errors
		Exec(ctx)

	if err != nil {
		if _, ok := err.(*db.ConstraintError); !ok {
			panic(err)
		}
	}
}
