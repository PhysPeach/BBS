<section class="log-in">
    <h1>Log in</h1>
    {{if .loginError}}
      <div class="errors">
        <ul>
          <li>{{.loginError}}</li>
        <ul>
      </div>
    {{end}}
    <div class="black-block">
      <form action="/login" method="POST">
        {{.xsrf}}
        <input type="text" class="form-underbar-input" placeholder="Name" name="Name">
        <br>
        <input type="Password" class="form-underbar-input" placeholder="Password" name="Password">
        <br>
        <input type="submit" class="form-submit" value="Log in">
      </form>
    </div>
  </section>