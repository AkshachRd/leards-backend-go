package services

import (
	"fmt"
	"github.com/AkshachRd/leards-backend-go/models"
)

type SearchResult struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Rating          uint     `json:"rating"`
	Type            string   `json:"type"`
	ProfileIconPath string   `json:"profileIconPath"`
	AuthorName      string   `json:"authorName"`
	Tags            []string `json:"tags"`
} // @name SearchResult

func (sr *SearchResult) fetchTags() error {
	tagNames, err := models.FetchTagNamesByStorageIdAndStorageType(sr.ID, sr.Type)
	if err != nil {
		return err
	}

	if tagNames == nil {
		tagNames = &[]string{}
	}

	sr.Tags = *tagNames
	return nil
}

type SearchService struct {
	page       int
	pageSize   int
	searchType string
	sortType   string
	orderBy    string
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func NewSearchService(searchType string, sortType string, orderBy string) (*SearchService, error) {
	validOrderBys := []string{"asc", "desc"}
	if !contains(validOrderBys, orderBy) {
		return &SearchService{}, fmt.Errorf("invalid order by: %s", orderBy)
	}

	validSortTypes := []string{"rating", "name"}
	if !contains(validSortTypes, sortType) {
		return &SearchService{}, fmt.Errorf("invalid sort type: %s", sortType)
	}

	return &SearchService{page: 1, pageSize: 10, sortType: sortType, searchType: searchType, orderBy: orderBy}, nil
}

func (ss *SearchService) GetPage() int {
	return ss.page
}

func (ss *SearchService) SetPage(page int) {
	if page <= 0 {
		ss.page = 1
		return
	}
	ss.page = page
}

func (ss *SearchService) GetPageSize() int {
	return ss.pageSize
}

func (ss *SearchService) SetPageSize(pageSize int) {
	switch {
	case pageSize > 100:
		ss.pageSize = 100
		return
	case pageSize <= 0:
		ss.pageSize = 10
		return
	}
	ss.pageSize = pageSize
}

func (ss *SearchService) Search(
	name string,
	tags []string,
) (*[]SearchResult, error) {
	results := make([]SearchResult, 0)

	var err error
	var searchResults *[]models.SearchResult
	switch ss.searchType {
	case "all":
		searchResults, err = models.SearchByNameOrTagsWithPagination(name, tags, ss.sortType, ss.orderBy, ss.page, ss.pageSize)
	case "tag":
		searchResults, err = models.SearchByTagsWithPagination(tags, ss.sortType, ss.orderBy, ss.page, ss.pageSize)
	case "name":
		searchResults, err = models.SearchByNameWithPagination(name, ss.sortType, ss.orderBy, ss.page, ss.pageSize)
	default:
		err = fmt.Errorf("unknown search type: %s", ss.searchType)
	}
	if err != nil {
		return &[]SearchResult{}, err
	}

	for _, searchResult := range *searchResults {
		result := SearchResult{
			ID:              searchResult.ID,
			Name:            searchResult.Name,
			Rating:          searchResult.Rating,
			Type:            searchResult.Type,
			AuthorName:      searchResult.AuthorName,
			ProfileIconPath: searchResult.ProfileIconPath.String,
			Tags:            make([]string, 0),
		}

		if err := result.fetchTags(); err != nil {
			return &[]SearchResult{}, err
		}

		results = append(results, result)
	}

	return &results, nil
}
