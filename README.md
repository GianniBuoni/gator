# 🐊 Gator

A simple cli blog aggregator!

## 💾 Installation & Dependencies

To install `gator`, you'll need Go installed on your system.  
With Go, install gator with this command:

```sh
go install "github.com/GianniBuoni/gator"
```

You'll also need a Postgress database to store all app data

## 🧾 Configuration

For the app to work, `gator` requires a `.gatorconfig.json` thag hold your database string and the name of the current user:

```json
{
  "db_url": "database/string?sslmode=disable",
  "current_user_name": "name"
}
```

## 🚀 Usage

List of commands, what they do and any needed args:

| COMMAND    | DESCRIPTION                                                                   | ARGUMENTS                                    |
| ---------- | ----------------------------------------------------------------------------- | -------------------------------------------- |
| `register` | Registers a name into the user database.                                      | 1: username (unique)                         |
| `login`    | Logs in a registered user.                                                    | 1: username                                  |
| `users`    | Lists all users.                                                              | None                                         |
| `addfeed`  | Adds a new feed into the database and follows that feed for the current user. | 1: feed name, 2: feed url                    |
| `feeds`    | Lists all added feeds.                                                        | None                                         |
| `follows`  | Lists current user's followed feeds.                                          | None                                         |
| `follow`   | Current users follows feed                                                    | 1: feed url (must be in database)            |
| `unfollow` | Current user removes feed froms following list.                               | 1: feed url                                  |
| `agg`      | Pulls post data from all feeds. Should be running in the background           | 1: refetch time (defaults to every 30s)      |
| `browse`   | Prints posts from feeds user follows                                          | 1: amount of posts displayed (defaults to 2) |
