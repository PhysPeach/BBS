<section class="Threads">
  {{if .sessName}}
  <h1>New Thread</h1>
  <form action="/{{.sessName}}" method="POST">
    <input type="text" class="form-underbar-input" placeholder="Title" name="Title">
    <br>
    <textarea placeholder="Description" name="Description"></textarea>
    <input type="submit" class="button" value="Create!">
  </form>
  {{else}}
    <h1>Resister</h1>
    <a href="/signup/new" class="button">Sign up</a>
    <a href="/login/new" class="button">Log in</a>
  {{end}}
  <h1>Threads</h1>
  {{range $key, $thread := .threads}}
    <a href="/{{$thread.HostAccount.Name}}/{{$thread.ID}}">{{$thread.Title}}</a>
    <p>Created by: <a href="/{{$thread.HostAccount.Name}}">{{$thread.HostAccount.Name}}<a></p>
    <p>{{$thread.Description}}</p>
    <p>Created at: {{$thread.CreatedAt}}</p>
  {{end}}
</section>