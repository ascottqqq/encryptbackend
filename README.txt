encryptbackend
==============

Description
-----------

A go encryption backend taking ajax queries from clients

Relies on the library github.com/ascottqqq/rfc7539

Installation
------------

This package can be installed with the go get command:

    go get github.com/aascottqqq/encryptbackend

To install dependencies do the following from the project directory:

    go get ./...

Notes
-----

Without a client to post the json requests this application is of
limited use.

Also, it is not a good idea to submit plaintext over the internet.
So should be used only on a secure network (or locally) or
updated to use https