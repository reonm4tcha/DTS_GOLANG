package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Friend struct {
	ID      int
	Name    string
	Address string
	Job     string
	Reason  string
}

func getFriendById(id int) (Friend, error) {
	friends := map[int]Friend{
    1: {
			ID:      1,
			Name:    "Amalia Rahardjo",
			Address: "Cileungsi",
			Job:     "Marketing Development",
			Reason:  "Belajar",
		},
		2: {
			ID:      2,
			Name:    "Budi Triono",
			Address: "Cileungsi",
			Job:     "Software Engineer",
			Reason:  "Belajar",
		},
		3: {
			ID:      3,
			Name:    "Dea Fitria",
			Address: "Cileungsi",
			Job:     "Data Analyst",
			Reason:  "Belajar",
		},
		4: {
			ID:      4,
			Name:    "Prayogi Adhitya",
			Address: "Gas Alam",
			Job:     "Fullstack Developer",
			Reason:  "Belajar",
		},
		5: {
			ID:      5,
			Name:    "Sandi Hidayat",
			Address: "Citayam",
			Job:     "Mobile Developer",
			Reason:  "Belajar",
		},
		6: {
			ID:      6,
			Name:    "Rifq Afta",
			Address: "Tanah Baru",
			Job:     "Network Engineer",
			Reason:  "Belajar",
		},
		7: {
			ID:      7,
			Name:    "Andy Yusuf",
			Address: "Tebet",
			Job:     "Quality Assurance",
			Reason:  "Belajar",
		},
		8: {
			ID:      8,
			Name:    "Wildan Fajrie",
			Address: "Plaju",
			Job:     "Freelance",
			Reason:  "Belajar",
		},
		9: {
			ID:      9,
			Name:    "Irene",
			Address: "Seoul",
			Job:     "Idol",
			Reason:  "Belajar",
		},
		10: {
			ID:      10,
			Name:    "Marvelio Prado",
			Address: "Citereup",
			Job:     "Frontend Developer",
			Reason:  "Belajar",
		},
	}

	if value, exist := friends[id]; exist {
		return value, nil
	}

	return Friend{}, errors.New(fmt.Sprintf("Friend with id %v not found", id))
}

func main() {
	if len(os.Args) == 1 {
		panic("ID belum dimasukkan!")
	}
	var friendID = os.Args[1]
	var friendIDInt, _ = strconv.Atoi(friendID)

	friend, err := getFriendById(friendIDInt)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan ikut kelas Golang: %s",
			friend.Name,
			friend.Address,
			friend.Job,
			friend.Reason)
	}
}
