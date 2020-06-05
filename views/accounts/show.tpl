<section class="profile">
  <section class="container">
    <h1>{{.accountname}}</h1>
    {{if .editable}}
      <a href="/{{.accountname}}/edit" class="button">Edit</a>
    {{end}}
  </section>
  <h1>Threads</h1>
  {{range $key, $thread := .threads}}
    <a href="/{{$thread.HostAccount.Name}}/{{$thread.ID}}">{{$thread.Title}}</a>
    <p>{{$thread.Description}}</p>
    <p>Created at {{$thread.CreatedAt}}</p>
  {{end}}
  </section>