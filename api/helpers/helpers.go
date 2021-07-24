package helpers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func ParsePageAndLimit(c *fiber.Ctx) (uint64, uint64, error) {
	pageQuery := c.Query("page", "1")
	limitQuery := c.Query("limit", "10")
	page, err := strconv.ParseUint(pageQuery, 10, 64)
	if err != nil {
		return 0, 0, err
	}
	if page == 0 {
		page = 1
	}
	limit, err := strconv.ParseUint(limitQuery, 10, 64)
	if err != nil {
		return 0, 0, err
	}
	if limit == 0 {
		limit = 10
	}

	return page, limit, nil
}

//func ParseQueries()
