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
	Tags            []string `json:"tags"`
} // @name SearchResult

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
	results := &[]SearchResult{}
	var err error

	switch ss.searchType {
	case "all":
	case "tag":
	case "name":
		results, err = ss.searchByName(name)
	default:
		return &[]SearchResult{}, fmt.Errorf("unknown search type: %s", ss.searchType)
	}
	if err != nil {
		return &[]SearchResult{}, err
	}

	for i, result := range *results {
		tags, err := models.FetchTagsByStorageIdAndStorageType(result.ID, result.Type)
		if err != nil {
			return &[]SearchResult{}, err
		}
		for _, tag := range *tags {
			(*results)[i].Tags = append((*results)[i].Tags, tag.Name)
		}
	}

	return results, nil
}

func (ss *SearchService) searchByName(name string) (*[]SearchResult, error) {
	searchResults, err := models.SearchByNameWithPagination(name, ss.sortType, ss.orderBy, ss.page, ss.pageSize)
	if err != nil {
		return &[]SearchResult{}, err
	}

	var results []SearchResult

	for _, searchResult := range *searchResults {
		results = append(results, SearchResult{
			ID:              searchResult.ID,
			Name:            searchResult.Name,
			Rating:          searchResult.Rating,
			Type:            searchResult.Type,
			ProfileIconPath: searchResult.ProfileIconPath,
			Tags:            make([]string, 0),
		})
	}

	return &results, nil
}
