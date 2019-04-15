# Helpers

---
Note: This file is auto-generated. Do Not Edit
---


# [`github.com/gobuffalo/helpers/content#New`](https://godoc.org/github.com/gobuffalo/helpers/content#New)


# [`github.com/gobuffalo/helpers/content#ContentFor`](https://godoc.org/github.com/gobuffalo/helpers/content#ContentFor)
<p>ContentFor stores a block of templating code to be re-used later in the template
via the contentOf helper.
An optional map of values can be passed to contentOf,
which are made available to the contentFor block.</p>

<pre><code>&lt;% contentFor(&#34;buttons&#34;) { %&gt;
    &lt;button&gt;hi&lt;/button&gt;
&lt;% } %&gt;
</code></pre>


# [`github.com/gobuffalo/helpers/content#ContentOf`](https://godoc.org/github.com/gobuffalo/helpers/content#ContentOf)
<p>ContentOf retrieves a stored block for templating and renders it.
You can pass an optional map of fields that will be set.</p>

<pre><code>&lt;%= contentOf(&#34;buttons&#34;) %&gt;
&lt;%= contentOf(&#34;buttons&#34;, {&#34;label&#34;: &#34;Click me&#34;}) %&gt;
</code></pre>


# [`github.com/gobuffalo/helpers/debug#New`](https://godoc.org/github.com/gobuffalo/helpers/debug#New)


# [`github.com/gobuffalo/helpers/debug#Debug`](https://godoc.org/github.com/gobuffalo/helpers/debug#Debug)
<p>Debug by verbosely printing out using &#39;pre&#39; tags.</p>


# [`github.com/gobuffalo/helpers/debug#Inspect`](https://godoc.org/github.com/gobuffalo/helpers/debug#Inspect)


# [`github.com/gobuffalo/helpers/encoders#New`](https://godoc.org/github.com/gobuffalo/helpers/encoders#New)


# [`github.com/gobuffalo/helpers/encoders#ToJSON`](https://godoc.org/github.com/gobuffalo/helpers/encoders#ToJSON)


# [`github.com/gobuffalo/helpers/encoders#Raw`](https://godoc.org/github.com/gobuffalo/helpers/encoders#Raw)


# [`github.com/gobuffalo/helpers/env#New`](https://godoc.org/github.com/gobuffalo/helpers/env#New)


# [`github.com/gobuffalo/helpers/escapes#New`](https://godoc.org/github.com/gobuffalo/helpers/escapes#New)


# [`github.com/gobuffalo/helpers/escapes#HTMLEscape`](https://godoc.org/github.com/gobuffalo/helpers/escapes#HTMLEscape)


# [`github.com/gobuffalo/helpers/bootstrap#New`](https://godoc.org/github.com/gobuffalo/helpers/bootstrap#New)


# [`github.com/gobuffalo/helpers/bootstrap#BootstrapForm`](https://godoc.org/github.com/gobuffalo/helpers/bootstrap#BootstrapForm)
<p>BootstrapForm implements a Plush helper around the
bootstrap.New function in the github.com/gobuffalo/tags/form/bootstrap package</p>


# [`github.com/gobuffalo/helpers/bootstrap#BootstrapFormFor`](https://godoc.org/github.com/gobuffalo/helpers/bootstrap#BootstrapFormFor)
<p>BootstrapFormFor implements a Plush helper around the
bootstrap.NewFormFor function in the github.com/gobuffalo/tags/form/bootstrap package</p>


# [`github.com/gobuffalo/helpers/forms#Form`](https://godoc.org/github.com/gobuffalo/helpers/forms#Form)
<p>Form implements a Plush helper around the
form.New function in the github.com/gobuffalo/tags/form package</p>


# [`github.com/gobuffalo/helpers/forms#FormFor`](https://godoc.org/github.com/gobuffalo/helpers/forms#FormFor)
<p>FormForimplements a Plush helper around the
form.NewFormFor function in the github.com/gobuffalo/tags/form package</p>


# [`github.com/gobuffalo/helpers/forms#New`](https://godoc.org/github.com/gobuffalo/helpers/forms#New)


# [`github.com/gobuffalo/helpers/docs#New`](https://godoc.org/github.com/gobuffalo/helpers/docs#New)


# [`github.com/gobuffalo/helpers/docs#Validate`](https://godoc.org/github.com/gobuffalo/helpers/docs#Validate)
<p>Validate that options are usuable</p>


# [`github.com/gobuffalo/helpers/hctx#Merge`](https://godoc.org/github.com/gobuffalo/helpers/hctx#Merge)


# [`github.com/gobuffalo/helpers/cmd#Execute`](https://godoc.org/github.com/gobuffalo/helpers/cmd#Execute)
<p>Execute adds all child commands to the root command and sets flags appropriately.
This is called by main.main(). It only needs to happen once to the rootCmd.</p>


# [`github.com/gobuffalo/helpers/helptest#NewContext`](https://godoc.org/github.com/gobuffalo/helpers/helptest#NewContext)


# [`github.com/gobuffalo/helpers/helptest#New`](https://godoc.org/github.com/gobuffalo/helpers/helptest#New)


# [`github.com/gobuffalo/helpers/helptest#Value`](https://godoc.org/github.com/gobuffalo/helpers/helptest#Value)


# [`github.com/gobuffalo/helpers/helptest#Set`](https://godoc.org/github.com/gobuffalo/helpers/helptest#Set)


# [`github.com/gobuffalo/helpers/helptest#Block`](https://godoc.org/github.com/gobuffalo/helpers/helptest#Block)


# [`github.com/gobuffalo/helpers/helptest#BlockWith`](https://godoc.org/github.com/gobuffalo/helpers/helptest#BlockWith)


# [`github.com/gobuffalo/helpers/helptest#HasBlock`](https://godoc.org/github.com/gobuffalo/helpers/helptest#HasBlock)


# [`github.com/gobuffalo/helpers/helptest#Render`](https://godoc.org/github.com/gobuffalo/helpers/helptest#Render)


# [`github.com/gobuffalo/helpers/inflections#New`](https://godoc.org/github.com/gobuffalo/helpers/inflections#New)


# [`github.com/gobuffalo/helpers/iterators#Between`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Between)


# [`github.com/gobuffalo/helpers/iterators#GroupBy`](https://godoc.org/github.com/gobuffalo/helpers/iterators#GroupBy)


# [`github.com/gobuffalo/helpers/iterators#Next`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Next)


# [`github.com/gobuffalo/helpers/iterators#New`](https://godoc.org/github.com/gobuffalo/helpers/iterators#New)


# [`github.com/gobuffalo/helpers/iterators#Range`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Range)


# [`github.com/gobuffalo/helpers/iterators#Next`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Next)


# [`github.com/gobuffalo/helpers/iterators#Until`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Until)


# [`github.com/gobuffalo/helpers/meta#New`](https://godoc.org/github.com/gobuffalo/helpers/meta#New)


# [`github.com/gobuffalo/helpers/meta#Len`](https://godoc.org/github.com/gobuffalo/helpers/meta#Len)


# [`github.com/gobuffalo/helpers/paths#Asset`](https://godoc.org/github.com/gobuffalo/helpers/paths#Asset)


# [`github.com/gobuffalo/helpers/paths#New`](https://godoc.org/github.com/gobuffalo/helpers/paths#New)


# [`github.com/gobuffalo/helpers/tags#CSS`](https://godoc.org/github.com/gobuffalo/helpers/tags#CSS)


# [`github.com/gobuffalo/helpers/tags#Img`](https://godoc.org/github.com/gobuffalo/helpers/tags#Img)


# [`github.com/gobuffalo/helpers/tags#JS`](https://godoc.org/github.com/gobuffalo/helpers/tags#JS)


# [`github.com/gobuffalo/helpers/tags#New`](https://godoc.org/github.com/gobuffalo/helpers/tags#New)


# [`github.com/gobuffalo/helpers/text#Markdown`](https://godoc.org/github.com/gobuffalo/helpers/text#Markdown)
<p>Markdown converts the string into HTML using GitHub flavored markdown.</p>


# [`github.com/gobuffalo/helpers/text#New`](https://godoc.org/github.com/gobuffalo/helpers/text#New)


# [`github.com/gobuffalo/helpers/text#Truncate`](https://godoc.org/github.com/gobuffalo/helpers/text#Truncate)


