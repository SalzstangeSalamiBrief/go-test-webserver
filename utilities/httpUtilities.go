package utilities

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func ReadBody[T interface{}](r *http.Request) (T, error) {
	var result T
	body := r.Body
	defer body.Close()
	if body == nil {
		return result, errors.New("could not access request Body because it is empty")
	}

	decodeErr := json.NewDecoder(body).Decode(&result)
	return result, decodeErr
}

func GetIdFromPath(r *http.Request) (int, error) {
	pathParts := strings.Split(r.URL.Path, "/")
	data, err := strconv.Atoi(pathParts[len(pathParts)-1])
	return data, err
}

func GetIdsFromUrl(r *http.Request, path string) []int {
	re := regexp.MustCompile(`\d+`)
	stringMatches := re.FindAllString(r.URL.Path, -1)

	result := make([]int, 0)
	for _, value := range stringMatches {
		integerValue, _ := strconv.Atoi(value)
		result = append(result, integerValue)
	}

	return result
}
