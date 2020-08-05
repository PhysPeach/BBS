<section class="profile">
  <h1>Profile</h1>
  <div class="black-block">
    <h1>{{.account.Name}}</h1>
    <div class="block-flex">
      {{if .editable}}
        <a href="/{{.account.Name}}/edit" class="button">Edit</a>
      {{end}}
      <div class="block-time">
        Created at {{.account.CreatedAt}}
      </div>
    </div>
  </div>
  {{if .editable}}
    <h1>New Thread</h1>
    <div class="black-block">
      <form action="/{{.sessAccountName}}" method="POST">
        <input type="text" class="form-underbar-input" placeholder="Title" name="Title">
        <br>
        <textarea class="form-textarea" placeholder="Description" name="Description"></textarea>
        <br>
        <input type="submit" class="form-submit" value="Create!">
      </form>
    </div>
  {{end}}
  <h1>Threads by {{.account.Name}}</h1>
  {{range $key, $thread := .account.Threads}}
    <section class="Thread">
      <div class="black-block">
        <div class="block-flex">
          <h1><a href="/{{$thread.HostAccount.Name}}/{{$thread.ID}}" class="text-link">{{$thread.Title}}</a></h1>
        </div>
        <a href="/{{$thread.HostAccount.Name}}/{{$thread.ID}}" class="text-link">
          <div class="white-block">
            <p>{{$thread.Description}}</p>
          </div>
        </a>
        <div class="block-flex">
          <div class="block-time">
            at {{$thread.CreatedAt}}
          </div>
        </div>
      </div>
    </section>
  {{end}}
  </section>