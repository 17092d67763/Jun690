package client

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// GetAllItems returns a list of all Items in the DB
func GetAllItems(itemType string, port string, verbose bool) []byte {

	host := viper.GetString("host")

	url := "http://" + host + ":" + port + "/api/v1/" + itemType

	if verbose {
		fmt.Println("GET: " + url)
	}

	resp, err := http.Get(url)
	if err != nil {
		// handle error
		fmt.Println("An error occurred")
		fmt.Println(err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	return data
}

// DeleteItem deletes the given item
func DeleteItem(id string, pathID string, pathName string, port string, verbose bool) []byte {

	// The ID parameter can be either NAME or ID. We are doing this to allow the user
	// enter either the name or the ID of an object to delete.
	// First, we try ID. If successful, stop. If unsuccessful, try name.

	host := viper.GetString("host")

	// Create client
	client := &http.Client{}

	// Try ID first
	url := "http://" + host + ":" + port + "/api/v1/" + pathID + id

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	if string(respBody) == "true" {
		// deleting with ID worked
		if verbose {
			fmt.Println("DELETE: " + url)
		}
		return respBody
	}

	if pathName == "" {
		fmt.Println("Deleting by ID failed: " + url)
	}

	// deleting with IF failed. trying with name/slug
	url = "http://" + host + ":" + port + "/api/v1/" + pathName + id

	req, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Fetch Request
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return respBody

}

// GetVersion returns the version of a service given its port
func GetVersion(port string) []byte {
	host := viper.GetString("host")

	url := "http://" + host + ":" + port + "/version"

	resp, err := http.Get(url)
	if err != nil {
		// handle error
		fmt.Println("An error occurred")
		fmt.Println(err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	return data
}
