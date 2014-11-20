package utils

/*
GetFilesWithDateTime returns a slice containing elements which match timestamp
*/
func GetFilesWithDateTime(fm FolderMeta, timestamp int64) FolderMeta {
	start := DateTimeFindFirst(fm, timestamp)
	end := DateTimeFindLast(fm, timestamp)

	if start > -1 {
		return fm[start : end+1]
	} else {
		return FolderMeta{}
	}
}

/*
GetFilesWithDateTimeRange returns a slice containing elements with DateTime bigger
than timestampStart and smaller than timestampEnd
*/
func GetFilesWithDateTimeRange(fm FolderMeta, timestampStart, timestampEnd int64) FolderMeta {
	start := DateTimeNearestFirst(fm, timestampStart)
	end := DateTimeNearestLast(fm, timestampEnd)

	if start > -1 && end > -1 {
		return fm[start : end+1]
	} else {
		return FolderMeta{}
	}
}

/*
GetExactMatches returns a slice containing elements which match timestamp
*/
func SearchHash(fm FolderMeta, hash []byte) FolderMeta {
	start := HashFindFirst(fm, hash)
	end := HahsFindLast(fm, hash)

	if start > -1 {
		return fm[start : end+1]
	} else {
		return FolderMeta{}
	}
}

func GetExactMatches(fmLeft, fmRight FolderMeta) FolderMeta {
	matches := make(FolderMeta, 0)

	for i := range fmLeft {
		start := HashFindFirst(fmRight, fmLeft[i].Hash)
		end := HahsFindLast(fmRight, fmLeft[i].Hash)

		if start > -1 {
			matches = append(matches, fmRight[start:end+1]...)
		}
	}

	return matches
}

func IntersectDatetTime(fmLeft, fmRight FolderMeta) FolderMeta {
	matches := make(FolderMeta, 0)

	for i := range fmLeft {
		start := DateTimeFindFirst(fmRight, fmLeft[i].DateTime)
		end := DateTimeFindLast(fmRight, fmLeft[i].DateTime)

		if start > -1 {
			matches = append(matches, fmRight[start:end+1]...)
		}
	}

	return matches
}
