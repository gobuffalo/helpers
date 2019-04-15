# Helpers

---
Note: This file is auto-generated. Do Not Edit
---


## [`content#New`](https://godoc.org/github.com/gobuffalo/helpers/content#New)


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


## [`debug#New`](https://godoc.org/github.com/gobuffalo/helpers/debug#New)


## [`debug#Debug`](https://godoc.org/github.com/gobuffalo/helpers/debug#Debug)
<p>Debug by verbosely printing out using &#39;pre&#39; tags.</p>


## [`debug#Inspect`](https://godoc.org/github.com/gobuffalo/helpers/debug#Inspect)


## [`encoders#New`](https://godoc.org/github.com/gobuffalo/helpers/encoders#New)


## [`encoders#ToJSON`](https://godoc.org/github.com/gobuffalo/helpers/encoders#ToJSON)


## [`encoders#Raw`](https://godoc.org/github.com/gobuffalo/helpers/encoders#Raw)


## [`env#New`](https://godoc.org/github.com/gobuffalo/helpers/env#New)


## [`escapes#New`](https://godoc.org/github.com/gobuffalo/helpers/escapes#New)


## [`escapes#HTMLEscape`](https://godoc.org/github.com/gobuffalo/helpers/escapes#HTMLEscape)


## [`bootstrap#New`](https://godoc.org/github.com/gobuffalo/helpers/forms/bootstrap#New)


## [`bootstrap#BootstrapForm`](https://godoc.org/github.com/gobuffalo/helpers/forms/bootstrap#BootstrapForm)
<p>BootstrapForm implements a Plush helper around the
bootstrap.New function in the github.com/gobuffalo/tags/form/bootstrap package</p>


## [`bootstrap#BootstrapFormFor`](https://godoc.org/github.com/gobuffalo/helpers/forms/bootstrap#BootstrapFormFor)
<p>BootstrapFormFor implements a Plush helper around the
bootstrap.NewFormFor function in the github.com/gobuffalo/tags/form/bootstrap package</p>


## [`forms#Form`](https://godoc.org/github.com/gobuffalo/helpers/forms#Form)
<p>Form implements a Plush helper around the
form.New function in the github.com/gobuffalo/tags/form package</p>


## [`forms#FormFor`](https://godoc.org/github.com/gobuffalo/helpers/forms#FormFor)
<p>FormForimplements a Plush helper around the
form.NewFormFor function in the github.com/gobuffalo/tags/form package</p>


## [`forms#New`](https://godoc.org/github.com/gobuffalo/helpers/forms#New)


## [`hctx#Merge`](https://godoc.org/github.com/gobuffalo/helpers/hctx#Merge)


## [`inflections#New`](https://godoc.org/github.com/gobuffalo/helpers/inflections#New)


## [`iterators#Between`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Between)


## [`iterators#GroupBy`](https://godoc.org/github.com/gobuffalo/helpers/iterators#GroupBy)


## [`iterators#Next`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Next)


## [`iterators#New`](https://godoc.org/github.com/gobuffalo/helpers/iterators#New)


## [`iterators#Range`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Range)


## [`iterators#Next`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Next)


## [`iterators#Until`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Until)


## [`meta#New`](https://godoc.org/github.com/gobuffalo/helpers/meta#New)


## [`meta#Len`](https://godoc.org/github.com/gobuffalo/helpers/meta#Len)


## [`paths#Asset`](https://godoc.org/github.com/gobuffalo/helpers/paths#Asset)


## [`paths#New`](https://godoc.org/github.com/gobuffalo/helpers/paths#New)


## [`tags#CSS`](https://godoc.org/github.com/gobuffalo/helpers/tags#CSS)


## [`tags#Img`](https://godoc.org/github.com/gobuffalo/helpers/tags#Img)


## [`tags#JS`](https://godoc.org/github.com/gobuffalo/helpers/tags#JS)


## [`tags#New`](https://godoc.org/github.com/gobuffalo/helpers/tags#New)


## [`text#Markdown`](https://godoc.org/github.com/gobuffalo/helpers/text#Markdown)
<p>Markdown converts the string into HTML using GitHub flavored markdown.</p>


## [`text#New`](https://godoc.org/github.com/gobuffalo/helpers/text#New)


## [`text#Truncate`](https://godoc.org/github.com/gobuffalo/helpers/text#Truncate)


