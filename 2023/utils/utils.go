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

	"github.com/PuerkitoBio/goquery"
)

type Config struct {
	Token string `json:"token"`
}

type OutPut interface {
	int | string
}

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)

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

func SubmitSolutions[T OutPut](year int, day int, results []T, configPath string) error {
	for part, result := range results {
		err := postSolution(configPath, year, day, part+1, result)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func postSolution[T any](configPath string, year int, day int, part int, solution T) error {

	config, err := GetConfig(configPath)

	out, err := exec.Command("curl",
		"-i",
		"-b",
		fmt.Sprintf("session=%v", config.Token),
		"-X",
		"POST",
		"-d",
		fmt.Sprintf("answer=%v&level=%v", solution, part),
		fmt.Sprintf("https://adventofcode.com/%v/day/%v/answer", year, day)).Output()

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

func GetDayDesc(year int, day int, path string) error {
	now := time.Now().UTC()
	expected := time.Date(year, 12, day, 5, 0, 0, 0, time.UTC)
	if now.Before(expected) {
		return errors.New(fmt.Sprintf("Puzzle for day %v did not release yet", day))
	}

	config, err := GetConfig(path)

	if err != nil {
		return err
	}

	_, err = os.Stat(fmt.Sprintf("day%02d/description.md", day))
	if os.IsNotExist(err) {
		err = requestDesc(config, year, day)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("already downloaded, skip")
	}

	return nil
}

func requestDesc(config Config, year int, day int) error {
	now := time.Now().UTC()
	expected := time.Date(year, 12, day, 5, 0, 0, 0, time.UTC)
	if now.Before(expected) {
		return errors.New(fmt.Sprintf("Puzzle for day %v did not release yet", day))
	}

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://adventofcode.com/%v/day/%v", year, day),
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
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	heading := doc.Find("article").Find("h2").Text()
	doc.Find("article").Find("h2").Remove()

	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("# %s\n", heading))
	builder.WriteString(doc.Find("article").Text())

	foldername := fmt.Sprintf("day%02d", day)

	_, err = os.Stat(foldername)
	if os.IsNotExist(err) {
		err = os.Mkdir(foldername, os.ModePerm)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(foldername + "/description.md")

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(builder.String())

	if err != nil {
		return err
	}

	return nil
}

func GetDayInput(year int, day int, path string) error {
	now := time.Now().UTC()
	expected := time.Date(year, 12, day, 5, 0, 0, 0, time.UTC)
	if now.Before(expected) {
		return errors.New(fmt.Sprintf("Puzzle for day %v did not release yet", day))
	}

	config, err := GetConfig(path)

	if err != nil {
		return err
	}

	_, err = os.Stat(fmt.Sprintf("day%02d/input", day))
	if os.IsNotExist(err) {
		err = requestInput(config, year, day)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("already downloaded, skip")
	}

	return nil
}

func requestInput(config Config, year int, day int) error {

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day),
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

	foldername := fmt.Sprintf("day%02d", day)

	_, err = os.Stat(foldername)
	if os.IsNotExist(err) {
		err = os.Mkdir(foldername, os.ModePerm)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(foldername + "/input")

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
