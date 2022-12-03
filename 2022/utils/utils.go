package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Token string `json:"token"`
}

func ReadFile(day int) ([]string, error) {
	file, err := os.Open(fmt.Sprintf("day%02d/input", day))

	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	var contents []string
	for fileScanner.Scan() {
		contents = append(contents, fileScanner.Text())
	}

	return contents, nil
}

func GetConfig(path string) (Config, error) {
	fileContent, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("could not open file")
		return Config{}, err
	}

	config := Config{}

	err = json.Unmarshal(fileContent, &config)

	if err != nil {
		fmt.Println("could not parse json")
		return Config{}, err
	}

	return config, nil
}

func SubmitSolutions(day int, results []int, configPath string) error {
	for part, result := range results {
		err := postSolution(configPath, day, part+1, result)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func postSolution(configPath string, day int, part int, solution int) error {

	config, err := GetConfig(configPath)

	out, err := exec.Command("curl",
		"-i",
		"-b",
		fmt.Sprintf("session=%v", config.Token),
		"-X",
		"POST",
		"-d",
		fmt.Sprintf("answer=%v&level=%v", solution, part),
		fmt.Sprintf("https://adventofcode.com/%v/day/%v/answer", 2022, day)).Output()

	if err != nil {
		return err
	}

	body := fmt.Sprintf("%s", out)
	if strings.Contains(body, "Did you already complete it?") {
		return errors.New(fmt.Sprintf("Day %v, Part %v: puzzle already completed", day, part))
	}
	if strings.Contains(body, "That's not the right answer.") {
		return errors.New(fmt.Sprintf("Day %v, Part %v: wrong solution", day, part))
	}

	fmt.Printf("Successfully submitted Day %v Part %v\n", day, part)

	return nil

}

func GetDayInput(day int, path string) error {
	config, err := GetConfig(path)

	if err != nil {
		return err
	}

	err = requestInput(config, day)
	if err != nil {
		return err
	}

	return nil
}

func requestInput(config Config, day int) error {
	now := time.Now().UTC()
	expected := time.Date(2022, 12, day, 5, 0, 0, 0, time.UTC)
	if now.Before(expected) {
		return errors.New(fmt.Sprintf("Puzzle for day %v did not release yet", day))
	}

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", 2022, day),
		nil)

	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: config.Token})

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("Status code: " + strconv.Itoa(resp.StatusCode))
		return err
	}

	_, err = os.Stat("inputs")
	if os.IsNotExist(err) {
		err = os.Mkdir("inputs", os.ModePerm)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(fmt.Sprintf("day%02d/input", day))

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		return err
	}
	return nil
}
