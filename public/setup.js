function setup() {
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
}

function reset() {
  for (var i = 0; i < 9; i++) {
    for (var j = 0; j < 9; j++) {
      $('.row'+i+'.col'+j).val("")
    }
  }
}

function getBoard() {
  var board = []
  for (var i = 0; i < 9; i++) {
    var row = []
    for (var j = 0; j < 9; j++) {
      var n = $('.row'+i+'.col'+j).val()
      if (n === "") n = "0"
      row.push(parseInt(n))
    }
    board.push(row)
  }
  return board
}
