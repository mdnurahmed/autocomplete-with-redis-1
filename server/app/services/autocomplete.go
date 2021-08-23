package services

import (
	"autocomplete/app/repositories"
	"strings"
)

type IAutocompleteService interface {
	Search(word string) (result []string, err error)
	Insert(word string) (err error)
	Delete(key string) (err error)
}
type AutocompleteService struct {
	keyName         string
	searchLength    int64
	redisRepository repositories.IRedisRepository
}

func NewInstanceOfAutocompleteService(redisRepository repositories.IRedisRepository, keyName string, searchLength int64) AutocompleteService {
	return AutocompleteService{redisRepository: redisRepository, keyName: keyName, searchLength: searchLength}
}

func (a *AutocompleteService) Search(word string) ([]string, error) {
	result, err := a.redisRepository.Search(word, a.keyName, a.searchLength)
	if err != nil {
		return []string{}, err
	}
	curatedResult := []string{}
	for i := 0; i < len(result); i++ {
		lastIndex := len(result[i]) - 1
		if lastIndex >= 0 && result[i][lastIndex] == '*' && strings.HasPrefix(result[i], word) {
			curatedResult = append(curatedResult, result[i])
		}
	}
	if len(curatedResult) == 0 {
		lastWordIndex := len(result)
		if lastWordIndex >= 0 && strings.HasPrefix(result[lastWordIndex], word) {
			curatedResult = append(curatedResult, result[lastWordIndex])
		}
	}
	return curatedResult, nil
}

func (a *AutocompleteService) Insert(word string) error {
	err := a.redisRepository.Insert(word, a.keyName)
	if err != nil {
		return err
	}
	for i := 1; i < len(word); i++ {
		var prefix []byte
		for j := 0; j < i; j++ {
			prefix = append(prefix, word[j])
		}
		if len(prefix) != 0 {
			err := a.redisRepository.Insert(string(prefix), a.keyName)
			if err != nil {
				return err
			}
		}
	}
	endWord := word + "*"
	err = a.redisRepository.Insert(endWord, a.keyName)
	if err != nil {
		return err
	}
	return nil
}

func (a *AutocompleteService) Delete(key string) error {
	err := a.redisRepository.Delete(key)
	return err
}
