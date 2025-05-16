package tools

import (
	"math"
)

type Page struct {
	TotalCount int
	Current    int
	Prev       int
	Next       int
	TotalPages int
	Pages      []int
}

var pageCount int = 2 // 默认每页显示3条

func SetPageCount(newPageCount int) {
	if newPageCount < 0 {
		return
	}
	pageCount = newPageCount
}

func GetPageCount() int {
	return pageCount
}

func GetPage(totalCount int, current int) Page {
	page := Page{
		TotalCount: totalCount,
		Current:    current,
		TotalPages: int(math.Ceil(float64(totalCount) / float64(pageCount))),
		Prev:       current - 1,
		Next:       current + 1,
	}

	// 上一页
	if current == 1 {
		page.Prev = 1
	}

	// 下一页
	if current == page.TotalPages {
		page.Next = page.TotalPages
	}

	// 安全处理 不能超出总页数
	if page.Current > page.TotalPages {
		page.Current = page.TotalPages
		page.Next = page.TotalPages
		page.Prev = page.Current - 1
	}

	if page.Current < 1 {
		page.Current = 1
		page.Prev = 1
		page.Next = page.Current + 1
	}

	for i := 0; i < page.TotalPages; i++ {
		page.Pages = append(page.Pages, i+1)
	}

	return page
}

// 1.只展示10页
// 2.一直展示有后4页，让当前页处于中间位置（前5页后4页）（第6页开始）
// 3.每页都只有10个数据

func baiduPage(page *Page) {
	// 传地址目的，不用返回值

	var start int = 1
	var end int = page.TotalPages

	// 大于10页要处理
	if page.TotalPages > 10 {
		// 1.如果前6页：显示前10条
		if page.Current <= 6 {
			end = 10
		}
		// 2.如果是后5页：显示后10页
		if page.Current >= page.TotalPages-4 {
			start = page.TotalPages - 4 - 5
		}

		// 3.中间页
		if page.Current > 6 && page.Current < page.TotalPages-4 {
			start = page.Current - 5
			end = page.Current + 4
		}
	}

	for i := start; i <= end; i++ {
		page.Pages = append(page.Pages, i)
	}
}
