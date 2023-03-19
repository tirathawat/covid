package covid

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/tirathawat/covid/pkg/errs"
)

var (
	// ErrFetchFailed represents the error when failed to fetch COVID data.
	ErrFetchFailed = errs.New("failed to fetch COVID data")
)

// Handler provides the covid handler.
type Handler interface {
	Summary(c *gin.Context)
}

type handler struct {
	fetcher Fetcher
}

// NewHandler creates a new covid handler.
func NewHandler(fetcher Fetcher) *handler {
	return &handler{
		fetcher: fetcher,
	}
}

// SummaryResponse represents the summary response.
type SummaryResponse struct {
	Province ProvinceCount `json:"Province"`
	AgeGroup AgeGroupCount `json:"AgeGroup"`
}

// Summary handles the summary request.
// It fetches the data from the COVID API and returns the summary.
// The summary contains the number of cases by province and age group.
func (h *handler) Summary(c *gin.Context) {
	data, err := h.fetcher.Fetch()
	if err != nil {
		log.Error().Err(err).Msg(ErrFetchFailed.Msg)
		c.JSON(http.StatusServiceUnavailable, ErrFetchFailed)
		return
	}

	c.JSON(http.StatusOK, SummaryResponse{
		Province: CountCasesByProvince(data.Provinces()),
		AgeGroup: CountCasesByAgeGroup(data.Ages()),
	})
}
