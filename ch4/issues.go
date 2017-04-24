// Issues prints a table of GitHub issues matching the search terms.
package ch4

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/angadthandi/gopl/ch4/github"
)

const monthsInYear = 12

func Issues() {
	var inLastMonth, inLastYear, beforeLastYear github.IssuesSearchResult

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		m := monthsDiff(item.CreatedAt)
		if m == 1 {
			inLastMonth.Items = append(inLastMonth.Items, item)
		} else if m < 12 {
			inLastYear.Items = append(inLastYear.Items, item)
		} else if m > 24 {
			beforeLastYear.Items = append(beforeLastYear.Items, item)
		}
	}

	if len(inLastMonth.Items) > 0 {
		fmt.Println("======================\n")
		fmt.Println("Issues in Last Month :\n")
		for _, bItem := range inLastMonth.Items {
			fmt.Printf("#%-5d %9.9s %.55s %v\n",
				bItem.Number, bItem.User.Login, bItem.Title,
				bItem.CreatedAt)
		}
	}
	if len(inLastYear.Items) > 0 {
		fmt.Println("======================\n")
		fmt.Println("Issues in Last Year :\n")
		for _, bItem := range inLastYear.Items {
			fmt.Printf("#%-5d %9.9s %.55s %v\n",
				bItem.Number, bItem.User.Login, bItem.Title,
				bItem.CreatedAt)
		}
	}
	if len(beforeLastYear.Items) > 0 {
		fmt.Println("======================\n")
		fmt.Println("Issues before Last Year :\n")
		for _, bItem := range beforeLastYear.Items {
			fmt.Printf("#%-5d %9.9s %.55s %v\n",
				bItem.Number, bItem.User.Login, bItem.Title,
				bItem.CreatedAt)
		}
	}
}

func monthsDiff(a time.Time) (months int) {
	// cur := b
	i_a, i_cur := 0, 0
	cur := time.Now().UTC()

	i_cur = int(cur.Month())
	i_a = int(a.Month())
	if cur.Year() > a.Year() {
		y := cur.Year() - a.Year()
		// addMonths := time.Month(monthsInYear * y)
		addMonths := (monthsInYear * y)
		months += (i_cur + addMonths) - i_a
	} else {
		// months += cur.Month() - a.Month()
		months += i_cur - i_a
	}

	return months
}
