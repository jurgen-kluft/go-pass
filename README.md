Introducing `gpw`
------------------

Password management should be simple and follow [Unix philosophy](http://en.wikipedia.org/wiki/Unix_philosophy). 

`gpw` makes managing these individual password files extremely easy. All passwords live in a `password store`, and `gpw` provides some nice commands for adding, editing, generating, and retrieving passwords. It is a very short and simple application. It's capable of temporarily putting passwords on your clipboard and tracking password changes using [`git`](http://en.wikipedia.org/wiki/Git_(software)). The `git` repository used is one that uses an encrypted store, this is provided by using an App called `Keybase` from http://keybase.io.

You can edit the password store using ordinary unix shell commands alongside the `gpw` command. There are no funky file formats or new paradigms to learn. There is [bash](http://en.wikipedia.org/wiki/Bash_(Unix_shell)) [completion](http://en.wikipedia.org/wiki/Command-line_completion) so that you can simply hit tab to fill in names and commands, as well as completion for [zsh](http://en.wikipedia.org/wiki/Z_shell) and [fish](http://en.wikipedia.org/wiki/Friendly_interactive_shell) available in the [completion](https://git.jay_at_gmail.com/password-store/tree/src/completion) folder. 

The `gpw` command is extensively documented in its [man page](https://github.com/jurgen-kluft/go-pass/docs).

### Using the password store

We can list all the existing passwords in the store:

    jay@mac ~ $ pass list
    Password Store
    ├── Sites
    │   ├── amazon.com
    │   └── tweakers.net
    ├── Email
    │   ├── sophia.hotmail.com
    │   └── jay.gmail.com
    └── Nederland
        ├── bank
        ├── tweakers
        └── mobilephone
    

And we can show passwords too:

    jay@mac ~ $ pass show email/jay_at_gmail.com
    sup3rh4x3rizmynam3

Or copy them to the clipboard:

    jay@mac ~ $ pass show -c email/jay_at_gmail.com
    Copied email/jay_at_gmail.com to clipboard.

We can add existing passwords to the store with `insert`:

    jay@mac ~ $ pass insert business/cheese-whiz-factory
    Enter password for business/cheese-whiz-factory: omg so much cheese what am i gonna do
    

This also handles multiline passwords or other data with `--multiline` or `-m`, and passwords can be edited in your default text editor using `pass edit pass-name`.

The utility can `generate` new passwords using `/dev/urandom` internally:

    jay@mac ~ $ pass generate -l 15 email/jasondonenfeld.com 
    The generated password to email/jasondonenfeld.com is:
    $(-QF&Q=IN2nFBx

It's possible to generate passwords with no symbols using `--no-symbols` or `-n`, and we can copy it to the clipboard instead of displaying it at the console using `--clip` or `-c`. 

And of course, passwords can be removed:

    jay@mac ~ $ pass rm business/cheese-whiz-factory
    rm: remove regular file ‘/home/jay_at_gmail/.password-store/business/cheese-whiz-factory.gpg’? y
    removed ‘/home/jay_at_gmail/.password-store/business/cheese-whiz-factory.gpg’
    
If the password store is a git repository, since each manipulation creates a git commit, you can synchronize the password store using `pass git push` and `pass git pull`.

You can read more examples and more features in the [man page](https://git.jay_at_gmail.com/password-store/about/).

### Setting it up

To begin, there is a single command to initialize the password store:

    jay@mac ~ $ pass init ‘$HOME/.password-store’
    mkdir: created directory ‘/home/jay_at_gmail/.password-store’
    Password store initialized for jay_at_gmail

We can additionally initialize the password store as a git repository:

    jay@mac ~ $ pass init --git ‘$HOME/.password-store’
    Initialized empty Git repository in /home/jay_at_gmail/.password-store/
    jay@mac ~ $ pass git remote add origin https://github.com/jay_at_gmail/pass-store

If a git repository is initialized, `gpw` creates a git commit each time the password store is manipulated.

Download
--------

The latest version is 0.1.1.

### Git Repository

You may [browse the git repository](https://github.com/jurgen-kluft/gpw/) or clone the repo:

    $ git clone https://github.com/jurgen-kluft/gpw/

Data Organization
-----------------

### Usernames, Passwords, PINs, Websites, Metadata, et cetera

The password store does not impose any particular schema or type of organization of your data, as it is simply a flat text file, which can contain arbitrary data. Though the most common case is storing a single password per entry, some power users find they would like to store more than just their password inside the password store, and additionally store answers to secret questions, website URLs, and other sensitive information or metadata. Since the password store does not impose a scheme of it's own, you can choose your own organization. There are many possibilities.

The password itself is stored on the first line of the file, and the additional information on subsequent lines. 
For example, `Amazon/bookreader` might look like this:

    Yw|ZSNH!}z"6{ym9pI
    URL: *.amazon.com/*
    Username: AmazonianChicken@example.com
    Secret Question 1: What is your childhood best friend's most bizarre superhero fantasy? Oh god, Amazon, it's too awful to say...
    Phone Support PIN #: 84719

_This is the scheme used here._ The `--clip` / `-c` options will only copy the first line of such a file to the clipboard, thereby making it easy to fetch the password for login forms, while retaining additional information in the same file.
