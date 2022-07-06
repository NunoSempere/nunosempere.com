# How to create an autarchic website

> Where to host the blog: Our Recommendation: Use Substack or try out a hosted Ghost website—it’s a Goldilocks solution of easy-to-use but highly customizable. Ghost makes for a great portfolio site, blog, and newsletter service.

[Effective Ideas: Where to host the blog](https://effectiveideas.org/how-to-start-a-blog/#pt3)

No. Vat do I vant, and in vat order? Uncensorability, ability to ramp up degrees of privacy and paranoia, control over ze whole stack, and finally, ze modern comforts. Vith that in mind, I present my guide to uncensorability.

What is this meant for? Some recent examples:

- The Canadian government [grossly overreached against Canadian truckers' freedom of speech](https://twitter.com/punk6529/status/1494444624630403083).
- [Youtube-dl was banned by Github](https://www.eff.org/deeplinks/2020/11/github-reinstates-youtube-dl-after-riaas-abuse-dmca).
- [EFF Sues Proctorio on Behalf of Student It Falsely Accused of Copyright Infringement to Get Critical Tweets Taken Down](https://www.eff.org/press/releases/eff-sues-proctorio-behalf-student-it-falsely-accused-copyright-infringement-get).

## Index

1. Preliminiaries

- Introduction (^)
- Get a password manager
- Get an email

2. Get a server

- Register a domain
- Get a server
- Point your domain to your server

3. Set up a simple server

- Create and set up users
- Install nginx
  - Get nginx working with a simple website
  - Install TLS certificates

4. Into the rabbit hole

- Install plan9port
- Install werc
- Configure werc
  - Install discount markdown
- Configure nginx to work with werc
  - Install fcgi

## 1. Preliminaries

### Get you a password manager for great security

For this, I recommend [pass](https://www.passwordstore.org/) (see also the [archlinux wiki](https://wiki.archlinux.org/title/Pass)). Simple, secure.

To generate a password independent from pass:

```
function newpassword(){
  characters="\!#\$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_\abcdefghijklmnopqrstuvwxyz{|}~"
  # ^ excludes `
  length=25
  read -r -n $length new_password < <(LC_ALL=C tr -dc "$characters" < /dev/urandom)
  echo $new_password
} ## Taken from pass generate. Maybe save this in your .bashrc
newpassword
pass insert -m autarchy/protonmail ## Then save both the email address and the password.
## You can then recall your password with pass show autarchy/protonmail
```

Personally, I've found it useful to wrap `pass insert -m` into an extension for pass (pass [append]()). I also found it useful to have a utility for [searching passwords](). And, as always, [backup](https://github.com/8go/pass-backup) stuff.

### Get an email

[Protonmail](https://protonmail.com/) is fine.

Mad Eye Moody recommends: Don't give Protonmail a backup email. Instead, make sure to make backups of your passwords. You can access Protonmail over [Tor](https://torproject.org/) [here](https://mail.protonmailrmez3lotccipshtkleegetolb73fuirgj7r4o4vfu7ozyd.onion). For added paranoia, you can register an [XMPP identity](https://www.shad0w.io/) instead of an email address. Also, do your own research on which service providers are the most secure.

## 2. Get a server

### Register a domain

For zis, ve will use, of course, [njal.la](https://njal.la/). The're the non-evil twin brother of [Epik domains](https://www.epik.com/). From their [FAQ](https://njal.la/faq/):

> We're a team of committed internet activists and we're also involved in other privacy projects such as the IPredator VPN service. Some of us have also been involved in projects like The Pirate Bay and Piratbyrån to mention a few things.

> We support signing up for our service using email or XMPP (and yes, we use OTR). For email we also support PGP so all of our outgoing email will be signed and encrypted.

> We accept payments via Bitcoin, Litecoin, Monero, ZCash, DASH, Bitcoin Cash, Ethereum and Paypal.

They also have [a blog](https://njal.la/blog/) vere they record takedown requests. All in all, zis is ze good stuff.

For this, you will need to setup an account on njal.la, using the email you previously created

```
newpassword
pass insert -m autarchy/njalla-account ## save both password and user
## You can later recall this with pass show autarchy/njalla,
## or with pass reveal njalla if you installed that extension
```

Mad Eye Moody recommends: Use [tornado cash](https://tornado.cash/) to anonymize your transaction history before paying njalla. Check which top-level domain names are [more uncensorable](https://www.eff.org/files/2017/08/02/domain_registry_whitepaper.pdf) (from casual reading, I would go with .br, .cr , .is, .org, .ru, of which njalla supports only .org and .ru). If if we wanted further anonymity and uncensorability, [set up an onion service](https://community.torproject.org/onion-services/setup/).

### Get a server

Previously, I had been using an Ubuntu server on [Digital Ocean](digitalocean.com), which offers servers from $5/month ($6/month including backups). But I have grown to dislike my system having 1GB worth of stuff that I don't understand. I also dislike wrangling with systemd services.

So I thought it would be fun to try Alpine Linux on [Njal.la's own servers](https://njal.la/servers/add/). This costs $15/month instead.

When buying a server from Njalla, we need to name it. And we will name it [`ramiel`](https://en.wikipedia.org/wiki/Ramiel) [^1]. We will also need to generate an ssh key to connect to it. The command to do this is `ssh-keygen -t ed25519`. After paying for our server, we'll want to set up payment to autorenew.

[^1]: Why Ramiel? Because Uriel is dead

### Point from your domain to your server.

Add a DNS record pointing from your domain to your server. In particular, you want an A record. Note that it might take a while for it to propagate, so at this point, you can either:

- Wait for ~24 until your server and domain name are connected
- Set up your users and a minimal nginx installation. This requires some additional fiddling, and I will not document it here.

## Step 3: Set up the server

### Set up a new user

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

Strictly speaking you could do everything as root, but this feels dirty. Instead, we will create two new users, one root user for admin tasks, called `ramiel`, and one user for normal operation of web assets, named `www`. The below process probably contains some bugs, which stackoverflow or the alpine and arch linux wikis should solve.

To create a new root user:

```
adduser ramiel
apk add doas
adduser ramiel wheel
mkdir -p /home/ramiel
mkdir /home/ramiel/.ssh
```

Then, add `permit persist :wheel` to `/etc/doas.d/doas.conf`

Then use [visudo](https://unix.stackexchange.com/questions/27594/why-do-we-need-to-use-visudo-instead-of-directly-modifying-the-sudoers-file) to add `ramiel ALL=(ALL) ALL` to `/etc/sudoers`.

Lastly, allow yourself to log into the server as this new user using ssh.

```
cp /root/.ssh/authorized_keys /home/ramiel/.ssh/authorized_keys
chown ramiel /home/ramiel/.ssh/authorized_keys
```

Then, in `/etc/ssh/sshd_config`, change:

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

And connect with your new user

```
ssh ramiel@your_ip
```

### Install nginx

Here we make our first—of many—compromises with the [suckless](https://suckless.org/) philosophy. Unlike simpler web servers, nginx—pronounced "engine x"—is not suckless. However, I do have a job, and I'm choosing nginx because it is well documented, widely used, and has good integration with certbot, the certificates bot from the Electronic Freedom Foundation which allows me to have https (also perhaps not suckless).

To install nginx, follow instructions from [here](https://wiki.alpinelinux.org/wiki/Nginx). Digital Ocean has a few good tutorials on how to get nginx up and running, e.g., [How To Install Nginx on Ubuntu 20.04](https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-ubuntu-20-04). Steps are very similar on Alpine. Though Alpine doesn't have the `ufw` firewall, the package manager is called `apk`, not `apt`, and so on.

In this case, I chose to create a folder in the `/home/www` directory, rather than directly in `/www`. This is probably some residual scaredness from when I didn't want to play around too much with directories outside the home directory.

In any case, the commands I used to set up nginx were:

```
apk update
apk add nginx
adduser -D -g 'www' www
chown -R www:www /var/lib/nginx ## This is important!!
mkdir /home/www
chown -R www:www /home/www
```

We also want nginx to boot up at startup

```
rc-update add nginx default
rc-service nginx status

```

At this point, we can get a minimal website working. Create an `index.html` file somewhere, and use a nginx setup such as [this](https://github.com/NunoSempere/werc-1.5.0-tweaks/blob/master/etc/contrib/nginx/sites-available/nunosempere.com-simple) to check everything is working out ok. This is an important step, because it lets you check your work so far.

After some debugging, you should have a very simple yet fairly hard to censor website. Congratulations!

### Set up TLS

At this point, one might as well add TLS working, by following instructions such as those in [How To Secure Nginx with Let's Encrypt on Ubuntu 20.04](https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-20-04).

Before you install SSL, you will need a DNS record, and in particular, an A record, pointing from your domain to your server. It might take a while for the DNS record to be updated; around 24 to 48 hours.

The commands for installing certificates are pretty trivial:

```
apk add --update python3 py3-pip
apk add certbot
pip install certbot-nginx
rc-service nginx stop
sudo certbot --nginx -d your-site.org
```

## 4. Into the rabbit hole

Our setup will looks as follows:

![](https://i.imgur.com/69cnAw2.png)

And the steps we will use to get this working are:

- Install plan9port
- Install werc
- Configure werc
  - Install discount markdown
- Configure nginx to work with werc
- Install fcgi

Why am I using werc in Alpine linux rather than, I don't know, [Pelican](https://blog.getpelican.com/) on Debian, or [Jekyll](https://jekyllrb.com/) on Github Pages or [Gatsby](https://www.gatsbyjs.com/) on netlify? Because I am a sadomasochist, that's why. Proba

### Install plan9port

Plan9port is werc's key dependency. I might be able to get werc working with musl. But I can't be arsed. Instead, I am installing plan9port side by side with musl.

To install plan9port, add the following line:

```
https://dl-cdn.alpinelinux.org/alpine/v3.15/community/
```

to the `/etc/apk/repositories` file, to get access to community packages.

Then install plan9port

```
apk install plan9port
## https://pkgs.alpinelinux.org/package/v3.15/community/x86_64/plan9port
```

By default, plan9port gets installed to `/usr/lib/plan9/bin`. But we want the `rc` shell to be easily available:

```
cp /usr/lib/plan9/bin/rc /bin/rc
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

Alternatively, pull in _my own personal tweaks_:

```
git init
git add .
git remote add tweaks https://github.com/NunoSempere/werc-1.5.0-tweaks.git
git pull tweaks master
```

These tweaks are a bit opinionated (e.g., they assume that you have set-up https). But they also make daily work more pleasant. Proceed at your own caution.

Mad Eye Moody recommends: Go through the [commit history](https://github.com/NunoSempere/werc-1.5.0-tweaks/commits/master) and pick and choose instead.

### Configure werc

There are several important configurations:

- Point to your plan9 binaries
- Decide whether to have a blog or not
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

To enable blog functionality, follow instructions [here](https://werc.cat-v.org/apps/blagh/)

To use https instead of http, grep for http and change it all to https. Find the necessary files using `grep http -R bin/`, or have a look at [this git commit](https://github.com/NunoSempere/werc-1.5.0-tweaks/commit/18893f3192721144e82497dce751864faa4437e4)

Personally, I'm a fan of discount markdown, because it has a few more features than the default awk markdown parser. To install discount markdown, run

```
apk install discount
```

And then have a look at [this commit](https://github.com/NunoSempere/werc-1.5.0-tweaks/commit/f492aa962b4c0627dd013f074c0d01000f6e70fe) for the necessary changes.

## Configure nginx to work with werc

Configuring nginx to work with werc is the tricky part. Luckily, [here](https://github.com/NunoSempere/werc-1.5.0-tweaks/tree/master/etc/contrib/nginx) is an example working configuration, i.e., the contents of my `/etc/nginx`. Best of luck!

One of the key pieces to connect nginx and werc is [fcgi](https://wikiless.org/wiki/FastCGI?lang=en) (fast common gateway interface). It can be installed with

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

into a new file, `spawn-fcgi2`, and then:

```
mv spawn-fcgi2 /etc/init.d/spawn-fcgi2
cd /etc/init.d
sudo chmod 755 spawn-fcgi2
sudo rc-update add spawn-fcgi2
```

Installing fcgi, copying [my configuration](https://github.com/NunoSempere/werc-1.5.0-tweaks/tree/master/etc/contrib/nginx) and making the relevant changes probably still requires some tweaking and debugging. For instance, it's necessary that the nginx folder is owned by the `www` user. If someone wants to document this better, I welcome improvements.

## Backups

Backups, and then backups for the backups. For this, I would recommend some combination of:

- [tarsnap](https://www.tarsnap.com/)
- an online git service, like GitHub, GitLab or Codeberg
- a local git repository in a different machine
- [ipfs](https://ipfs.io/)

## Conclusion

We have seen how to create a website which is decently anonymous and hard to attack. This isn't secure enough to support a drug empire (unless?), but it should be secure enough to post slightly edgy content online. We did this the hard way in order to not rely on services like WordPress or Netlify, which would probably buckle under external pressure.
