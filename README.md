# Helpers

[![Build Status](https://dev.azure.com/markbates/buffalo/_apis/build/status/gobuffalo.helpers?branchName=master)](https://dev.azure.com/markbates/buffalo/_build/latest?definitionId=49&branchName=master)[![GoDoc](https://godoc.org/github.com/gobuffalo/helpers?status.svg)](https://godoc.org/github.com/gobuffalo/helpers)
---
Note: This file is auto-generated. Do Not Edit
---


## [`bootstrap#Form`](https://godoc.org/github.com/gobuffalo/helpers/forms/bootstrap#Form)
<p>Form implements a Plush helper around the
bootstrap.New function in the github.com/gobuffalo/tags/form/bootstrap package</p>


## [`bootstrap#FormFor`](https://godoc.org/github.com/gobuffalo/helpers/forms/bootstrap#FormFor)
<p>FormFor implements a Plush helper around the
bootstrap.NewFormFor function in the github.com/gobuffalo/tags/form/bootstrap package</p>


## [`bootstrap#New`](https://godoc.org/github.com/gobuffalo/helpers/forms/bootstrap#New)


## [`content#ContentFor`](https://godoc.org/github.com/gobuffalo/helpers/content#ContentFor)
<p>ContentFor stores a block of templating code to be re-used later in the template
via the contentOf helper.
An optional map of values can be passed to contentOf,
which are made available to the contentFor block.</p>

<pre><code>&lt;% contentFor(&#34;buttons&#34;) { %&gt;
    &lt;button&gt;hi&lt;/button&gt;
&lt;% } %&gt;
</code></pre>


## [`content#ContentOf`](https://godoc.org/github.com/gobuffalo/helpers/content#ContentOf)
<p>ContentOf retrieves a stored block for templating and renders it.
You can pass an optional map of fields that will be set.</p>

<pre><code>&lt;%= contentOf(&#34;buttons&#34;) %&gt;
&lt;%= contentOf(&#34;buttons&#34;, {&#34;label&#34;: &#34;Click me&#34;}) %&gt;
</code></pre>


## [`content#New`](https://godoc.org/github.com/gobuffalo/helpers/content#New)


## [`debug#Debug`](https://godoc.org/github.com/gobuffalo/helpers/debug#Debug)
<p>Debug by verbosely printing out using &#39;pre&#39; tags.</p>


## [`debug#Inspect`](https://godoc.org/github.com/gobuffalo/helpers/debug#Inspect)


## [`debug#New`](https://godoc.org/github.com/gobuffalo/helpers/debug#New)


## [`encoders#New`](https://godoc.org/github.com/gobuffalo/helpers/encoders#New)


## [`encoders#Raw`](https://godoc.org/github.com/gobuffalo/helpers/encoders#Raw)


## [`encoders#ToJSON`](https://godoc.org/github.com/gobuffalo/helpers/encoders#ToJSON)


## [`env#New`](https://godoc.org/github.com/gobuffalo/helpers/env#New)


## [`escapes#HTMLEscape`](https://godoc.org/github.com/gobuffalo/helpers/escapes#HTMLEscape)


## [`escapes#New`](https://godoc.org/github.com/gobuffalo/helpers/escapes#New)


## [`forms#Form`](https://godoc.org/github.com/gobuffalo/helpers/forms#Form)
<p>Form implements a Plush helper around the
form.New function in the github.com/gobuffalo/tags/form package</p>


## [`forms#FormFor`](https://godoc.org/github.com/gobuffalo/helpers/forms#FormFor)
<p>FormForimplements a Plush helper around the
form.NewFormFor function in the github.com/gobuffalo/tags/form package</p>


## [`forms#New`](https://godoc.org/github.com/gobuffalo/helpers/forms#New)


## [`inflections#New`](https://godoc.org/github.com/gobuffalo/helpers/inflections#New)


## [`iterators#Between`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Between)


## [`iterators#GroupBy`](https://godoc.org/github.com/gobuffalo/helpers/iterators#GroupBy)


## [`iterators#New`](https://godoc.org/github.com/gobuffalo/helpers/iterators#New)


## [`iterators#Next`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Next)


## [`iterators#Next`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Next)


## [`iterators#Range`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Range)


## [`iterators#Until`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Until)


## [`meta#Len`](https://godoc.org/github.com/gobuffalo/helpers/meta#Len)


## [`meta#New`](https://godoc.org/github.com/gobuffalo/helpers/meta#New)


## [`paths#New`](https://godoc.org/github.com/gobuffalo/helpers/paths#New)


## [`paths#PathFor`](https://godoc.org/github.com/gobuffalo/helpers/paths#PathFor)
<p>PathFor takes an <code>interface{}</code>, or a <code>slice</code> of them,
and tries to convert it to a <code>/foos/{id}</code> style URL path.
Rules:</p>

<ul>
<li>if <code>string</code> it is returned as is</li>
<li>if <code>Pathable</code> the <code>ToPath</code> method is returned</li>
<li>if <code>slice</code> or an <code>array</code> each element is run through the helper then joined</li>
<li>if <code>struct</code> the name of the struct, pluralized is used for the name</li>
<li>if <code>Paramable</code> the <code>ToParam</code> method is used to fill the <code>{id}</code> slot</li>
<li>if <code>struct.Slug</code> the slug is used to fill the <code>{id}</code> slot of the URL</li>
<li>if <code>struct.ID</code> the ID is used to fill the <code>{id}</code> slot of the URL</li>
</ul>


## [`tags#CSS`](https://godoc.org/github.com/gobuffalo/helpers/tags#CSS)


## [`tags#Img`](https://godoc.org/github.com/gobuffalo/helpers/tags#Img)


## [`tags#JS`](https://godoc.org/github.com/gobuffalo/helpers/tags#JS)


## [`tags#LinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#LinkTo)


## [`tags#New`](https://godoc.org/github.com/gobuffalo/helpers/tags#New)


## [`text#Markdown`](https://godoc.org/github.com/gobuffalo/helpers/text#Markdown)
<p>Markdown converts the string into HTML using GitHub flavored markdown.</p>


## [`text#New`](https://godoc.org/github.com/gobuffalo/helpers/text#New)


## [`text#Truncate`](https://godoc.org/github.com/gobuffalo/helpers/text#Truncate)


