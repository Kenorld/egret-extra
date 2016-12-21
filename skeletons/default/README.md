# Welcome to Eject

## Getting Started

A high-productivity web framework for the [Go language](http://www.golang.org/).

### Start the web server:

    eject run myapp

   Run with <tt>--help</tt> for options.

### Go to http://localhost:9000/ and you'll see:

"It works"

### Description of Contents

The default directory structure of a generated Eject application:

    myapp               App root
      app               App sources
        controllers     App controllers
          init.go       Interceptor registration
        models          App domain models
        routes          Reverse routes (generated code)
        views           Templates
      tests             Test suites
      conf              Configuration files
        app.conf        Main configuration file
        routes          Routes definition
      messages          Message files
      public            Public assets
        css             CSS files
        js              Javascript files
        images          Image files

app

    The app directory contains the source code and templates for your application.

conf

    The conf directory contains the application’s configuration files. There are two main configuration files:

    * app.conf, the main configuration file for the application, which contains standard configuration parameters
    * routes, the routes definition file.


messages

    The messages directory contains all localized message files.

public

    Resources stored in the public directory are static assets that are served directly by the Web server. Typically it is split into three standard sub-directories for images, CSS stylesheets and JavaScript files.

    The names of these directories may be anything; the developer need only update the routes.

test

    Tests are kept in the tests directory. Eject provides a testing framework that makes it easy to write and run functional tests against your application.

### Follow the guidelines to start developing your application:

* The README file created within your application.
* The [Getting Started with Eject](http://eject.github.io/tutorial/index.html).
* The [Eject guides](http://eject.github.io/manual/index.html).
* The [Eject sample apps](http://eject.github.io/samples/index.html).
* The [API documentation](https://godoc.org/bitbucket.org/kenorld/eject-core).

## Contributing
We encourage you to contribute to Eject! Please check out the [Contributing to Eject
guide](https://bitbucket.org/kenorld/eject-core/blob/master/CONTRIBUTING.md) for guidelines about how
to proceed. [Join us](https://groups.google.com/forum/#!forum/eject-framework)!
