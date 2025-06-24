package repository

import (
	"errors"
	"go-rest-api/internal/domain"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

// UserRepository is a struct that implements the UserRepository interface
type UserRepository struct {
	db    map[int]*domain.User // db is a map that simulates a database, where the key is the user ID and the value is a pointer to a User struct
	mutex sync.Mutex           // mutex is used to synchronize access to the db map
	ID    int                  // ID is used to simulate an auto-incrementing ID for new users
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db:    make(map[int]*domain.User),
		ID:    0,            // Initialize ID to 0
		mutex: sync.Mutex{}, // Initialize the mutex
	}
}

func (r *UserRepository) Create(user *domain.User) error {
	r.mutex.Lock()         // Lock the mutex to ensure thread-safe access to the db map
	defer r.mutex.Unlock() // Ensure the mutex is unlocked after the function completes

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost) // Hash the user's password
	if err != nil {
		return err
	}

	user.Password = string(hasedPassword) // Store the hashed password in the user struct

	user.ID = r.ID    // Set the ID of the user to the current value of ID
	r.db[r.ID] = user // Add the user to the map with the ID as the key
	r.ID++            // Increment the ID counter for the next user
	return nil
}

func (r *UserRepository) GetByID(id int) (*domain.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock() // Lock the mutex to ensure thread-safe access to the db map

	user, exist := r.db[id] // Look up the user by ID in the map
	if !exist {
		return nil, errors.New("user not found") // Return an error if the user is not found
	}

	return user, nil
}

func (r *UserRepository) GetByUsername(username string) (*domain.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, user := range r.db {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found") // Return an error if the user is not found
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	users := make([]domain.User, 0, len(r.db)) // Create a slice to hold all users
	for _, user := range r.db {
		users = append(users, *user) // Append each user to the slice
	}

	return users, nil
}

func (r *UserRepository) Update(user *domain.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exist := r.db[user.ID]; !exist {
		return errors.New("user not found") // Return an error if the user is not found
	}

	r.db[user.ID] = user // Update the user in the map
	return nil
}
