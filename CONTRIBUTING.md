# Contributing to `kapp`

Thanks for thinking about contributing to `kapp`! All types of contributions are
welcome:

*   [Issues](#issues)
*   [Documentation](#documentation)
*   [Tests](#tests)
*   [Features](#features)

Start with setting up your environment by following the instructions in
[development setup](#development-setup).

## Issues

Please use [GitHub issues](https://github.com/peterj/kapp/issues) to report bugs
or submit new feature request. Before creating a new issue, do a quick search to
see if same/similar issue is already logged and consider adding +1 to the
existing issue. If no issue exists, feel free to open a new one and include as
much information as possible (things like tools versions, steps to reproduce it,
etc.) to be able to reproduce the behavior.

## Documentation

People make mistekes, forget things, don't explain concepts well and so forth.
Documentation contributions are extremely important as docs are the first thing
users see. Feel free to expand on the existing documentation, change or update
it to make it clearer or add more examples.

## Tests

Check out the [current code coverage](https://codecov.io/gh/peterj/kapp) to see
where test coverage is missing. Follow the
[development setup](#development-setup) to get your machine ready for
development.

## Features

If you have an idea for a new feature, feel free to log an issue and describe
what the feature would be and get feedback from others. Next, get your machine
ready for development by following the instructions in [development
setup)[#development-setup] and write some code (don't forget to add tests for
any new code your write).

## Development setup

1.  Fork and clone the repo.
    > Note: your `master` branch should be pointing at the original repo and you
    > will be making pull requests from branches on your fork. Do this to set it
    > up:
    >
    > ```
    > git remote add upstream https://github.com/peterj/kapp.git
    > git fetch upstream
    > git branch --set-upstream-to=upstream/master master
    > ```

2)  Run `dep ensure` to get all dependencies.
3)  Run `make all` to build the app, run tests and other tools.
4)  Run the app `./kapp` to ensure it works.

## Make changes, commit and open a pull request

After you've made your changes, make sure you run `make all` to ensure there are
no issues (no failing tests, no vet/fmt issues). Finally, create a pull request
from your fork against the original repo.

### Create new release

To create a new release, follow the steps below:

1.  `make bump-version`
2.  `make all`
3.  `make release`
4.  `make tag`
5.  Push the tag (e.g. `git push origin [TAG]`)
6.  Create a new release on [GitHub](https://github.com/peterj/kapp/releases)
    and upload the binaries
