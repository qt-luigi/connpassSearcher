package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tenntenn/connpass"
)

const (
	count = 100
)

var (
	keyword string
	year    int
	month   int

	weekdays = []string{"日", "月", "火", "水", "木", "金", "土"}
)

func init() {
	flag.StringVar(&keyword, "k", "", "keyword for event (default empty)")
	flag.IntVar(&year, "y", 0, "year to hold event (default this year)")
	flag.IntVar(&month, "m", 0, "month to hold event (default this month)")
}

func main() {
	flag.Parse()

	if year == 0 || month == 0 {
		t := time.Now()
		if year == 0 {
			year, _ = strconv.Atoi(t.Format("2006"))
		}
		if month == 0 {
			month, _ = strconv.Atoi(t.Format("01"))
		}
	}

	param := []connpass.Param{
		connpass.YearMonth(year, time.Month(month)),
		connpass.Count(count),
		connpass.Order(connpass.OrderByDate),
	}
	if keyword != "" {
		param = append(param, connpass.Keyword(keyword))
	}
	params, err := connpass.SearchParam(param...)
	if err != nil {
		log.Fatal(err)
	}

	cli := connpass.NewClient()
	ctx := context.Background()
	r, err := cli.Search(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	if r.Available > 100 {
		fmt.Fprintln(os.Stderr, "results_available over 100")
		os.Exit(1)
	}

	events := make([]string, 0)

	for _, e := range r.Events {
		// date
		sm, _ := strconv.Atoi(e.StartedAt.Format("01"))
		sd, _ := strconv.Atoi(e.StartedAt.Format("02"))
		smd := fmt.Sprintf("%d/%d", sm, sd)
		em, _ := strconv.Atoi(e.EndedAt.Format("01"))
		ed, _ := strconv.Atoi(e.EndedAt.Format("02"))
		emd := fmt.Sprintf("%d/%d", em, ed)
		var d = fmt.Sprintf("%s(%s)", smd, weekdays[e.StartedAt.Weekday()])
		if smd != emd {
			d = d + fmt.Sprintf("-%s(%s)", emd, weekdays[e.EndedAt.Weekday()])
		}

		// time
		sh, _ := strconv.Atoi(e.StartedAt.Format("15"))
		sn := e.StartedAt.Format("04")
		eh, _ := strconv.Atoi(e.EndedAt.Format("15"))
		en := e.EndedAt.Format("04")
		t := fmt.Sprintf("%d:%s〜%d:%s", sh, sn, eh, en)

		// place
		var p string
		if e.Address != "" {
			p = e.Address
		} else {
			p = e.Place
		}

		// collect lines of title containing keyword
		if strings.Contains(strings.ToLower(e.Title), strings.ToLower(keyword)) {
			events = append(events, fmt.Sprintf("%s %s [%s] %s %s", d, t, p, e.Title, e.URL))
		}
	}

	// output order by datetime ascending
	for i := len(events) - 1; i >= 0; i-- {
		fmt.Println(events[i])
	}
}
