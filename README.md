# Old Gamers CLI tool

It's a simple tool that you can generate codes. For example, I'm using it to generate CRUD APIs!

## How it works?

og will load your template and payload files, then executes the template generator engine with your payload! So you need
to provide your template and payload files. It uses [text/template](https://pkg.go.dev/text/template) package
and [JSON](https://en.wikipedia.org/wiki/JSON) for the payload file.

For example, create a file named 'user.tmpl' with the following content.

```text
{{.user.name}}
```

Then, create a file named 'user.json' with the following content.

```json
{
  "user": {
    "name": "Bob"
  }
}
```

And finally, execute the following command on your terminal so you will see the final generation of the template and the
payload.

```bash
og user.json user.tmpl
```

If you want to save the output on a files, you can use the following command.

```bash
og user.json user.tmpl > user.txt
```

## Contribution

Feel free to send Pull Requests, fix typos, grammatical mistakes. We appreciate your help!
