package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`    // Username is required for user identification
	Password string `json:"password" binding:"required"`    // Password is required for authentication
	Email    string `json:"email" binding:"required,email"` // Email is required and must be a valid email format
	Role     string `json:"role" binding:"required"`        // Role is required to define user permissions
}

type AuthService interface {
	Login(username, password string) (string, error) // Login method to authenticate a user and return a token
	Logout(token string) error                       // Logout method to invalidate a user's token
	ValidateToken(token string) (string, error)      // ValidateToken method to check if a token is valid and return the associated username
}

type UserRepository interface {
	Create(user *User) error                      // Create a new user in the repository
	GetByID(id int) (*User, error)                // Get a user by ID from the repository
	GetByUsername(username string) (*User, error) // Get a user by username from the repository
	GetAll() ([]User, error)                      // Get all users from the repository
	Update(user *User) error                      // Update an existing user in the repository
	Delete(id int) error                          // Delete a user by ID from the repository
}

type UserUsecase interface {
	Create(user *User) error                         // Create a new user using the use case
	GetByID(id int) (*User, error)                   // Get a user by ID using the use case
	GetByUsername(username string) (*User, error)    // Get a user by username using the use case
	GetAll() ([]User, error)                         // Get all users using the use case
	Update(user *User) error                         // Update an existing user using the use case
	Delete(id int) error                             // Delete a user by ID using the use case
	Login(username, password string) (string, error) // Login a user and return a token
	Logout(token string) error                       // Logout a user and invalidate the token
}
