package dto

import "strconv"

func IDToDTO(id uint64) string {
	return strconv.FormatUint(id, 36)
}

func IDFromDTO(idDTO string) (uint64, error) {
	return strconv.ParseUint(idDTO, 36, 64)
}
