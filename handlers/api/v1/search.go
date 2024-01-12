package v1

import (
	"github.com/AkshachRd/leards-backend-go/httputils"
	_ "github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SearchPublicStorages godoc
// @Id			 searchPublicStorages
// @Summary      Get public storages with search
// @Description  fetches public storages from the database
// @Tags         search
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        page         query     int       false    "Page" minimum(1)
// @Param        page_size    query     int       false    "Page size" minimum(10)    maximum(100)
// @Param		 search_type  query     string    true     "Search type" Enums(all, tag, name)
// @Param        sort_type    query     string    true     "Sort type" Enums(rating, name)
// @Param        order_by     query     string    true     "Order by" Enums(asc, desc)
// @Param        name		  query     string    false    "Name"
// @Param        tags		  query     []string  false    "Tags" collectionFormat(multi)
// @Success      200  {array}   httputils.SearchResult
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /search [get]
func SearchPublicStorages(c *gin.Context) {
	searchService, err := services.NewSearchService(
		c.Query("search_type"),
		c.Query("sort_type"),
		c.Query("order_by"),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if pageStr, ok := c.GetQuery("page"); ok {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page input"})
			return
		}
		searchService.SetPage(page)
	}

	if pageSizeStr, ok := c.GetQuery("page_size"); ok {
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size input"})
			return
		}
		searchService.SetPageSize(pageSize)
	}

	searchResults, err := searchService.Search(c.Query("name"), c.QueryArray("tags"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, httputils.ConvertSearchResults(searchResults))
}
