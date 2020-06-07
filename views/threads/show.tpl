<section class="Thread Description">
  <h1>{{.thread.Title}}</h1>
  <p>{{.thread.Description}}</p>
  {{if .editable}}
      <form action="/{{.thread.HostAccount.Name}}/{{.thread.ID}}" method="POST">
      <input type="hidden" name="_method" value="DELETE">
      <input type="submit" class="button" value="Delete Thread">
    </form>
  {{end}}
  <h1>Comments</h1>
  {{if .sessName}}
    <h1>New Comment</h1>
    <div class="black-block">
      <form action="/{{.thread.HostAccount.Name}}/{{.thread.ID}}" method="POST">
        <br>
        <textarea class="form-textarea" placeholder="Comment" name="Comment"></textarea>
        <br>
        <input type="submit" class="form-submit" value="Submit!">
      </form>
    </div>
  {{end}}
</section>