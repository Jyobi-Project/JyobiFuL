package answerquestion

func GetMine(language string) string {
	switch language {
	case "java":
		return ".java"
	case "python":
		return ".py"
	default:
		return ".txt"
	}
}
