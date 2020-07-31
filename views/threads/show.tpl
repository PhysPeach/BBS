<section class="Thread Description">
  <h1></h1>
  <div class="black-block">
    <h1>{{.thread.Title}}</h1>
    <div class="white-block">
      <p>{{.thread.Description}}</p>
    </div>
    {{if .editable}}
      <form action="/{{.thread.HostAccount.Name}}/{{.thread.ID}}" method="POST">
        <input type="hidden" name="_method" value="DELETE">
        <input type="submit" class="form-submit" value="Delete Thread">
      </form>
    {{end}}
    <div class="block-flex">
      <div class="block-time">
        at {{.thread.CreatedAt}}
      </div>
    </div>
  </div>
  <div class="main">
    {{range $key, $comment := .comments}}
      <section class="Comment">
      <div class="black-block">
        <h1><a href="/{{$comment.HostAccount.Name}}" class="text-link">by {{$comment.HostAccount.Name}}</a></h1>
        <p>{{$comment.Content}}</p>
        <div class="block-flex">
          <div class="block-time">
            at {{$comment.CreatedAt}}
          </div>
        </div>
      </div>
    </section>
  {{end}}
    {{if .sessAccountName}}
      <h1>New Comment</h1>
      <div class="black-block">
        <form action="/{{.thread.HostAccount.Name}}/{{.thread.ID}}" method="POST">
          <br>
          <textarea class="form-textarea" placeholder="Comment" name="Content"></textarea>
          <br>
          <input type="submit" class="form-submit" value="Submit!">
        </form>
      </div>
    {{end}}
  </div>
</section>