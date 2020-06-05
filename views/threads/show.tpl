<section>
  <section>
    <h1>{{.thread.Title}}</h1>
    <p>{{.thread.Description}}</p>
    {{if .editable}}
      <form action="/{{.thread.HostAccount.Name}}/{{.thread.ID}}" method="POST">
        <input type="hidden" name="_method" value="DELETE">
        <input type="submit" class="button" value="Delete Thread">
      </form>
    {{end}}
  </section>
  <h1>Comments</h1>
  </section>