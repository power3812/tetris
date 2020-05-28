(function($){
  $(document).ready(function(){
    $('#start-button').on('click', function() {
      var game_start = $('#game_start');
      game_start[0].volume = 0.1;
      game_start[0].play();

    });
  });
})(jQuery);
