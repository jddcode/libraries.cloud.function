package cloudFunction

import (
	FireStore "cloud.google.com/go/firestore"
	"context"
	"errors"
	"os"
	"strings"
)

	type FireStoreClient struct {

		client *FireStore.Client
		setup bool
	}

	func (fs *FireStoreClient) doSetup() error {

		if fs.setup {

			return nil
		}

		var err error
		fs.client, err = FireStore.NewClient(context.Background(), os.Getenv("GCP-Project"))

		if err != nil {

			return err
		}

		fs.setup = true
		return nil
	}

	func (fs *FireStoreClient) Write(path string, object interface{}) error {

		err := fs.doSetup()

		if err != nil {

			return err
		}

		// Cut this up; the last part is the name of the item
		bits := strings.Split(strings.TrimRight(path, "/"), "/")

		if len(bits) < 2 {

			return errors.New("You must supply a path and filename")
		}

		filename := bits[len(bits) - 1]

		_, err = fs.getCollection(strings.Join(bits[0:len(bits) - 1], "/")).Doc(filename).Set(context.Background(), object)
		return err
	}

	func (fs *FireStoreClient) Delete(filename string)  error {

		err := fs.doSetup()

		if err != nil {

			return err
		}

		// Cut this up; the last part is the name of the item
		bits := strings.Split(strings.TrimRight(filename, "/"), "/")

		if len(bits) < 2 {

			return errors.New("You must supply a path and filename")
		}

		actualFile := bits[len(bits) - 1]

		_, err = fs.getCollection(strings.Join(bits[0:len(bits) - 1], "/")).Doc(actualFile).Delete(context.Background())
		return err
	}

	func (fs *FireStoreClient) Search(path, field, operator string, filter interface{}) ([]*FireStore.DocumentSnapshot, error) {

		err := fs.doSetup()

		if err != nil {

			return nil, err
		}

		return fs.getCollection(path).Where(field, operator, filter).Documents(context.Background()).GetAll()
	}

	func (fs *FireStoreClient) SearchOne(path, field, operator string, filter, result interface{}) error {

		err := fs.doSetup()

		if err != nil {

			return err
		}

		results, err := fs.Search(path, field, operator, filter)

		if err != nil {

			return err
		}

		if len(results) < 1 {

			return errors.New("No results for your search")
		}

		results[0].DataTo(&result)
		return nil
	}

	func (fs *FireStoreClient) Fetch(filename string, destination interface{}) error {

		err := fs.doSetup()

		if err != nil {

			return err
		}

		// Cut this up; the last part is the name of the item
		bits := strings.Split(strings.TrimRight(filename, "/"), "/")

		if len(bits) < 2 {

			return errors.New("You must supply a path and filename")
		}

		actualFile := bits[len(bits) - 1]

		result, err := fs.getCollection(strings.Join(bits[0:len(bits) - 1], "/")).Doc(actualFile).Get(context.Background())

		if err == nil {

			err = result.DataTo(&destination)
		}

		return err
	}

	func (fs *FireStoreClient) getCollection(path string) *FireStore.CollectionRef {

		err := fs.doSetup()

		if err != nil {

			return nil
		}

		bits := strings.Split(strings.TrimRight(path, "/"), "/")

		if len(bits) < 2 {

			return fs.client.Collection(bits[0])
		}

		if len(bits)%2 == 0 {

			bits = append(bits, "items")
		}

		return fs.client.Collection(strings.Join(bits, "/"))
	}
