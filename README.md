This is stahnma's experimental bot.


Basically, I got really tired of coffeescript. I saw that Dan Ryan was writing some similar in go. I should use that, becuase golang sounds way more fun than stuff that uses node.

# HOWTO

## Using this
The easiest thing to do right now is create a file with all the env vars you need.

    export HAL_HIPCHAT_USER=147859_1068038
    export HAL_HIPCHAT_PASSWORD=secret
    export HAL_HIPCHAT_ROOMS="Websages,foobar"
    export HAL_HIPCHAT_RESOURCE=bot
    export HAL_ADAPTER=hipchat
    export HAL_LOG_LEVEL=debug

Then source that file. Then make; then run the binary.

## godep
This project uses godep for library management.

Layout your code like this http://golang.org/doc/code.html
Here's another article on godep http://www.goinggo.net/2013/10/manage-dependencies-with-godep.html


#TODO

There are oh so many things to do to make this worth while.


    * Package for deployment
    * Have init scripts
    * Figure out logging
    * Perhaps put it in a container
    * Remove manual step of sourcing env vars

