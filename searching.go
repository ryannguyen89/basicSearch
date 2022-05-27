package basicSearch

func SearchOrg(key, val string, indexing Indexing) *Result {
	if idx, ok := indexing.OIndex.Indexes[key]; ok {
		if data, ok := idx[val]; ok {
			return &Result{
				Hits: int32(len(data)),
				Data: data,
			}
		}
		return nil
	}
	return nil
}

func SearchUser(key, val string, indexing Indexing) *Result {
	if idx, ok := indexing.UIndex.Indexes[key]; ok {
		if data, ok := idx[val]; ok {
			return &Result{
				Hits: int32(len(data)),
				Data: data,
			}
		}
		return nil
	}
	return nil
}

func SearchTicket(key, val string, indexing Indexing) *Result {
	if idx, ok := indexing.TIndex.Indexes[key]; ok {
		if data, ok := idx[val]; ok {
			return &Result{
				Hits: int32(len(data)),
				Data: data,
			}
		}
		return nil
	}
	return nil
}
