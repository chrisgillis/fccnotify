## fccnotify

fccnotify helps new hams know when their callsign is ready


### usage

fccnotify -frn <your frn>

if you want an e-mail notification you'll need a gmail account

turn on access for less secure apps in your gmail settings:
https://www.google.com/settings/u/1/security/lesssecureapps

fccnotify -frn <your frn> -gmailaddr <gmail email address> -gmailpass <gmail password>

### interval

the command will check the fcc database every 30 minutes by default. you can increase the time interval with `-m <minutes>` but it cannot be less than 30. It's pointless to check more quickly than that anyway.

### License

MIT License