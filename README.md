## ON Esports Discord Bot - Go Developer Task

---

### The task

Develop a Go-based Discord bot with unique features to engage users.

### Requirements

- **Language:** Go (Golang)
- **Libraries:** DiscordGo (or similar)
- **Discord API:** Understand bot creation and API usage
- **Version Control:** Git (GitHub)

### Core Features

- Command parsing and response
- Help command listing available functions
- Asynchronous task processing

### Unique Features to add

- Poll creation with voting
- Weather updates by location
- Message translation
- Reminder system
- Simple text-based game
- Showcase your creativity and Go skills!


## Meet Anatoliy

---

**Anatoliy** is a gamer and cybersports enthusiast. He loves to play games a lot, but more than anything, he loves to prompt chatting game results and schedules. Also, when Anatoly is not playing games, he works part-time as a bookmaker and will be able to take your bet on a game. But will he return your winnings...

The bot intends to provide a variety of features to help users stay up to date with the latest esports news, game and players statistics, and provide some gamification in the form of "bets" to upcoming games.

### Anatoliy's Features

- `/help` - List of available commands and their descriptions.
- `/quote` - Get a random quote from Anatoliy.
- `/on-esports` - Gives a quote about **ON Esports**.

**Features that could'n be implemented, but in the process of development**:

- TODO: Add command to check latest news of given discipline
- TODO: Add command to check given player’s avg. statistics
- TODO: Add command to check given player's performance on given match
- TODO: Add command to check past game result of given team in the given discipline
- TODO: Add command to get upcoming games and tournaments of specified team or discipline.
- TODO: Add command to get H2H (Head-to-Head) stats of two commands
- TODO: Add "gamification": to implement coins
- TODO: Add command to create a bet on specified match using those coins
- TODO: Add command to get user's rating based on their balance / bet winning rate etc.


## How to run

---

1. Clone the repository
2. Create a `.env` file in the root of the project and add the following:  

    ```env
    DISCORD_BOT_TOKEN=<YOUR_DISCORD_BOT_TOKEN>
    ```

3. Run the bot using the following command:

    ```bash
    go run main.go
    ```

OR

Just use the [link](https://discord.com/api/oauth2/authorize?client_id=1203244936498257930&permissions=8&scope=bot) to add the bot to your server!
