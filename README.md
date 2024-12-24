# Blog AggreGATOR

A command-line RSS feed aggregator named gator written in Go that allows users to manage and follow RSS feeds, collecting posts and storing them in a PostgreSQL database.

## Features

- User Management
  - Register new users
  - Login functionality
  - View all users
  - Reset user data

- Feed Management
  - Add new RSS feeds
  - View all feeds
  - Follow/unfollow feeds
  - View followed feeds
  - Browse posts from followed feeds
  - Automatic feed aggregation with configurable intervals

## Prerequisites

- Go 1.21+
- PostgreSQL
- Dependencies:
  - github.com/lib/pq
  - github.com/google/uuid

## Installation

1. Clone the repository:
```bash
git clone https://github.com/jamesfulreader/blog-aggreGATOR.git
cd blog-aggreGATOR
```

2. Install dependencies:
```bash
go mod download
```

3. Set up the PostgreSQL database and run migrations:
```bash
# Create database
createdb your_database_name

# Run migrations (using goose)
goose postgres "postgresql://username:password@localhost:5432/your_database_name?sslmode=disable" up
```

4. Create a configuration file at `~/.gatorconfig.json`:
```json
{
    "db_url": "postgresql://username:password@localhost:5432/your_database_name?sslmode=disable",
    "current_user_name": ""
}
```

## Usage

### Basic Commands

```bash
# Register a new user
./gator register <username>

# Login as a user
./gator login <username>

# View all users
./gator users

# Add a new feed
./gator addfeed <feed_name> <feed_url>

# List all feeds
./gator feeds

# Follow a feed
./gator follow <feed_url>

# View followed feeds
./gator following

# Unfollow a feed
./gator unfollow <feed_url>

# Browse posts (default limit: 2)
./gator browse [limit]

# Start feed aggregation (runs continuously)
./gator agg <time_between_requests>
```