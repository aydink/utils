package utils

import (
	"bytes"
)

func BinarySearch(a []int, value int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (low + high) / 2

		if a[mid] > value {
			high = mid - 1
		} else if a[mid] < value {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func FindFirst(a []int, value int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (low + high) / 2

		if a[mid] > value {
			high = mid - 1
		} else if a[mid] < value {
			low = mid + 1
		} else {
			//return mid
			if (mid == 0) || (a[mid-1] != value) {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func FindLast(a []int, value int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (low + high) / 2

		if a[mid] > value {
			high = mid - 1
		} else if a[mid] < value {
			low = mid + 1
		} else {
			//return mid
			if (mid == high) || (a[mid+1] != value) {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

func NearestFirst(a []int, value int) int {
	low := 0
	high := len(a) - 1
	seen_higher := false

	var mid int

	for low <= high {
		mid = (low + high) / 2

		if a[mid] > value {
			high = mid - 1
			seen_higher = true
		} else if a[mid] < value {
			low = mid + 1
		} else {
			//return mid
			if (mid == 0) || (a[mid-1] != value) {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	if seen_higher {
		return mid
	} else {
		return -1
	}
}

func NearestLast(a []int, value int) int {
	low := 0
	high := len(a) - 1
	seen_lower := false

	var mid int

	for low <= high {
		mid = (low + high) / 2

		if a[mid] > value {
			high = mid - 1
		} else if a[mid] < value {
			low = mid + 1
			seen_lower = true
		} else {
			//return mid
			if (mid == high) || (a[mid+1] != value) {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	if seen_lower {
		if mid == high {
			return mid
		} else {
			return mid - 1
		}
	} else {
		return -1
	}
}

func DateTimeFindFirst(a FolderMeta, value int64) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (low + high) / 2

		if a[mid].DateTime > value {
			high = mid - 1
		} else if a[mid].DateTime < value {
			low = mid + 1
		} else {
			//return mid
			if (mid == 0) || (a[mid-1].DateTime != value) {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func DateTimeFindLast(a FolderMeta, value int64) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (low + high) / 2

		if a[mid].DateTime > value {
			high = mid - 1
		} else if a[mid].DateTime < value {
			low = mid + 1
		} else {
			//return mid
			if (mid == high) || (a[mid+1].DateTime != value) {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

func DateTimeNearestFirst(a FolderMeta, value int64) int {
	low := 0
	high := len(a) - 1
	seen_higher := false

	var mid int

	for low <= high {
		mid = (low + high) / 2

		if a[mid].DateTime > value {
			high = mid - 1
			seen_higher = true
		} else if a[mid].DateTime < value {
			low = mid + 1
		} else {
			//return mid
			if (mid == 0) || (a[mid-1].DateTime != value) {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	if seen_higher {
		return mid
	} else {
		return -1
	}
}

func DateTimeNearestLast(a FolderMeta, value int64) int {
	low := 0
	high := len(a) - 1
	seen_lower := false

	var mid int

	for low <= high {
		mid = (low + high) / 2

		if a[mid].DateTime > value {
			high = mid - 1
		} else if a[mid].ModTime < value {
			low = mid + 1
			seen_lower = true
		} else {
			//return mid
			if (mid == high) || (a[mid+1].ModTime != value) {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	if seen_lower {
		if mid == high {
			return mid
		} else {
			return mid - 1
		}
	} else {
		return -1
	}
}

func HashFindFirst(a FolderMeta, value []byte) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (low + high) / 2

		if bytes.Compare(a[mid].Hash, value) > 0 {
			high = mid - 1
		} else if bytes.Compare(a[mid].Hash, value) < 0 {
			low = mid + 1
		} else {
			//return mid
			if (mid == 0) || (bytes.Compare(a[mid-1].Hash, value) != 0) {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func HahsFindLast(a FolderMeta, value []byte) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (low + high) / 2

		if bytes.Compare(a[mid].Hash, value) > 0 {
			high = mid - 1
		} else if bytes.Compare(a[mid].Hash, value) < 0 {
			low = mid + 1
		} else {
			//return mid
			if (mid == high) || (bytes.Compare(a[mid+1].Hash, value) != 0) {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}
