$(function () {
  for (var i = 0; i < 9; i++) {
    for (var j = 0; j < 9; j++) {
      var $input = $('<input type="number" class="row'+ i+ ' col'+ j +'">')
      $('sudoku-board').append($input)
    }
  }
  $('input').keydown(function (e) {
    console.log(e.which);
    if ($(this).val().length > 0 && e.which != 8) {
      e.preventDefault()
    } else if (e.which < 41 && e.which > 36) {
      e.preventDefault()
    }
  })
})
