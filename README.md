<a name="readme-top"></a>

# goatsapp

## Getting started

Once you clone the repository, you need to make sure you have already installed the necessary tools and set up your project by following these simple steps below:

### Prerequisites

Tools needed:

- [golang](https://go.dev/dl/) >= 1.22
- [bun](https://bun.sh/docs/installation) >= 1.1

### Env var

Adjust env variables with your own configuration:

```sh
cp .env.example .env
```

### DB Setup

Run database migrations with following command:

```sh
# in the root dir
bun run db:migrate:apply
```

### Build

Build the web assets with following command:

```sh
# change directory to web app folder
cd web/app
bun install --frozen-lockfile
bun run build
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## License

This project is [licensed][LICENSE] under the MIT License.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Contact

Hirzi Nurfakhrian - hirzinurfakhrian@gmail.com

Project Link: [https://github.com/hrz8/goatsapp](https://github.com/hrz8/goatsapp)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

[LICENSE]: (./LICENSE)
