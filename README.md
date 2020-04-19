## go-react

I made this tool so I can quickly startup a webapp with a go server and react frontend whenever I feel like it.
If you also want to use this go ahead.

install using
```
go install github.com/lokucrazy/go-react
```

and run:
```
go-react PROJECT_NAME USER/REPO
```

PROJECT_NAME is the name of the project and the folder name

USER/REPO is the module path for the go.mod file

it will then create the necessary files and then tell you to go into the frontend folder and do `npm install`.
I was getting some permission errors on windows when trying to exec `npm install` in the go code, I'm like writing this README at 3 am so I'll deal with it later.

if you have any questions or you like/dislike this repo let me know.

Have fun.