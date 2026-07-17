package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

type ClassModel struct {
	Id   int    `l:"id"`
	Name string `l:"name"`
}

func camelToSnake(s string) string {
	var result []rune

	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}

	return string(result)
}

func Find(obj any, query ...any) (sql string, err error) {
	var where string
	var nameList []string
	t := reflect.TypeOf(obj)

	// 必须要判断类型，若不是结构体，则获取不到相关信息
	if t.Kind() != reflect.Struct {
		err = errors.New("obj必须是结构体")
		return
	}

	// 处理查询条件
	if len(query) > 0 {
		q := query[0]
		// 第一个参数必须是字符串
		qs, ok := q.(string)
		if !ok {
			err = errors.New("查询的第一个参数必须是字符串")
			return
		}
		// 算？个数
		if strings.Count(qs, "?")+1 != len(query) {
			err = errors.New("查询参数个数不匹配")
			return
		}
		end := query[1:]
		for _, v := range end {
			switch s := v.(type) {
			case string:
				// strings.Contains(s, "x")      // 是否包含
				// strings.HasPrefix(s, "x")     // 是否以 x 开头
				// strings.HasSuffix(s, "x")     // 是否以 x 结尾
				// strings.Split(s, ",")         // 分割
				// strings.Join(arr, ",")        // 拼接
				// strings.Replace(s, "?", "x", 1)  // 替换 1 次
				// strings.ReplaceAll(s, "?", "x")  // 替换全部
				// strings.TrimSpace(s)          // 去掉前后空格
				// strings.ToLower(s)            // 转小写
				// strings.ToUpper(s)            // 转大写
				qs = strings.Replace(qs, "?", fmt.Sprintf("'%s'", s), 1)

			case int:
				qs = strings.Replace(qs, "?", fmt.Sprintf("%d", s), 1)
			}
		}

		where = "where " + qs
		// endCount := query[1:]
	}
	for i := 0; i < t.NumField(); i++ {
		// 拼接所有 有tag 的字段
		filed := t.Field(i).Tag.Get("l")
		if filed == "" {
			continue
		}
		nameList = append(nameList, filed)
	}

	sql = fmt.Sprintf("select %s form %ss %s", strings.Join(nameList, ","), camelToSnake(t.Name()), where)

	return
}

func main() {
	sql, err := Find(ClassModel{}, "name = ? ", "三年一班")
	fmt.Println(sql, err)
	// //  select name,id form class where name = '三年一班'

	sql, err = Find(ClassModel{}, "name = ?  and id = ?", "三年一班", 1)
	fmt.Println(sql, err)

	// select name,id form class where name = '三年一班' and id = 1

	sql, err = Find(ClassModel{})
	fmt.Println(sql, err)

	// select name,id form class
}
