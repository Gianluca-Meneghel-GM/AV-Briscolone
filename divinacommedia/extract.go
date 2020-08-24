package divinacommedia

import (
	"antaresvision.com/corso-go-2/httputils"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type rowsResponse struct {
	TotalNumberOfRows int           `json:"total_number_of_rows"`
	Rows              []rowResponse `json:"rows"`
}

type rowResponse struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

func ExtractRows(divinaCommediaReader io.ReadSeeker, fromNumber int, toNumber int, output http.ResponseWriter) error {
	divinaCommediaReader.Seek(0, io.SeekStart)

	if divinaCommediaReader == nil {
		return errors.New(`input reader must be valid`)
	}
	if fromNumber > toNumber {
		return errors.New(`"from" must be less than or equal "to"`)
	}
	if output == nil {
		return errors.New(`output writer must be valid`)
	}

	buff, err := ioutil.ReadAll(divinaCommediaReader)
	if err != nil {
		return errors.New(err.Error())
	}
	str := string(buff)

	rows := strings.Split(str, "\n")
	if fromNumber >= len(rows) {
		return errors.New(`"from" must be less than max length`)
	}

	resp := rowsResponse{}
	resp.TotalNumberOfRows = len(rows)
	for i := fromNumber; (i <= toNumber) && (i < len(rows)); i++ {
		resp.Rows = append(resp.Rows, rowResponse{
			Number: i,
			Text:   rows[i],
		})
	}
	httputils.WriteJsonResponse(output, resp)

	return nil
}
