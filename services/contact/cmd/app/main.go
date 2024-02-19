package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	repoContact "architecture_go/services/contact/internal/repository/contact"
	repoGroup "architecture_go/services/contact/internal/repository/group"
	ucContact "architecture_go/services/contact/internal/usecase/contact"
	ucGroup "architecture_go/services/contact/internal/usecase/group"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	dbInfo := &postgres.Database{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "0000",
		Name:     "golang",
	}

	db, err := postgres.OpenDB(dbInfo)
	if err != nil {
		panic(err.Error())
	}

	contactRepo := repoContact.NewContactRepository(db)
	groupRepo := repoGroup.NewGroupRepository(db)

	contactUC := ucContact.NewContactUseCase(contactRepo)
	groupUC := ucGroup.NewGroupUseCase(groupRepo)

	http.HandleFunc("/contact/create", func(w http.ResponseWriter, req *http.Request) {
		id, err := contactUC.CreateContact(contact.Contact{})
		if err != nil {
			panic("cannot create contact")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("new contact with id %d created", id)))
	})

	http.HandleFunc("/group/create", func(w http.ResponseWriter, req *http.Request) {
		id, err := groupUC.CreateGroup(group.Group{})
		if err != nil {
			panic("cannot create group")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("new group with id %d created", id)))
	})

	err = http.ListenAndServe("127.0.0.1:4000", nil)
	log.Fatal(err)
}
