
# ecrypto


ecrypto is in devlopment.  It is not yet a super hyper cool program but
we are going there.

###  name

*ecrypto* really sucks.  Let's find a better name.

*  *ecrypto* (easy-crypto)
*  *nfc* (npm for cryptography)
*  *monsieur gustave*  one needs to watch the Grand Budapest Hotel





## Use case

ecrypto is the *npm* for cryptography.  It revolutionizes the way the people
use cryptography on a everyday use.  The general design is to do what programs
like *npm* did for dependencies management but for cryptography.

ecrypto user interface is a command line user interface (*CLI*) and its source
code.  For us, the readability of the source code is as important as the
functionality of the software.

###  using ecrypto

#### creating keys

The first thing you want to do when using ecrypto is to create your personal
private-public key pair.  your keys will be kept in
your `~/.ecrypto` directory.

~~~
ecrypto createkey
~~~

Then, you might want to publish your public key to the public repository

~~~
ecrypto publish -email alice@lechef.com
~~~

#### encrypting file

To encrypt a file to send to *bob@lechef.com*, you can do the following
things

1.  if *bob@lechef.com*'s public key is on the public repository, you can
    simply use
    ~~~
    ecrypto encrypt -to bob@lechef.com -file mysecretdocument.doc -out mycypherdoc.cyp
    ~~~

2.  if *bob@lechef.com*'s public key is on your local file path you can use
    ~~~
    ecrypto encrypt -to /path/to/bob@lechef.com -file mysecretdocument.doc -out mycypherdoc.cyp
    ~~~

3.  if *bob@lechef.com*'s public key is on accessible on a public url you can
    use
    ~~~
    ecrypto encrypt -to https://example.com/path/to/bob@lechef.com -file mysecretdocument.doc -out mycypherdoc.cyp
    ~~~

In anycase, the command line output will be a cypher file called 
mycypherdoc.cyp.  You are then responsible to send the file to its recipient
by the mean you wish; email, dropbox, floppy disc...

#### decrypting file

Once you received a file `encrypted_file.cyp`, you can decrypt it via the
following command.

~~~
ecrypt decrypt encrypted_file.cyp
~~~


#### so you want to hack ecrypto



That's cool because it is made for that!  That you are a high school student
with basic programming knownledge, a math student that wants to see an 
implementation of cryptographic algorithm or a senior security engineer, we
welcome you!

Before everything, you will need to learn a little bit of [Go](http://golang.org).
We really like the official [go tour](http://tour.golang.org/) because you
do not need to download or install anything; everything runs in your web
browser.  Then, you only need to dig into the code!

We are proud of our source code.  The unit test code coverage is almost
100%.  All the functions are bellow 10 line of code.  So if you do not
understand a small piece of code, just look at its unit tests.

All our file formats are well documented [yaml](http://www.yaml.org/) files.
You can open any file and read the comments and make sure everything
make sense.

---

## first road map

*  make a simple web service to use as remote key location
*  make blue print for the file

---

## things to do not forget

In the long run, there are some things we will want to do and we need not
to forget about it or lock ourselves out of those goals.

*  Use many encryption algorithms (symetrics, asymetrics, modes and hashes)
*  Has many keys for the same email
*  Rotate its keys
*  Encrypt a file to many people at the same time
*  Keys could be some where else like in [npm](https://www.npmjs.org/doc/cli/npm-install.html)
*  Using threshold cryptography
*  make a neet web site like npm.org to store keys 
*  make alias for known public keys
*  make it easy for rapper applications

---
