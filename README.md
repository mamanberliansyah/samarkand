# samarkand
Basic Web Server Written in GoLang. Pretty much a lightweight pseudo alternative to NGINX and Apache. Because it's my first shot on this language, don't expect much out of it.

### Installation :
All you have to do is run the install.sh script and let the magic happen.

### Configuration :
Configuration can be accessed in :<br>`/etc/samarkand/config.toml` <br>
There are two things you can configure, port and webpage directory, both are pretty self explanatory. Default port is 80 and default directory is : `/var/www/html`

### Uninstall :
I know some of you may get tired because of my program, so I've also made an exit. To uninstall, just run the uninstall.sh script.

### Sources & Libraries :
GoLang Web Application Tutorial by Soham Kamani<br>
`https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/`<br>
URL Router and Dispatcher by Gorilla<br>
`https://github.com/gorilla/mux`<br>
TOML Parser by BurntSushi<br>
`https://github.com/BurntSushi/toml`<br>
