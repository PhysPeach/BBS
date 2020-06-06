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
  <h1>Threads by {{.account.Name}}</h1>
  {{range $key, $thread := .threads}}
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