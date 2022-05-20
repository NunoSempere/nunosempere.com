# How to create an autarchic website: the easy way

> Where to host the blog: Our Recommendation: Use Substack or try out a hosted Ghost website—it’s a Goldilocks solution of easy-to-use but highly customizable. Ghost makes for a great portfolio site, blog, and newsletter service.

[Effective Ideas: Where to host the blog](https://effectiveideas.org/how-to-start-a-blog/#pt3)

Fuck that. Vat do I vant, and in vat order? Uncensorability, ability to ramp up degrees of privacy and paranoia, control over ze whole stack, and finally, ze modern comforts. Vith that in mind, I present: Doctor Nefarious' guide to uncensorability.

What is this meant for? Some recent examples:

- [Youtube-dl being banned by Github](https://www.eff.org/deeplinks/2020/11/github-reinstates-youtube-dl-after-riaas-abuse-dmca)
- The Canadian government [fucking around with Canadian truckers' freedom of speech](https://twitter.com/punk6529/status/1494444624630403083)

## Get you a password manager for great security

For this, I recommend [pass](https://www.passwordstore.org/) (see also the [archlinux wiki](https://wiki.archlinux.org/title/Pass)). Simple, secure.

## Get an email

[Protonmail](https://protonmail.com/) is fine.

```
function newpassword(){
  characters="\!#\$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_\`abcdefghijklmnopqrstuvwxyz{|}~"
  length=25
  read -r -n $length new_password < <(LC_ALL=C tr -dc "$characters" < /dev/urandom)
  echo $new_password
} ## Taken from pass generate. Maybe save this in your .bashrc
newpassword
pass insert -m autarchy/protonmail ## Then save both the email address and the password.
## You can then recall your password with pass show autarchy/protonmail
```

Mad Eye Moody recommends: Don't give Protonmail a backup email. But make sure to make backups of your passwords. You can access Protonmail over [Tor](https://torproject.org/) [here](https://mail.protonmailrmez3lotccipshtkleegetolb73fuirgj7r4o4vfu7ozyd.onion). For previous steps, you can register an [XMPP identity](https://www.shad0w.io/). Better yet, do your own research on which services are the most secure.

## Register a domain

For zis, ve will use, of course, [njal.la](https://njal.la/). The're the non-evil twin brother of [Epik domains](https://www.epik.com/). From their [FAQ](https://njal.la/faq/):

> We're a team of committed internet activists and we're also involved in other privacy projects such as the IPredator VPN service. Some of us have also been involved in projects like The Pirate Bay and Piratbyrån to mention a few things.

> We support signing up for our service using email or XMPP (and yes, we use OTR). For email we also support PGP so all of our outgoing email will be signed and encrypted.

> We accept payments via Bitcoin, Litecoin, Monero, ZCash, DASH, Bitcoin Cash, Ethereum and Paypal.

They also have [a blog](https://njal.la/blog/) vere they record takedown requests:

![](https://njal.la/static/downloads/2019njetski/3.jpg)

All in all, zis is ze good shit.

For this, you will need to setup an account on njal.la, using the email you previously created

```
newpassword
pass insert -m autarchy/njalla-account ## save both password and user
## You can later recall this with pass show autarchy/njalla
```

Mad Eye Moody recommends: Use tornado cash (https://tornado.cash/) to anonymize your transaction history before paying njalla. Check which top-level domain names are more uncensorable (https://www.eff.org/files/2017/08/02/domain_registry_whitepaper.pdf). Ultimately, if we wanted further anonymity and uncensorability, we would [set up an onion service](https://community.torproject.org/onion-services/setup/).

## Step 2: Get a server

Previously, I have been using an Ubuntu server on [Digital Ocean](digitalocean.com), which offers servers from $5/month ($6/month including backups). But I have grown to dislike my system having 1GB worth of stuff that I don't understand. I also dislike wrangling with systemd services.

So I thought it would be fun to try Alpine Linux on [Njal.la's own servers](https://njal.la/servers/add/). This costs $15/month instead.

When buying a server from Njalla, we need to name it. And we will name it [`ramiel`](https://en.wikipedia.org/wiki/Ramiel) [^1]. We will also need to generate an ssh key to connect to it. The command to do this is `ssh-keygen -t ed25519`. After paying for our server, we'll want to set up payment to autorenew.

[^1]: Why Ramiel? Because Uriel is dead

## Step 3: Set up the server

Our setup will looks as follows:

![](https://i.imgur.com/69cnAw2.png)

We can connect to our server with

```
ssh root@your_ip
```

Then add some essential tooling:

```
apk add vim
apk add sudo
apk add git
```

Create a new user with root permissions

### Add a new user with root permissions

Strictly speaking you could do everything as root, but this feels dirty. Instead, we will create two new users, one root user for admin tasks, called `ramiel`, and one user for normal operation of web assets, named `www`. The below process probably contains some bugs, which stackoverflow or the alpine/arch linux wiki should solve.

#### Create new user:

```
adduser ramiel
apk add doas
adduser ramiel wheel
mkdir -p /home/ramiel
mkdir /home/ramiel/.ssh
```

#### Give it wide permissions and add it to the sudoers group

Add `permit persist :wheel` to `/etc/doas.d/doas.conf`

Then use `visudo` to edit `/etc/sudoers`, and add

```
ramiel ALL=(ALL) ALL
```

### Allow us to log in with ssh

```
cp /root/.ssh/authorized_keys /home/ramiel/.ssh/authorized_keys
chown ramiel /home/ramiel/.ssh/authorized_keys
```

In `/etc/ssh/sshd_config`, change:

```
PasswordAuthentication no
PermitRootLogin prohibit-password
```

to:

```
PasswordAuthentication no
PermitRootLogin yes
```

Then restart the ssh service

```
service sshd restart
exit
```

Connect with our new user

```
ssh ramiel@your_ip
```

## Install nginx

Here we make our first—of many—compromises with the suckless philosophy. Unlike simpler web servers, nginx is not suckless. However, I do have a job, and I'm choosing nginx because it has good integration with certbot, the certificates bot from the Electronic Freedom Foundation.

To install nginx, follow instructions from [here](https://wiki.alpinelinux.org/wiki/Nginx). In this case, I chose to create a folder in the `/home/www` directory, rather than directly in `/www`. This is probably some residual scaredness from when I didn't want to play around too much with directories outside the home directory.

```
apk update
apk add nginx
adduser -D -g 'www' www
chown -R www:www /var/lib/nginx
mkdir /home/www
chown -R www:www /home/www
```

We also want nginx to boot up at startup

```
rc-update add nginx default
```

## Install werc

[Werc](https://werc.cat-v.org/) is a work of genius by the mad architect Uriel. It describes itself as a "a minimalist web anti-framework built following the Unix and Plan 9 tool philosophy of software design". It is a more hardcore version of Jekyll, Pelican, Blogger, or Wordpress.

```
cd /home/www
mkdir werc
cd werc
wget --no-check-certificate https://werc.cat-v.org/download/werc-1.5.0.tar.gz
sha512sum werc-1.5.0.tar.gz
## 06695bf0798d28821500d84339f873a712ceb53cec6dc5bc641a47945a811515657e749c1d8c6cfeba42432ec8557397db9231ff4b5d2eec581a5ff063ab6648  werc-1.5.0.tar.gz
tar x -f werc-1.5.0.tar.gz
cd werc-1.5.0
```

Initialize a git repository

```
git init
git add .
git commit -m "feat: start keeping track of atomic changes on top of werc-1.5.0"
git remote add
```

It's good discipline to keep track of changes which are more or less atomic. This facilitates you remembering what changes you've done, and allows you to contribute back.

Alternatively, pull in my tweaks:

```
git init
git add .
git remote add tweaks https://github.com/NunoSempere/werc-1.5.0-tweaks.git
git pull tweaks master
```

These tweaks are a bit opinionated (e.g., they assume that you will set-up https). But they also make daily work more pleasant. Proceed at your own caution.

Mad Eye Moody recommends: Go through the [commit history](https://github.com/NunoSempere/werc-1.5.0-tweaks/commits/master) and pick and choose instead.

### Install plan9port

I might be able to get werc working with musl. But I can't be arsed. Instead, I am installing plan9port side by side.

Add the following line:

```
https://dl-cdn.alpinelinux.org/alpine/v3.15/community/
```

to the `/etc/apk/repositories` file, to get access to community packages.

Then install plan9port

```
apk install plan9port
## https://pkgs.alpinelinux.org/package/v3.15/community/x86_64/plan9port
```

By default, plan9port gets installed to `/usr/lib/plan9/bin`. But we want `rc` to be easily available:

```
cp /usr/lib/plan9/bin/rc /bin/rc
```

### Configure werc

There are several important configurations:

- Point to your plan9 binaries
- Whether to have a blog or not
- Whether to use http or https
- Whether to use [discount markdown](https://www.pell.portland.or.us/~orc/Code/discount/) or not, in which case you will have to install it

Of these, the more important is the first one. To let werc know where your plan9 binaries are installed

```
cd /home/www/werc/werc-1.5.0/etc/
cp initrc initrc.local
vim initrc.local
## Change one of the first lines to:
## plan9port=/usr/lib/plan9
```

## Configure nginx to work with werc

[Here](https://github.com/NunoSempere/werc-1.5.0-tweaks/tree/master/etc/contrib/nginx) is an example working configuration, i.e., the contents of my `/etc/nginx`

Here, I would recommend:

1. Get nginx up and running with a simple website using somethng like [this](https://github.com/NunoSempere/werc-1.5.0-tweaks/blob/master/etc/contrib/nginx/sites-available/nunosempere.com-simple)
2. Set it up to use werc with fastcgi
3. Add a DNS record pointing from your domain to your server
4. Add SSL/TSL certificates using certbot
5. Switch werc to https

Digital Ocean has a few good tutorials on how to get nginx up and running:

- [How To Install Nginx on Ubuntu 20.04](https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-ubuntu-20-04)
- [How To Secure Nginx with Let's Encrypt on Ubuntu 20.04](https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-20-04)

Steps are very similar on Alpine. Though, e.g., Alpine doesn't have the `ufw` firewall, and so on.

Before you install SSL, you will need a DNS record, and in particular, an A record, pointing from your domain to your server. It might take a while for the DNS record to be updated; around 24 to 48 hours.

The commands for installing certificates are pretty trivial:

```
apk add --update python3 py3-pip
apk add certbot
pip install certbot-nginx
rc-service nginx stop
sudo certbot --nginx -d your-site.org
```

```
sudo apk add spawn-fcgi
sudo apk add fcgiwrap

```

### Installing fcgi

```
sudo apk add spawn-fcgi
sudo apk add fcgiwrap
```

You can then spawn this with

```
/usr/bin/spawn-fcgi -a 127.0.0.1 -p 9000 -f /usr/bin/fcgiwrap
```

Alternatively, and more conveniently, you can set this up so that fcgi is spawned when the system starts. To do so, copy the following:

```
#!/sbin/openrc-run
## move to /etc/init.d/spawn-fcgi2
## Then sudo rc-update add spawn-fcgi2
## test with sudo rc-service spawn-fcgi2 start
## and sudo rc-service spawn-fcgi2 stop
PID_file="/var/tmp/spawn"

depend() {
  need net
}

start() {
  /usr/bin/spawn-fcgi -a 127.0.0.1 -p 9000 -f /usr/bin/fcgiwrap -P "$PID_file"
}

stop() {
  kill -15 "$(cat "$PID_file")"
  rm "$PID_file"
}

case $1 in
  start|stop) "$1" ;;
esac
```

into `spawn-fcgi2`, and then:

```
mv spawn-fcgi2 /etc/init.d/spawn-fcgi2
cd /etc/init.d
sudo chmod 755 spawn-fcgi2
sudo rc-update add spawn-fcgi2
```

## Install discount markdown

## Debug problems
