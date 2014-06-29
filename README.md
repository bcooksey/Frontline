==== Frontline ====

A game.

==== Building and Running ====

Navigate to the root directory of the application and run:

```
go get
go build
npm install
./node_modules/.bin/uglifyjs static/js/libs/jquery-1.11.0.min.js static/js/libs/underscore-min.js static/js/libs/backbone-min.js static/js/libs/jquery-cookie.js -o static/compiled/js/libs.js -m --screw-ie8 -c
./frontline
```

==== Testing ====

`go test`
