# Shortbot
The official bot of the Short Gang's Discord server. What does it do?

Made for fun with no expectations. Anyone can contribute :)

## How to run the bot/contribute
Software: [Go](https://go.dev/) and [Git](https://git-scm.com/)
1. Clone the repository (fork it if you want to contribute).
2. Make a bot via the Discord developer portal and copy your bot's token.
3. Create a ".env" file from the ".env.example" and paste the token in the ".env" file.
4. Run with `go run .`.
5. Make changes.
6. Create a pull request.
7. Profit?

## Specific Contributions
### Adding a command
Create the command in  [/commands/actions](/commands/actions). Commands are of type `func(*discordgo.Session, *discordgo.InteractionCreate)`. Then add your command to BotCommands in [/commands/registercommands.go](/commands/registercommands.go). 
### Adding an automatic response
Create the response in [/responses](responses). Responses are of type `func(s *discordgo.Session, m *discordgo.MessageCreate)`. Then add your response in [/eventhandlers/responsemanager.go](eventhandlers/responsemanager.go).
