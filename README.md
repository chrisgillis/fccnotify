## fccnotify

fccnotify helps new hams know when their callsign is ready

### Installation

I recommend installing via docker.

#### via Docker

```
docker run -d -e="frn=x" -e="gmailaddr=x" -e "gmailpass=x" --restart always gillisct/fccnotify
```

#### via Source

You must have Golang installed on your system.

```
git clone https://github.com/chrisgillis/fccnotify.git
cd fccnotify
go install
fccnotify -frn <your frn>
```

### Usage

`fccnotify -frn <your frn> [-m <minutes>] [-gmailaddr] [-gmailpass]`

* `-frn`, _required_, your FCC FRN
* `-m`, _optional_, check the fcc database every x minutes, minimum is 30
* `-gmailaddr`, _optional_, your gmail e-mail address
* `-gmailpass`, _optional_, your gmail password

Run the command without the gmail options and it will log to console when your callsign is ready.

If you want it to notify you via e-mail, you'll need a [Gmail](http://www.gmail.com/) account.

Make sure you [turn on access for less secure apps](https://www.google.com/settings/u/1/security/lesssecureapps) in your gmail settings.

### License

MIT License
