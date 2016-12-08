# The Things Network Wiki

This is the content of The Things Network's wiki.

> Note: The wiki rewrite is still work in progress, so make sure you check out the branch `rewrite`.  
> You should also add the `--ref rewrite` flag when running gollum.

## Running Locally (simple version)

* Clone the repository [github.com/TheThingsNetwork/wiki](https://github.com/TheThingsNetwork/wiki)
* Install [gollum](https://github.com/gollum/gollum/wiki/Installation)
* Run `gollum`.
  * While on rewrite branch, run with `gollum -r rewrite`.
* Go to `localhost:4567`

## Running Locally (with TTN styles)

* Clone the **wiki** repository [github.com/TheThingsNetwork/wiki](https://github.com/TheThingsNetwork/wiki)
* Clone the **gollum** repository [github.com/TheThingsNetwork/gollum](https://github.com/TheThingsNetwork/gollum)
* In the location of the **gollum** repository, run
  * `bundle install` (if something goes wrong, you might miss some [dependencies](https://github.com/gollum/gollum/wiki/Installation))
  * `export ACCOUNT_SERVER_URL=https://account.thethingsnetwork.org`
  * `bundle exec bin/gollum ../wiki --config config.rb` (where `../wiki` points to the location of the **wiki** repository)
    * While on rewrite branch, run with `bundle exec bin/gollum ../wiki --config config.rb --ref rewrite`.
* Go to `localhost:4567`
