<h1 align="center">gospotify</h1>
<h3 align="center">A simple Golang-based Spotify web client</h3>

## Overview

This project is a simple Golang-based web client for interacting with the Spotify API. It provides basic functionalities for authenticating with Spotify, retrieving access tokens, and accessing Spotify data.

## Features

- Authenticate with Spotify
- Retrieve access tokens
- Access Spotify data
- Access Spotify data with custom scopes

## Prerequisites

- Go 1.18 or later
- A Spotify Developer account
- Spotify API credentials (Client ID, Client Secret, Redirect URI)

Visit Spotify [Web API](https://developer.spotify.com/documentation/web-api) documentation to get started.

## Installation

1. **Clone the Repository**

```bash
git clone https://github.com/alicse3/gospotify.git
```

2. **Install Dependencies**
```bash
go mod tidy
```

## Configuration

Before running the client, set up your Spotify API credentials. You can set these as environment variables or configure them directly in the code.

```bash
export SPOTIFY_CLIENT_ID="your_client_id"
export SPOTIFY_CLIENT_SECRET="your_client_secret"
export SPOTIFY_REDIRECT_URL="your_redirect_uri"
```

## Usage
1. **Run the Web Client**
```bash
go run main.go
```

2. **Authorize the Application**

App tries to open your default browser for the authorization. Follow the instructions to authenticate with Spotify and access your data.

## Examples

### Basic Initialization with `NewClient`

The `NewClient` function initializes the client with default settings and no specific scopes.

Here's an example of how to use the `NewClient` to authenticate and retrieve user data:
```go
package main

import (
    "log"
    "github.com/alicse3/gospotify"
)

func main() {
	// Initialize the client with your Spotify credentials
	client, err := gospotify.NewClient(
		&gospotify.Credentials{
			ClientId:     "your_client_id",
			ClientSecret: "your_client_secret",
			RedirectUrl:  "your_redirect_uri",
		},
	)
	if err != nil {
		log.Fatalf("Failed to initialize client: %v", err)
	}

	// Retrieve the current user's profile data
	userData, err := client.UserService.GetCurrentUserProfile
	if err != nil {
		log.Fatalf("Failed to get user data: %v", err)
	}

	// Print the user data
	log.Printf("User data: %v", userData)
}
```

### Basic Initialization with `DefaultClient`

If you have already exported credentials to environment variables, you can use the `DefaultClient` to authenticate:

```go
package main

import (
    "log"
    "github.com/alicse3/gospotify"
)

func main() {
	// Initialize the client using default credentials from environment variables
	client, err := gospotify.DefaultClient()
	if err != nil {
		log.Fatalf("Failed to create default client: %v", err)
	}

	// Perform additional operations using the client
	// ...
}
```

### Initialization with Custom Scopes using `NewClientWithCustomScopes`

The `NewClientWithCustomScopes` function initializes the client with user-specified scopes, allowing you to request only the permissions necessary for your application's functionality.

Here's an example of how to use the `NewClientWithCustomScopes`:
```go
package main

import (
    "log"
    "github.com/alicse3/gospotify"
)

func main() {
	// List of custom scopes
	customScopes := []string{gospotify.ScopeUserPrivateRead, gospotify.ScopeUserReadEmail}

	// Initialize the client with custom scopes
	client, err := gospotify.NewClientWithCustomScopes(
		&gospotify.Credentials{
			ClientId:     "your_client_id",
			ClientSecret: "your_client_secret",
			RedirectUrl:  "your_redirect_uri",
		},
        customScopes, // Passing custom scopes here
	)

	// Perform additional operations using the client
	// ...
}
```

### Initialization with token using `NewClientWithToken`

The `NewClientWithToken` function initializes the client with user-specified token.

Here's an example of how to use the `NewClientWithToken`:
```go
package main

import (
    "log"
    "github.com/alicse3/gospotify"
)

func main() {
	// Token to use
	token := "xxxxx"

	// Initialize the client with token
	client, err := gospotify.NewClientWithToken(token)
	if err != nil {
		log.Fatalf("Failed to initialize client with token: %v", err)
	}

	// Perform additional operations using the client
	// ...
}
```

## Testing
There are currently no tests written for this project. Contributions for adding tests are welcome and highly encouraged!

## Contributing
Contributions are welcome!

Please fork the repository and submit a pull request with your changes. Make sure to follow the coding conventions and write tests for any new features or bug fixes.

## License
This project is licensed under the MIT License. See the `LICENSE` file in the project root directory for more details.

## Contact
For any questions or support, please contact alicse3@gmail.com.
