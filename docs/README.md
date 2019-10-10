# NLX documentation site

The NLX documentation site is a resource for developers that want to use and provide services on the NLX network.

This site is generated with [Hugo](https://gohugo.io/), a static site generator written in Go.

The current `master` version of the documentation is deployed at [docs.nlx.io](https://docs.nlx.io/)

## Editing the content of the site

To edit pages, edit the [MarkDown](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet) files in the [`/content` folder](content/).

Each file needs to have Hugo [front-matter](https://gohugo.io/content-management/front-matter/) at the top.
The front-matter sets the settings for each page, including `title` and the `weight` that determines the order that a page displays in lists.


## Development

Install hugo

    $ sudo apt-get install hugo

    or
    
    $ snap install hugo

    or 

    $ brew install hugo


Serve the html with live changes

    $ hugo server
