Soothing software
=================

I have this concept of my mind of "soothing software", a cluster of software which is just right, which is competently made, which contains no surprises, which is a joy to use. Here are a few examples:

### pass: "the standard unix password manager"

[pass](https://www.passwordstore.org/) is a simple password manager based on the Unix philosophy. It saves passwords on a git repository, encrypted with [gpg](https://en.wikipedia.org/wiki/GNU_Privacy_Guard). To slightly tweak the functionality of its native commands (pass show and pass insert), I usually use two extensions, [pass reveal](https://git.nunosempere.com/open.source/pass-reveal), and [pass append](https://git.nunosempere.com/open.source/pass-append).

### lf

[lf](https://github.com/gokcehan/lf) is a terminal file manager, similar to [ranger](https://ranger.github.io/) and probably less powerful than [Midnight commander](https://midnight-commander.org/timeline). You can navigate, copy, cut or delete files, or enter into an editor. 

![](https://i.imgur.com/3TQyxz9.png)

### werc

[werc](https://werc.cat-v.org/) is the software which runs this website. It was a pain to get up and running, but I find it soothing to edit my posts directly in markdown, and to know what each part of my server is doing:

![](https://i.imgur.com/X6bNWoY.png)

### vim/neovim

[neovim](https://neovim.io/) is a successor to [vim](https://www.vim.org/). It's a text editor: it allows me to write, anvigate and manipulate text. And through various plugins, it allows for wide-ranging customization, syntax highlighting, autocomplete, and so on. Though its plugins do have set-up costs. I am writting this post in neovim.

![](https://i.imgur.com/UWQCf5I.png)

Unlike its traditional rival, [Emacs](https://www.gnu.org/software/emacs/), vim/neovim is pretty much just a text editor, rather than an ersatz operating system, hence the need for...

### tmux

[tmux](https://github.com/tmux/tmux) is a "terminal multiplexer", i.e., a program that divides a terminal into tabs and parts. Here is a screenshot:

![](https://i.imgur.com/Kisw18E.png)

and here is it annotated:

![](https://i.imgur.com/AFlTtgJ.png)

Normally instead of a blogpost I would have a diary tab open, but I didn't want to show my diary. The point is, though, that tmux makes it very easy to spin up tabs and to subdivide my screen, and then to return to just one window.

### Newsboat

[Newsboat](https://github.com/newsboat/newsboat) is an rss reader. I was previously using [inoreader](https://www.inoreader.com/) and then previously [feedly](https://feedly.com/). But Newsboat is just better for me, it shows no nagging appeals to subscribe to the premium version, no advertisements. It has a simpler counterpart on [newsraft](https://codeberg.org/grisha/newsraft).

Here is an image showing the newsboat feed for Matt Levine's newsletter:

![](https://i.imgur.com/xDYwBpC.png)

### moc

[moc](moc.daper.net) (music-on-console) is a console audio player. It just does its job competently. Recently I figured out that you can save the current playlist into a file, and then [reload that playlist](https://moc.daper.net/node/1368), which made me happy.

![](https://i.imgur.com/SJ7Fgdg.png)

### postmark (paid)

[postmark](https://postmarkapp.com/) is a service that sends email on my behalf. Previously, I was relying on Mailchimp, Substack, or on a hacky connection to gmail through SMPT. Now I'm sending emails through the [cli](https://github.com/activecampaign/postmark-cli), or through [listmonk](https://listmonk.app/)—the software I use to manage my [personal newsletter](https://nunosempere.com/.subscribe/).

### tarsnap (paid)

[tarsnap](https://www.tarsnap.com/) offers "online backups for the truly paranoid", which appeals to me. I'm backing up everything on my computer every hour, while excluding e.g., `node_modules` and a few large filetypes and folders. I find this soothing because it means that if my computer dies, I've lost nothing of importance.

### teddit

[teddit](https://codeberg.org/teddit/teddit) is an alternative frontend to Reddit. I am hosting my own instance [here](https://teddit.nunosempere.com/); people reading this are welcome to use it. I consider the UX to be much superior to the normal reddit interface.

![](https://i.imgur.com/VVxYBM8.png)

### rose / rosenrot

[rose](https://github.com/mini-rose/rose) is a minimal browser based on webkit2 and gtk. [rosenrot](https://github.com/NunoSempere/rosenrot-browser) is my version of it, which has accumulated some cruft I like. I find it soothing to know what many parts of the code are doing, and to be able to add my own custom functionality.

![](https://i.imgur.com/VVxYBM8.png)
<figcaption>Teddit, as displayed in the rose browser</figcaption>

### prt

[prt](https://github.com/NunoSempere/PredictResolveTally), meaning "predict, resolve and tally" is a minimal bash utility I wrote to keep track of predictions.

### jq

[jq](https://stedolan.github.io/jq/) is a command-line json processor. Here is it in a command line utility used to get comments from the EA forum which include the string " bet ":

```
data='{ "query": "{ comments( input: {terms: {limit: 5000}} ){ results{ postId, htmlBody, postedAt } } }" }'

curl 'https://forum.effectivealtruism.org/graphql/' \
  -X POST  \
  -H 'content-type: application/json' \
  --data "$data"                                      |
jq .data.comments.results                             |
jq  'map(select(.htmlBody | contains(" bet ")))'      |
jq 'reverse'                                          |  
ack --passthru ' bet '

```

### Honorable mentions

Some software which I appreciate, but which I don't find soothing

- this command-line [spaced repetition program](https://github.com/MorganJamesSmith/TerminalFlashcards).
- git
- [deno](https://deno.com/)
- gitea
- f-droid
- the EA Forum
- [balena etcher](https://github.com/balena-io/etcher)
- [kiss launcher](https://kisslauncher.com/) for android


### In comparison with suckless programs 

A related concept is the [suckless](https://suckless.org/) philosophy, which aims for a fairly radical focus on simplicity and elegance. I use, enjoy and admire some of their projects. Though for now I find their setup costs a bit too costly, and their projects lacking a bit in quality of life features.

Anyways, that's all I have for today.

<p>
  <section id='isso-thread'>
  <noscript>Javascript needs to be activated to view comments.</noscript>
  </section>
</p>

