# Wikara - Simple Wiki Application

## Overview
Wikara is a lightweight wiki application written in Go. It allows users to create, edit, and view wiki pages through a web interface. Wikara uses Markdown syntax for formatting wiki content.

## Features
- View existing wiki pages
- Edit wiki pages
- Save edited wiki pages
- Create new wiki pages

## Configuration
Wikara can be configured using a YAML file named `config.yaml` that is placed in the same directory as the server binary. The configuration file allows you to customize various aspects of the application, such as server settings and site title. Here are the available configuration options:

- `Port`: The port on which the server will listen (default: 8080)
- `Host`: The host address on which the server will listen (default: 0.0.0.0)
- `SSL`: Whether to enable SSL/TLS (default: false)
- `SSLCertFile`: Path to the SSL certificate file (required if SSL is enabled)
- `SSLKeyFile`: Path to the SSL private key file (required if SSL is enabled)
- `ContentDir`: Directory where wiki page content will be stored (default: data)
- `FrontPageTitle`: Title of the front page (default: FrontPage)
- `SiteTitle`: Title of the wiki site (default: Wikara)
- `SiteLogoURL`: URL to an image that will act as the sites logo (Recommended size of 100x100 pixels)

Example `config.yaml`:
```yaml
Port: 8080
Host: 0.0.0.0
SSL: false
ContentDir: data
FrontPageTitle: FrontPage
SiteTitle: MyWiki
SiteLogoURL: /assets/mylogo.png
```

## Getting Started

1. Clone the Wikara repository.

``` console
git clone https://github.com/bepisdev/wikara.git
cd wikara
```

2. Run make to generate the server binary, along with a configuration file.

``` console
make
```

3. Copy the `dist` folder to your server that you plan on using to host the Wiki.

4. Edit the configuration file to your needs, and then run the server binary

``` console
./wikara
```

5. You can now visit the Wiki at `http://<your-ip-address>:8080` (or whatever you configure it to).

## Todo

- Implement users/authentication to protect pages and edit functionality
	- Page ownership
- Add logo to the sidebar
