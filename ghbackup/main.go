package main

// This file is part of tools released under the MIT license.
// See the LICENSE file for more information.

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func SetGitwebProperties(description string, owner string) {
	fmt.Printf("\tSetting the description property to: %s\n", description)
	cmd := exec.Command("git", "config", "--local", "--add", "gitweb.description", "'"+description+"'")
	err := cmd.Run()
	if err != nil {
		fmt.Println("an error occurred.\n")
		log.Fatal(err)
		return
	}
	fmt.Printf("\tSetting the owner property to: %s\n", owner)
	cmd = exec.Command("git", "config", "--local", "--add", "gitweb.owner", "'"+owner+"'")
	err = cmd.Run()
	if err != nil {
		fmt.Println("an error occurred.\n")
		log.Fatal(err)
		return
	}
}

func RetrieveRepo(name string, clone_url string, description string, owner string) {
	fmt.Printf("\nStarting retrieving repository %s with:\n\tname : %s \n\tdesc : %s\n\towner: %s\n", clone_url, name, description, owner)
	var elements []string = strings.Split(clone_url, "/")
	var directory string = elements[len(elements)-1]

	isExist, err := exists("./" + directory)
	if err != nil {
		fmt.Println("Can't figure out the state for the directory " + directory + ". Please check!")
		return
	}

	if isExist {
		fmt.Println("\tthe repository already exists: time to update it!")
		cmd := exec.Command("git", "remote", "update")
		err = cmd.Run()
		if err != nil {
			fmt.Print("Couldn't update the repository:\n\t")
			fmt.Println(err)
			return
		}
		fmt.Println("The repository is now up-to-date.")
		return
	}

	//fmt.Printf("CMD: %s %s %s %s\n", "git", "clone", "--mirror", clone_url)
	cmd := exec.Command("git", "clone", "--mirror", clone_url)
	err = cmd.Run()
	if err != nil {
		fmt.Print("Couldn't retrieve the repository:\n\t")
		fmt.Println(err)
		return
	}

	os.Chdir(directory)
	desc := description
	if desc == "" {
		desc = name
	}
	SetGitwebProperties(desc, owner)
	os.Chdir("..")

	fmt.Println("Done!")
}

func main() {
	dec := json.NewDecoder(os.Stdin)
	for {
		var v []map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			fmt.Println("\n--THE END--")
			return
		}
		fmt.Println("Github repository mirroring")
		fmt.Printf("About to mirror %d repositories!\n---\n", len(v))
		for i := 0; i < len(v); i++ {
			//for i := 0; i < 2; i++ {
			name := v[i]["name"].(string)
			description := v[i]["description"].(string)
			clone_url := v[i]["clone_url"].(string)
			var owner map[string]interface{} = v[i]["owner"].(map[string]interface{})
			owner_login := owner["login"].(string)

			RetrieveRepo(name, clone_url, description, owner_login)
		}
	}
}
