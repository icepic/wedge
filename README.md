<p align="center">Wedge is a general-purpose HTTP/2 web server that serves HTTPS by default.</p>
<p align="center"><strong>Wedge is a fork of Caddy</strong>, with sponsor headers removed. The project is unofficial.</p>
<p align="center">As a result of issue [#2](https://github.com/WedgeServer/wedge/issues/2), references to Caddy in this documentation are being changed to reference the name "Wedge" instead.</p>
---

Wedge is fast, easy to use, and makes you more productive.

Available for Windows, Mac, Linux, BSD, Solaris, and Android.

## Building Wedge

```
go get github.com/caddyserver/builds
go get github.com/WedgeServer/wedge/caddy
```

Then `cd` to `wedge/caddy` and `go run build.go`.

## Menu

- [Features](#features)
- [Install](#install)
- [Quick Start](#quick-start)
- [Running in Production](#running-in-production)
- [Contributing](#contributing)
- [Donors](#donors)
- [About the Project](#about-the-project)

## Features

- **Easy configuration** with the configuration file
- **Automatic HTTPS** on by default (via [Let's Encrypt](https://letsencrypt.org))
- **HTTP/2** by default
- **Virtual hosting** so multiple sites just work
- Experimental **QUIC support** for those that like speed
- TLS session ticket **key rotation** for more secure connections
- **Extensible with plugins** because a convenient web server is a helpful one
- **Runs anywhere** with **no external dependencies** (not even libc)


## Install

Follow the instructions above for building from source. Binaries will be available soon.

## Quick Start

To serve static files from the current working directory, run:

```
caddy
```

The default port is 2015, so open your browser to [http://localhost:2015](http://localhost:2015).

### Go from 0 to HTTPS in 5 seconds

If the binary has permission to bind to low ports and your domain name's DNS records point to the machine you're on:

```
caddy -host example.com
```

This command serves static files from the current directory over HTTPS. Certificates are automatically obtained and renewed for you!

### Customizing your site

To customize how your site is served, create a file named Caddyfile by your site and paste this into it:

```plain
localhost

push
browse
websocket /echo cat
ext    .html
log    /var/log/access.log
proxy  /api 127.0.0.1:7005
header /api Access-Control-Allow-Origin *
```

When you run `caddy` in that directory, it will automatically find and use that Caddyfile.

This simple file enables server push (via Link headers), allows directory browsing (for folders without an index file), hosts a WebSocket echo server at /echo, serves clean URLs, logs requests to an access log, proxies all API requests to a backend on port 7005, and adds the coveted  `Access-Control-Allow-Origin: *` header for all responses from the API.

Wow! Caddy can do a lot with just a few lines.

### Doing more with Wedge 

To host multiple sites and do more with the Caddyfile, please see the [Caddyfile tutorial](https://caddyserver.com/tutorial/caddyfile).

Sites with qualifying hostnames are served over [HTTPS by default](https://caddyserver.com/docs/automatic-https).

Wedge has a command line interface. Run `caddy -h` to view basic help or see the [CLI documentation](https://caddyserver.com/docs/cli) for details.


## Running in Production

Wedge is production-ready if you find it to be a good fit for your site and workflow.

**Running as root:** We advise against this. You can still listen on ports < 1024 on Linux using setcap like so: `sudo setcap cap_net_bind_service=+ep ./caddy`

How you choose to run Wedge is up to you. Many users are satisfied with `nohup caddy &`. Others use `screen`. Users who need Wedge to come back up after reboots either do so in the script that caused the reboot, add a command to an init script, or configure a service with their OS.

If you have questions or concerns about Wedge' underlying crypto implementations, consult Go's [crypto packages](https://golang.org/pkg/crypto), starting with their documentation, then issues, then the code itself; as Wedge uses mainly those libraries.


## Contributing


Please see our [contributing guidelines](https://github.com/WedgeServer/wedge/blob/master/.github/CONTRIBUTING.md) for instructions.
 
We use GitHub issues and pull requests only for discussing bug reports and the development of specific changes.

Thanks for making Wedge -- and the Web -- better!


## Donors

- [DigitalOcean](https://m.do.co/c/6d7bdafccf96) is hosting the Caddy project.
- [DNSimple](https://dnsimple.link/resolving-caddy) provides DNS services for Caddy's sites.
- [DNS Spy](https://dnsspy.io) keeps an eye on Caddy's DNS properties.

We thank them for their services. **If you want to help keep Caddy free, please [become a sponsor](https://caddyserver.com/pricing)!**


## About the Project

Caddy was born out of the need for a "batteries-included" web server that runs anywhere and doesn't have to take its configuration with it. Caddy took inspiration from [spark](https://github.com/rif/spark), [nginx](https://github.com/nginx/nginx), lighttpd,
[Websocketd](https://github.com/joewalnes/websocketd) and [Vagrant](https://www.vagrantup.com/), which provides a pleasant mixture of features from each of them.

**The name "Caddy":** The name of the software is "Caddy", not "Caddy Server" or "CaddyServer". Please call it "Caddy" or, if you wish to clarify, "the Caddy web server". See [brand guidelines](https://caddyserver.com/brand).

*Author on Twitter: [@mholt6](https://twitter.com/mholt6)*
