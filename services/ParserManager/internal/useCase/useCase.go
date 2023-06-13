package useCase

import (
	"context"
	"final/services/ParserManager/internal/domain"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ParserUseCase interface {
	GetNewBooks(context.Context) ([]*domain.Book, error)
}

type parserUseCase struct {
}

func (p *parserUseCase) GetNewBooks(ctx context.Context) ([]*domain.Book, error) {
	s1 := rand.NewSource(time.Now().Unix())
	r1 := rand.New(s1)

	size := r1.Intn(10)

	var books []*domain.Book

	for i := 1; i <= size; i++ {
		URL := fmt.Sprintf("https://openlibrary.org/trending/daily?page=%v", i)
		res, err := http.Get(URL)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, err
		}

		linkAll := doc.Find(".list-books").Find(".searchResultItem")
		linkAll.Each(func(j int, s *goquery.Selection) {
			var wg sync.WaitGroup
			wg.Add(3)
			var title string
			var yearInt int
			var author string

			go func() {
				defer wg.Done()
				title = s.Find("h3").Text()
				title = strings.TrimSpace(title)
			}()

			go func() {
				defer wg.Done()
				author = s.Find(".bookauthor").Find("a").Text()
				author = strings.TrimSpace(author)
			}()

			go func() {
				defer wg.Done()
				year := s.Find(".publishedYear").Text()
				year = strings.TrimSpace(year)
				year = strings.Replace(year, "First published in ", "", 1)
				yearInt, err = strconv.Atoi(year)
				if err != nil {
					log.Fatal(err)
					return
				}
			}()

			wg.Wait()

			fmt.Println(title, yearInt, author)

			books = append(books, &domain.Book{Title: title, Year: int32(yearInt), Author: author})
		})
	}

	return books, nil
}

func NewParserUseCase() ParserUseCase {
	return &parserUseCase{}
}
