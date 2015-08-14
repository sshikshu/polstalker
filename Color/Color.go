package Color

func Black(s string, color bool) string {
	if color {
		return "\x1b[30m" + s + "\x1b[39;49m"
	}
	return s
}

func Red(s string, color bool) string {
	if color {
		return "\x1b[31m" + s + "\x1b[39;49m"
	}
	return s
}

func Green(s string, color bool) string {
	if color {
		return "\x1b[32m" + s + "\x1b[39;49m"
	}
	return s
}

func Yellow(s string, color bool) string {
	if color {
		return "\x1b[33m" + s + "\x1b[39;49m"
	}
	return s
}

func Blue(s string, color bool) string {
	if color {
		return "\x1b[34m" + s + "\x1b[39;49m"
	}
	return s
}

func Magenta(s string, color bool) string {
	if color {
		return "\x1b[35m" + s + "\x1b[39;49m"
	}
	return s
}

func Cyan(s string, color bool) string {
	if color {
		return "\x1b[36m" + s + "\x1b[39;49m"
	}
	return s
}

func White(s string, color bool) string {
	if color {
		return "\x1b[37m" + s + "\x1b[39;49m"
	}
	return s
}
