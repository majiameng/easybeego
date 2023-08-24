package widget

type SelectParams struct {
	from   string
	option string
	value  interface{}
}

func Select(from string, option string, value interface{}) (out string) {
	out = `
    <select>
      <option value="">请选择</option>
      <option value="AAA">选项 A</option>
      <option value="BBB">选项 B</option>
      <option value="CCC">选项 C</option>
    </select>
`
	return out
}
