---
name: Hugo
author: cbrgm, maurerle
description: Use Hugo static site generator to create HTML output
tags: [ tool, hugo, generation, static ]
icon: https://raw.githubusercontent.com/gohugoio/gohugoioTheme/refs/heads/master/static/images/icon-h/hugo-h-1.svg
repo: maurerle/woodpecker-hugo
containerImage: ghcr.io/maurerle/woodpecker-hugo
containerImageUrl: https://github.com/maurerle/woodpecker-hugo/pkgs/container/woodpecker-hugo
url: https://github.com/maurerle/woodpecker-hugo
---

The Hugo plugin automatically generates static web page files, which can be published afterwards!

The example below demonstrates how you can use the plugin to automatically create static web page files using Hugo plugin:

```yml
steps:
- name: build
  image: ghcr.io/maurerle/woodpecker-hugo
  settings:
    hugo_version: 0.144
```

#### Customize source, output, theme, config, layout OR content directory paths

You can customize the paths for e. g. the theme, layout, content directory and output directory and much more!

```yml
steps:
- name: build
  image: ghcr.io/maurerle/woodpecker-hugo
  settings:
    hugo_version: 0.144
    url: https://foo.com
+   config: path/to/config
+   content: path/to/content/
+   layout: path/to/layout
+   output: path/to/public
+   source: path/to/source
+   theme: path/themes/THEMENAME/
```

#### Set hostname (and path) to the root

You can also define a base URL directly in the pipeline, which is used when generating the files.

```yml
steps:
- name: build
  image: ghcr.io/maurerle/woodpecker-hugo
  settings:
    hugo_version: 0.144
    url: https://foo.com
    config: path/to/config
    content: path/to/content/
    layout: path/to/layout
    output: path/to/public
    source: path/to/source
    theme: path/themes/THEMENAME/
+   url: https://foo.com
```

#### Build sites and include expired, drafts or future content

You can set the `buildDrafts`, `buildExpired`, `buildFuture` settings to configure the generated files.

- `buildDrafts` - include content marked as draft
- `buildExpired` - include expired content
- `buildFuture` - include content with publishdate in the future

```yml
steps:
- name: build
  image: ghcr.io/maurerle/woodpecker-hugo
  settings:
    hugo_version: 0.144
    url: https://foo.com
    config: path/to/config
    content: path/to/content/
    layout: path/to/layout
    output: path/to/public
    source: path/to/source
    theme: path/themes/THEMENAME/
    url: https://foo.com
+   buildDrafts: true
+   buildExpired: true
+   buildFuture: true
```

#### **Example**: Generate Hugo static files and publish them to remote directory using scp

Here is a short example of how to define a pipeline that automatically generates the static web page files with Hugo and then copies them to a remote server via scp. This makes publishing websites a breeze!

```yml
steps:
- name: build
  image: ghcr.io/maurerle/woodpecker-hugo
  settings:
    hugo_version: 0.144
    pull: always
    url: https://foo.com

- name: deploy
  image: appleboy/drone-scp
  settings:
    host: 192.168.162.1
    target: /var/www/site
    source: public/*
    username:
      from_secret: ssh_username
    password:
      from_secret: ssh_password
    port:
      from_secret: ssh_port
  when:
    branch:
    - master
    event:
      exclude:
      - pull_request
```

# Parameter Reference

buildDrafts
: Include content marked as draft

buildExpired
: include expired content

buildFuture
: include content with publishdate in the future

config
: config file (default is path/config.yaml|json|toml)

content
: filesystem path to content directory

layout
: filesystem path to layout directory

output
: filesystem path to write files to

source
: filesystem path to read files relative from

theme
: theme to use (located in /themes/THEMENAME/)

url
: hostname (and path) to the root

# Contributing

You have suggestions for improvements or features you miss? You are welcome to express all your wishes here. Just create a new Issue on Github and it will be taken care of quickly!

Github: https://github.com/maurerle/woodpecker-hugo