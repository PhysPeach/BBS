<section class="profile">
  <h1>Edit Profile</h1>
  <form action="/{{.accountname}}" method="POST">
    <input type="hidden" name="_method" value="PUT">
    <input type="text" class="form-underbar-input" placeholder="{{.accountname}}" name="Name">
    <br>
    <input type="submit" class="button" value="Change Profile!">
  </form>
  <br>
  <a href="/{{.accountname}}" class="button">Cancel</a>
  <br>
  <form action="/{{.accountname}}" method="POST">
    <input type="hidden" name="_method" value="DELETE">
    <input type="submit" class="button" value="Delete Account">
  </form>
</section>