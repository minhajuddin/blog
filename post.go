package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Post struct {
	ID    int
	Title string
	Body  string
}

const dataDir = "./data"

func findPost(ID int) *Post {
	fp := filepath.Join(dataDir, strconv.Itoa(ID))
	f, err := os.Open(fp)
	if err != nil {
		log.Println(err)
		return nil
	}
	p := Post{}
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &p
}

//save this post
func (p *Post) save() {
	//create parent directory if it is not present
	err := os.MkdirAll(dataDir, 0755)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fp := filepath.Join(dataDir, strconv.Itoa(p.ID))
	f, err := os.Create(fp)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	//write as json
	err = json.NewEncoder(f).Encode(p)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
}

//loads all posts
func allPosts() []*Post {
	posts := make([]*Post, 0, 10)

	err := filepath.Walk(dataDir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			posts = make([]*Post, 0)
			return err
		}

		if info.IsDir() {
			return nil
		}

		idstr := filepath.Base(p)
		id, err := strconv.Atoi(idstr)
		if err != nil {
			log.Println(err)
			return err
		}
		posts = append(posts, findPost(id))
		return nil
	})

	if err != nil {
		log.Println(err)
		return posts
	}

	return posts
}
