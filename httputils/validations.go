package httputils

const maxAvatarSize = int64(2 * 1024 * 1024) // 2 MB

func IsAvatarInSizeRange(size int64) bool {
	return size <= maxAvatarSize && size > 0
}

func getAvatarAllowedExtensions() map[string]bool {
	return map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}
}

func IsAvatarHasAllowedExtension(extension string) bool {
	if _, ok := getAvatarAllowedExtensions()[extension]; !ok {
		return false
	}

	return true
}
