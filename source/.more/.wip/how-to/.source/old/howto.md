w -> v
th -> z

> Where to host the blog: Our Recommendation: Use Substack or try out a hosted Ghost website—it’s a Goldilocks solution of easy-to-use but highly customizable. Ghost makes for a great portfolio site, blog, and newsletter service.

[Where to host the blog](https://effectiveideas.org/how-to-start-a-blog/#pt3)

No. Vat do I vant, and in vat order? Uncensorability, ability to ramp up degrees of privacy and paranoia, control over ze whole stack, and finally, ze modern comforts. Vith that in mind, Doktor Nefario presents:

# How to create an autarchic website: the easy way

## Step 0: Get you a password manager for great security.

For this, I recommend [pass](https://www.passwordstore.org/) (see also the [archlinux wiki](https://wiki.archlinux.org/title/Pass)). Seems simple, secure.

## Step 1: Register a domain.

For zis, ve will use, of course, [njal.la](https://njal.la/). The're the non-evil twin brother of [Epik domains](https://www.epik.com/). From their [FAQ](https://njal.la/faq/):

> We're a team of committed internet activists and we're also involved in other privacy projects such as the IPredator VPN service. Some of us have also been involved in projects like The Pirate Bay and Piratbyrån to mention a few things.

> We support signing up for our service using email or XMPP (and yes, we use OTR). For email we also support PGP so all of our outgoing email will be signed and encrypted.

> We accept payments via Bitcoin, Litecoin, Monero, ZCash, DASH, Bitcoin Cash, Ethereum and Paypal.

They also have [a blog](https://njal.la/blog/) vere they record takedown requests:

![](https://njal.la/static/downloads/2019njetski/3.jpg)

All in all, zis is ze good shit.

```
Mad Eye Moody recommends: Use tornado cash (https://tornado.cash/) to anonymize your transaction history before paying njalla. Check which top-level domain names are more uncensorable (https://www.eff.org/files/2017/08/02/domain_registry_whitepaper.pdf).
```

Ultimately, if we wanted further anonymity and uncensorability, we would [set up an onion service](https://community.torproject.org/onion-services/setup/).

## Step 2: Get a server

Previously, I have been using an Ubuntu server on [Digital Ocean](digitalocean.com), which offers servers from $5/month ($6/month including backups). But I have grown to dislike my system having 1GB worth of stuff that I don't understand. I also dislike wrangling with systemd services.

So I thought it would be fun to try Alpine Linux on [Njal.la's own servers](https://njal.la/servers/add/). This costs $15/month instead.

When buying a server from Njalla, we need to name it. And we will name it [`ramiel`](https://en.wikipedia.org/wiki/Ramiel) [^1]. We will also need to generate an ssh key to connect to it (`ssh-keygen -t ed25519`). After paying for our server, we'll want to set up payment to autorenew.

[^1]: Why Ramiel? Because Uriel is dead.

## Set up the server

We can connect to our server with

```
ssh root@your_ip
```

### Essential tooling

Then add some essential tooling:

```
apk add vim
apk add sudo
apk add git
```

### Add password to root user

This is useful in order to be able to drop into su as a user.

(Unclear whether njala asks this)

```
pass generate /njal.la/root
```

### Add a user with root permissions

Then add a root user with

```
adduser ramiel
apk add doas
adduser ramiel wheel
mkdir -p /home/ramiel
mkdir /home/ramiel/.ssh

```

Then add `permit persist :wheel` to `/etc/doas.d/doas.conf`

Then use `visudo` to edit `/etc/sudoers`, and add

```
ramiel ALL=(ALL) ALL
```

When creating a new user, we are prompted with a password. You can generate one with `pass`:

```
pass generate njalla/ramiel
```

and then show it with

```
pass show njalla/ramiel
```

### Allow ssh connection to that root user

But we want to log in with ssh, so we add:

```
cp /root/.ssh/authorized_keys /home/ramiel/.ssh/authorized_keys
chown ramiel /home/ramiel/.ssh/authorized_keys # https://stackoverflow.com/questions/61833713/how-to-login-by-ssh-in-alpine-linux-without-passwords
```

and we change `/etc/ssh/sshd_config`

```
PasswordAuthentication no
PermitRootLogin prohibit-password
```

to:

```
PasswordAuthentication no
PermitRootLogin yes
```

Then we restart the ssh service

```
service sshd restart
exit
```

and we connect with our new user

```
ssh ramiel@your_ip
```

#### lf

```
wget https://github.com/gokcehan/lf/releases/download/r26/lf-linux-amd64.tar.gz
tar x -f lf-linux-amd64.tar.gz
sha512sum lf ## 4c2a2907987bfd20a11a6c99004a524f91ee7c29f8329626af5bc315ba348966021672a24d6bce5ca3674108e82b9ee1d9f83b8a6d30b539a2e788aefbec0918
mv lf /usr/bin/lf0
```

Then paste the following into your `~/.profile`

```
lf () {
    tmp="$(mktemp)"
    lf0 -last-dir-path="$tmp" "$@"  --command "set hidden"
    if [ -f "$tmp" ]; then
        dir="$(cat "$tmp")"
        rm -f "$tmp"
        if [ -d "$dir" ]; then
            if [ "$dir" != "$(pwd)" ]; then
                cd "$dir"
            fi
        fi
    fi
}
```

## Install plan9port

Add community packages to your `/etc/apk/repositories`, by adding the following line:

```
https://dl-cdn.alpinelinux.org/alpine/v3.15/community/
```

```
apk install plan9port
## https://pkgs.alpinelinux.org/package/v3.15/community/x86_64/plan9port
```

By default, plan9port gets installed to `/usr/lib/plan9/bin`

```
cp /usr/lib/plan9/bin/rc /bin/rc
```

Also fix it somewhere in the config files [Note: important]

## Install werc

```
cd ~
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
```

## Set up nginx

Following instructions from [here](https://wiki.alpinelinux.org/wiki/Nginx)

Here we make our first—of many—compromises with the suckless philosophy. Unlike simpler web servers, nginx is not suckless. However, I do have a job, and I'm choosing nginx because it has good integration with certbot, the certificates bot from the Electronic Freedom Foundation.

```
su
apk update
apk add nginx
adduser -D -g 'www' www
```

```
cd /home/ramiel/werc/werc-1.5.0/sites/
mkdir example.com
```

Configure nginx

```
mv /etc/nginx/nginx.conf /etc/nginx/nginx.conf.orig
```

Example config page:

```
user                            www;
worker_processes                auto; # it will be determinate automatically by the number of core

error_log                       /var/log/nginx/error.log warn;
pid                             /var/run/nginx/nginx.pid; # it permit you to use /etc/init.d/nginx reload|restart|stop|start

events {
    worker_connections          1024;
}

http {
    include                     /etc/nginx/mime.types;
    default_type                application/octet-stream;
    sendfile                    on;
    access_log                  /var/log/nginx/access.log;
    keepalive_timeout           3000;
    server {
        listen                  80;
        root                    /home/ramiel/werc/werc-1.5.0/sites/example.com;
        index                   index.html index.htm;
        server_name             localhost;
        client_max_body_size    32m;
        error_page              500 502 503 504  /50x.html;
        location = /50x.html {
              root              /var/lib/nginx/html;
        }
    }
}
```

Autostart:

```
rc-service nginx start
```

## Add a DNS record to your site

In particular an A record, that points to the IP of your server. Then, create a

## Tinker with nginx until it works.

Change the nginx config file (`/etc/nginx/nginx.conf `) to something like what follows. This has the advantage of being able to manage more than one domain at a time.

```
user www;
worker_processes auto;
pid /run/nginx.pid;
include /etc/nginx/modules-enabled/*.conf;

events {
        worker_connections 768;
        # multi_accept on;
}

http {

        ##
        # Basic Settings
        ##

        sendfile on;
        tcp_nopush on;
        types_hash_max_size 2048;
        # server_tokens off;

        server_names_hash_bucket_size 64;
        # server_name_in_redirect off;

        include /etc/nginx/mime.types;
        default_type application/octet-stream;

        ##
        # SSL Settings
        ##

        ssl_protocols TLSv1 TLSv1.1 TLSv1.2 TLSv1.3; # Dropping SSLv3, ref: POODLE
        ssl_prefer_server_ciphers on;

        ##
        # Logging Settings
        ##

        access_log /var/log/nginx/access.log;
        error_log /var/log/nginx/error.log;

        ##
        # Gzip Settings
        ##

        gzip on;

        # gzip_vary on;
        # gzip_proxied any;
        # gzip_comp_level 6;
        # gzip_buffers 16 8k;
        # gzip_http_version 1.1;
        # gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript;

        ##
        # Virtual Host Configs
        ##

        include /etc/nginx/conf.d/*.conf;
        include /etc/nginx/sites-enabled/*;
}
```

Make some necessary directories

```
mkdir -p /etc/nginx/sites-available/
mkdir -p /etc/nginx/sites-enabled/
```

Add the following to `/etc/nginx/sites-available/example.com`

```
server {
    server_name  samotsvety.org; # Replace with your domain name.

    #charset utf-8;

    #access_log  logs/host.access.log  main;

    location / {

    # FastCGI params, usually stored in fastcgi_params
    # and imported with a command like the following:
    #include        fastcgi_params;

    # Typical contents of fastcgi_params (inlined here):
    fastcgi_pass localhost:9000;

    fastcgi_param  QUERY_STRING       $query_string;
    fastcgi_param  REQUEST_METHOD     $request_method;
    fastcgi_param  CONTENT_TYPE       $content_type;
    fastcgi_param  CONTENT_LENGTH     $content_length;

    #fastcgi_param  SCRIPT_FILENAME   /var/www/werc/bin/werc.rc;
    fastcgi_param  SCRIPT_NAME        /home/uriel/workspace/werc-1.5.0/bin/werc.rc;
    #fastcgi_param  SCRIPT_NAME        $fastcgi_script_name;

    fastcgi_param  REQUEST_URI        $request_uri;
    fastcgi_param  DOCUMENT_URI       $document_uri;
    fastcgi_param  DOCUMENT_ROOT      $document_root;
    fastcgi_param  SERVER_PROTOCOL    $server_protocol;

    fastcgi_param  GATEWAY_INTERFACE  CGI/1.1;
    fastcgi_param  SERVER_SOFTWARE    nginx/$nginx_version;

    fastcgi_param  REMOTE_ADDR        $remote_addr;
    fastcgi_param  REMOTE_PORT        $remote_port;
    fastcgi_param  SERVER_ADDR        $server_addr;
    fastcgi_param  SERVER_PORT        $server_port;
    fastcgi_param  SERVER_NAME        $server_name;
    fastcgi_param  REMOTE_USER        $remote_user;

    #root /home/uriel/workspace/werc-1.5.0/sites/samotsvety.org;
    #root   /var/www/werc/sites/$server_addr; # XXX This doesn't work, not sure why :(
    #nuno: I confirm that the above doesn't work.
    root /;
    #index  index.html index.htm;
  }

}

server {
    if ($host = samotsvety.org) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


  listen 80;
    server_name  samotsvety.org;
    return 404; # managed by Certbot


}
```

Link that to sites-enabled

```
sudo ln -s /etc/nginx/sites-available/example.com /etc/nginx/sites-enabled/
```

Right now we haven't really done anything.

## Add certbot

```
apk add --update python3 py3-pip
apk add certbot
pip install certbot-nginx
rc-service nginx stop
sudo certbot --nginx -d samotsvety.org

```

This asks you to input an email address.

You can also see autorenewal instructions [here](https://certbot.org/renewal-setup)

##

Create a werc site

```
cp -r /home/ramiel/werc/werc-1.5.0/sites/example.com /home/ramiel/werc/werc-1.5.0/sites/samotsvety.org

```

## Install fgci

```
sudo apk add spawn-fcgi
sudo apk add fcgiwrap

```

/usr/bin/spawn-fcgi -a 127.0.0.1 -p 9000 -f /usr/bin/fcgiwrap

/usr/bin/spawn-fcgi
/usr/bin/fcgiwrap

rc-update add nginx default
sudo chmod 755 spawn-fcgi2
sudo rc-update add spawn-fcgi2
touch /tmp/spawnfcgi-werc
sudo chmod 666 /tmp/spawnfgci-werc

## Add spawn-fcgi2 to run

copy file to /etc/init.d/spawn-fcgi2
chmod 755
rc-update add nginx default

## change http to https in ../bin/werc.rc

lines 37 and 99

```
chown www -R werc
```

## Change initrc.local and the .rc files to point to plan9, e.g., /usr/lib/plan9

## Install discount markdown

sudo apk install discount
Chang the discount markdown

## Debug problems.

## Wormhole

https://gist.github.com/stchris/68686875b7c9290040fc907d60473aa7
apk update && apk add build-base python3-dev libffi-dev openssl-dev && pip3 install magic-wormhole
I am sacrificing 600MB for this, which does seem worth it because it's sooo convenient, but which still hurts.
