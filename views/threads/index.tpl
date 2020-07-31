<section class="Threads">
  {{if .sessAccountName}}
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
  {{else}}
    <h1>Resister</h1>
    <div class="black-block">
      <a href="/signup/new" class="button">Sign up</a>
      <a href="/login/new" class="button">Log in</a>
    </div>
  {{end}}
  <h1>Threads</h1>
  {{range $key, $thread := .threads}}
    <section class="Thread">
      <div class="black-block">
        <div class="block-flex">
          <h1><a href="/{{$thread.HostAccount.Name}}/{{$thread.ID}}" class="text-link">{{$thread.Title}}</a></h1>
          <div class="block-right">
            <a href="/{{$thread.HostAccount.Name}}" class="text-link">by {{$thread.HostAccount.Name}}</a>
          </div>
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