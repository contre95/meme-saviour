# Meme saviour

<img src="https://external-content.duckduckgo.com/iu/?u=http%3A%2F%2Fclipart-library.com%2Fimages_k%2Fangry-pepe-transparent%2Fangry-pepe-transparent-11.png" alt="pepe the frog" width="100"/>

A Telegram bot that saves memes (or any other image really) to a [provided storage](https://github.com/contre95/meme-saviour/blob/327baffd78cf1be353e79b3dc376f3be46f9e352/app/domain.go#L11-L16).

https://github.com/contre95/meme-saviour/blob/327baffd78cf1be353e79b3dc376f3be46f9e352/app/domain.go#L11-L16

# Configuration ⚙️

Configuration can be set using env variables. There's no validation for config value. ([see code](https://github.com/contre95/meme-saviour/blob/327baffd78cf1be353e79b3dc376f3be46f9e352/main.go#L17-L18))

* `BOT_MEMESAVE_TELEGRAM_TOKEN`: Token of the Telegram bot you've created.
* `BOT_ALLOWED_USERNAMES`: CSV of usernames that are allowed to use the bot (eg. user1,user2,usern)
