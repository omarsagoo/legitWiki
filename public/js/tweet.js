$(document).ready(() => {
    $("#tweetModal").on("show.bs.modal", event => {
        document.getElementById("container").innerHTML = `<div class="text-center">
        <div class="spinner-border" role="status">
          <span class="sr-only">Loading...</span>
        </div>
      </div>`
        var button = $(event.relatedTarget) // Button that triggered the modal
        var recipient = button.data('articleId') // Extract info from data-* attributes
        console.log(recipient)

        $.get("/tweet/" + recipient, tweet => {
            document.getElementById("container").innerHTML = "";
            twttr.widgets.createTweet(
                tweet,
                document.getElementById('container'),
                {
                  theme: 'dark'
                }
              );
        })
        var modal = $(this)
        modal.find('.modal-title').text('New message to ' + recipient)
    })
})