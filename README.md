# Gift Mini App
The Gift Redemption and Rating Application is a web-based platform that allows users to log in, redeem gifts, and rate them. This application provides an interactive and engaging experience for users to access and utilize a variety of gift options while providing feedback through ratings.

## Features
- User Authentication: Create an account or log in securely.
- Gift Redemption: Browse, select, and redeem available gifts.
- Rating System: Rate and provide feedback on redeemed gifts.
- Role-Based Access Control: Admins manage the gift catalog.
- Token-Based Authentication: Secure API endpoints.
## Tech Specifications

### Libraries/Frameworks Used

- Go (version 1.19)
- Gorilla Mux
- PostgreSQL

### Architecture/Modularity

The project follows a clean architecture pattern, separating concerns into different layers:

- **Presentation Layer**: Contains the API handlers and controllers.
- **Service Layer**: Contains the business logic.
- **Repository Layer**: Deals with data storage and retrieval.
- **Database Layer**: Connects to the database.
- **Test Layer**: Contains unit tests and integration tests.

## Quick Start
### Installation guide
#### 1. install go version 1.19
```bash
# Please Read This Link Installation Guide of GO

# Link Download -> https://go.dev/dl/
# Link Install -> https://go.dev/doc/install

```

#### 2. Run the application
```bash
# run command :
git clone https://Dito_Adriel@bitbucket.org/Bhenedicto_Adriel/dito-rgb-golang-test.git

# install dependency
go mod tidy

# setup env
DB_DRIVER=postgres
DB_USERNAME=        #change to your db username
DB_PASSWORD=        #change to your db password
DB_HOST=            #change to your db host
DB_PORT=            #change to your db port 
DB_DATABASE=        #change to your db name 
DB_URL=             #postgres://{DB_USERNAME}:{DB_PASSWORD}@{DB_HOST}:{DB_PORT}/{DB_DATABASE}?sslmode=disable

KEY=                #change to your key
EXPIRED=            #change to your expiration time
# Run App
make start

# Migrate db
make migrate-up //this for up migrations
make migrate-down //this for down migrations
```