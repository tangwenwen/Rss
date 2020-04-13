package pagination

import "errors"

func GetORMLimitOffset(pageNo, pageRows, totalCount int) (int, int, error) {
	limit := 0
	offset := 0

	if pageNo < 1 {
		return 0, 0, errors.New("page rows need more then 1 or equals 1")
	} else {
		limit = pageRows
	}

	pageNo--
	if pageNo > 0 {
		offset = pageNo * pageRows
	} else {
		offset = 0
	}

	pageSize := GetPageCnt(limit, totalCount)
	if pageNo >= pageSize {
		if pageSize-1 >= 0 {
			offset = (pageSize - 1) * pageRows
		}
	}

	return limit, offset, nil

}
func GetPageCnt(limit, totalCount int) int {
	mod := totalCount % limit
	pageCnt := totalCount / limit
	if mod > 0 {
		pageCnt++
	}
	return pageCnt
}
