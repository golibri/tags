[![Built with Spacemacs](https://cdn.rawgit.com/syl20bnr/spacemacs/442d025779da2f62fc86c2082703697714db6514/assets/spacemacs-badge.svg)](http://github.com/syl20bnr/spacemacs)

# golibri/tags
Calculate relevant tags (aka keywords) from text string

**work in progress**

# installation
`go get github.com/dchest/stemmer/german`

`go get github.com/dchest/stemmer/porter2`

`go get github.com/bbalet/stopwords`

`go get github.com/golibri/tags`

# usage
````go
result := tags.Calculate("big-text-string", "en")
tags := result.Words
stems := result.Stems
dict := result.Dictionary // lookup stem -> words
````

# license
LGPLv3. (You can use it in commercial projects as you like, but improvements/bugfixes must flow back to this lib.)
