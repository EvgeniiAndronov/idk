package main

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func Test_CreatingUrl(t *testing.T) {
	testCases := []struct {
		name      string
		question  string
		sorted_by int
		expected  string
	}{
		{
			name:      "На_русском",
			question:  "Вопрос",
			sorted_by: 0,
			expected:  fmt.Sprintf("https://newsapi.org/v2/everything?q=Вопрос&from=2024-11-18&sortBy=relevancy&apiKey=%s", API_KEY),
		}, {
			name:      "На_английском",
			question:  "engl",
			sorted_by: 0,
			expected:  fmt.Sprintf("https://newsapi.org/v2/everything?q=engl&from=2024-11-18&sortBy=relevancy&apiKey=%s", API_KEY),
		}, {
			name:      "popularity",
			question:  "Вопрос",
			sorted_by: 1,
			expected:  fmt.Sprintf("https://newsapi.org/v2/everything?q=Вопрос&from=2024-11-18&sortBy=popularity&apiKey=%s", API_KEY),
		}, {
			name:      "publishedAt",
			question:  "Вопрос",
			sorted_by: 2,
			expected:  fmt.Sprintf("https://newsapi.org/v2/everything?q=Вопрос&from=2024-11-18&sortBy=publishedAt&apiKey=%s", API_KEY),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := createUrl(tc.question, tc.sorted_by)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func Test_MakeTitles(t *testing.T) {
	testCases := []struct {
		name     string
		article  []Article
		expected []string
	}{
		{
			name: "2 obj",
			article: []Article{
				{
					Sources: Source{
						Id:   "Id",
						Name: "Name 1",
					},
					Author:      "Author",
					Title:       "Title 1",
					Url:         "Url",
					UrlToImage:  "ImgUrl",
					PublishedAt: "PublishedAt",
					Content:     "Content",
				}, {
					Sources: Source{
						Id:   "Id",
						Name: "Name 2",
					},
					Author:      "Author",
					Title:       "Title 2",
					Url:         "Url",
					UrlToImage:  "ImgUrl",
					PublishedAt: "PublishedAt",
					Content:     "Content",
				},
			},
			expected: []string{"Title 1", "Title 2"},
		}, {
			name: "1 obj",
			article: []Article{
				{
					Sources: Source{
						Id:   "Id",
						Name: "Name 1",
					},
					Author:      "Author",
					Title:       "Title 1",
					Url:         "Url",
					UrlToImage:  "ImgUrl",
					PublishedAt: "PublishedAt",
					Content:     "Content",
				},
			},
			expected: []string{"Title 1"},
		}, {
			name: "IDK",
			article: []Article{
				{
					Sources: Source{
						Id:   "Id",
						Name: "Name 1",
					},
					Author:      "Author",
					Title:       "Title 1",
					Url:         "Url",
					UrlToImage:  "ImgUrl",
					PublishedAt: "PublishedAt",
					Content:     "Content",
				}, {
					Sources: Source{
						Id:   "Id",
						Name: "Name 2",
					},
					Author:      "Author",
					Title:       "Title 2",
					Url:         "Url",
					UrlToImage:  "ImgUrl",
					PublishedAt: "PublishedAt",
					Content:     "Content",
				}, {
					Sources: Source{
						Id:   "Id",
						Name: "Name 3",
					},
					Author:      "Author",
					Title:       "Title 3",
					Url:         "Url",
					UrlToImage:  "ImgUrl",
					PublishedAt: "PublishedAt",
					Content:     "Content",
				},
			},
			expected: []string{"Title 1", "Title 2", "Title 3"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := makeTitles(tc.article)
			assert.Equal(t, tc.expected, result)
		})
	}
}
