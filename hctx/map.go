package hctx

type Map map[string]interface{}

func Merge(maps ...Map) Map {
	mx := map[string]interface{}{}
	for _, m := range maps {
		for k, v := range m {
			mx[k] = v
		}
	}
	return mx
}
