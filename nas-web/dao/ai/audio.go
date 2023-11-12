package ai

type audio struct{}

var Audio chat

func (a *audio) TextToSpeech(modelName string, question string) (string, error) {
	return "", nil
}
