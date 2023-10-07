# Meme saviour

<img src="https://external-content.duckduckgo.com/iu/?u=http%3A%2F%2Fclipart-library.com%2Fimages_k%2Fangry-pepe-transparent%2Fangry-pepe-transparent-11.png" alt="pepe the frog" width="100"/>

A Telegram bot that saves memes (or any other image really) to a [provided storage](https://github.com/contre95/meme-saviour/blob/327baffd78cf1be353e79b3dc376f3be46f9e352/app/domain.go#L11-L16).

https://github.com/contre95/meme-saviour/blob/327baffd78cf1be353e79b3dc376f3be46f9e352/app/domain.go#L11-L16

## Configuration ⚙️

Configuration can be set using env variables. Storage variables are optional but at least one Storage must be for the app to work. Currently only Local storage is implemented therefore must be set.

* `TELEGRAM_TOKEN` (required): Token of the Telegram bot you've created.
* `TELEGRAM_ALLOWED_USERNAMES (required)`: CSV of usernames that are allowed to use the bot (*eg. user1,user2,usern*)
* `LOCAL_STORAGE` (optional): if set to `1` then is enabled (default: `0`)
* `LOCAL_STORAGE_PATH` (optional): Set to the path where you want to store the memes (*eg. /data/memes/*, default: `/data`)

# Run

Run manually from for testing.
```shell

```

Run it on a container.

```shell
podman container run --rm \
	--name memesaviour
	-v "$(pwd)/data:/data" \
	-e TELEGRAM_TOKEN=<telegram-bot-token> \
	-e LOCAL_STORAGE=1 \
	-e TELEGRAM_ALLOWED_USERNAMES=user1,user2,user3 \
	docker.io/contre95/memesaviour
```
