# CronTalk

Translate the cron expressions from your terminal into english(or any other supported language) words.

If you are building an application which requires cron jobs by parsing cron expressions,  this tool will help to check if your expression is the desired one. While there are multiple online cron expression parsing tools available, this is made with the sole intention of making things faster & easier. Because we, as developers use the terminal(CLI) very often & rather than depending on something online, it would be faster for us to get this little information with a single line of command.

So try this out.

### Installation

1. [Download the binary](https://github.com/Anondo/crontalk-bin) for your OS

| OS      	| x86                                                                                      	| x86_64                                                                                      	|
|---------	|------------------------------------------------------------------------------------------	|---------------------------------------------------------------------------------------------	|
| Mac     	| [Download](https://github.com/Anondo/crontalk-bin/blob/master/v2/mac_386.zip)     	| [Download](https://github.com/Anondo/crontalk-bin/blob/master/v2/mac_amd64.zip)     	|
| Linux   	| [Download](https://github.com/Anondo/crontalk-bin/blob/master/v2/linux_386.zip)   	| [Download](https://github.com/Anondo/crontalk-bin/blob/master/v2/linux_amd64.zip)   	|
| FreeBSD 	| [Download](https://github.com/Anondo/crontalk-bin/blob/master/v2/freebsd_386.zip) 	| [Download](https://github.com/Anondo/crontalk-bin/blob/master/v2/freebsd_amd64.zip) 	|
| Windows 	| [Download](https://github.com/Anondo/crontalk-bin/blob/master/v2/windows_386.zip) 	| [Download](https://github.com/Anondo/crontalk-bin/blob/master/v2/windows_amd64.zip) 	|



2. **For Mac/Linux**, add the binary to your $PATH variable or copy it in /usr/bin  to access this from everywhere in the terminal.

   **For Windows**  , set the environmental variable to the path of the binary for accessing the app from the terminal

1. Test your installation by hitting,

```
$ crontalk

```

or

```
$ crontalk version

```

from the terminal.

If you can see the version of the app then you are good to go!.

### Usage

#### Translate

Translate any cron expressions like this,
```
$ crontalk translate --cron="6 12 * * *"
```

And the result will be ```Every Day At 12:06PM```.

Also **CronTalk** supports multiple languages (english & bangla for now with english being the default). So try,

```
$ crontalk translate --cron="6 12 * * *" --bangla

```
And you will get something like ```প্রতিদিন সময় ১২:০৬PM```

#### Generate

Generate a cron expression from english words. For now the valid & proper english words are prompted by the app for the user to provide input. The ```generate``` command works like this,
```
$ crontalk generate
```

And you will be prompted with the english words required for all the sub-expression for a cron expression(minute,hour/day of month, month,  day of week) one by one like this,

```
Use the arrow keys to navigate: ↓ ↑ → ←
? Minute:
  ▸ done
    every minute
    <input>
    every <input> minutes
↓   every <input> minutes from <input> to 59
```
Until the ```done``` option is selected the particular sub-expression keeps on taking input. This is done because cron sub-expressions can often be listed. Selecting the ```<input>``` for example will block the app for a user input. After providing the input the user will be prompted for the next sub-expression like ```Hour``` and so on upto ```Week```. And finally the result will be something like this,

```
The cron expression: 30 12 * 6 *
Translation: Every June At 12:30PM

```

#### Other Available Commands

List the next occurrence(s) for a cron expression,
```
$ crontalk next --cron="6 12 * * *" -o 5
```

Doing this will give you

```
2019-05-23 12:06PM
2019-05-24 12:06PM
2019-05-25 12:06PM
2019-05-26 12:06PM
2019-05-27 12:06PM
```

The next 5 occurrences of the given expression. The **-o** / **--occurrence** flag , as you have guessed it, will determine the number of occurrence to display.

A command called ```serve``` will soon be fully working(hopefully on the next release) which will open a port on the local machine to browse a web based UI for ```CronTalk```.

### Constraints

  \* \* \* \* \*  = (minute) (hour) (day of month) (month) (day of week)

1. Should contain exactly 5 values/sub-expressions

1. Valid values are:

                      minute: 0-59

                      hour:   0-23

                      day of month: 1-31

                      month: 1-12 or jan-dec

                      day of week: 0-6 or sun-sat

                      list values example: 1,2,3

                      range values example: 1-4 , mon-thu

                      step values example: 1/8

### Contributing

See the contributions guide [here](CONTRIBUTING.md).

### License

CronTalk is licensed under the [MIT License](LICENSE).
