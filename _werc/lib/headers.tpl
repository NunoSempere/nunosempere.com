<!DOCTYPE HTML>
<html>
<head>

    <title>%($pageTitle%)</title>

    <link rel="stylesheet" href="/pub/style/style.css" type="text/css" media="screen, handheld" title="default">
    <link rel="shortcut icon" href="/favicon.ico" type="image/vnd.microsoft.icon">
% if(test -f $sitedir/_werc/pub/style.css)
%    echo '    <link rel="stylesheet" href="/_werc/pub/style.css" type="text/css" media="screen" title="default">'
    <link rel="alternate" type="application/rss+xml" title="RSS for Measure is Unceasing" href="/blog/index.rss" />
    <meta charset="UTF-8">
% # Legacy charset declaration for backards compatibility with non-html5 browsers.
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8"> 
    <meta property="og:image" content="https://cards.nunosempere.com/api/dynamic-image?endpoint=%($req_path%)">
    <meta name="twitter:card" content="summary_large_image" />
    <meta name="twitter:title" content="Measure is unceasing" />
    <meta name="twitter:description" content="%($pageTitle%)" />
    <meta name="twitter:url" content="https://nunosempere.com/" />
    <meta name="twitter:image" content="https://cards.nunosempere.com/api/dynamic-image?endpoint=%($req_path%)" />
    <meta name="twitter:site" content="@NunoSempere" />

% if(! ~ $#meta_description 0)
%   echo '    <meta name="description" content="'$"meta_description'">'
% if(! ~ $#meta_keywords 0)
%   echo '    <meta name="keywords" content="'$"meta_keywords'">'

% h = `{get_lib_file headers.inc}
% if(! ~ $#h 0)
%   cat $h

    %($"extraHeaders%)
    <script data-isso="//comments.nunosempere.com/" data-isso-max-comments-top="inf" data-isso-max-comments-nested="inf" data-isso-postbox-text-text-en="On the Internet, nobody knows you are a dog" src="//comments.nunosempere.com/js/embed.min.js"></script>

% # To add math
% # <script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script>
% # <script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>

</head>
<body>

