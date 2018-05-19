package sportsdatabase

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
)

//Excel convert SQDL response to xlsx
func (response SQDLResponse) Excel(sheetName string) (*xlsx.File, error) {
	f := xlsx.NewFile()
	sheet, err := f.AddSheet(sheetName)
	if err != nil {
		return nil, errors.Wrap(err, "can't create sheet")
	}

	if sheet.MaxRow == 0 {
		headers := sheet.AddRow()
		for _, h := range response.Headers {
			cell := headers.AddCell()
			cell.SetValue(h)
		}
	}

	raw := response.Groups[0].Columns

	colCount, rowCount := len(raw), len(raw[0])
	for r := 0; r < rowCount; r++ {
		row := sheet.AddRow()
		for c := 0; c < colCount; c++ {
			v := raw[c][r]
			switch x := v.(type) {
			case []interface{}:
				arr := make([]string, len(x))
				for i, y := range x {
					arr[i] = fmt.Sprint(y)
				}
				v = strings.Join(arr, ",")
			}
			// log.Printf("<%d,%d> %+v", c, r, v)
			cell := row.AddCell()
			cell.SetValue(v)
		}
	}

	return f, nil
}
