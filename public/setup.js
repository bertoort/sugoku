function setup() {
  for (var i = 0; i < 9; i++) {
    for (var j = 0; j < 9; j++) {
      var $input = $('<input type="number" class="row'+ i+ ' col'+ j +'">')
      $('sudoku-board').append($input)
    }
  }
  $('input').keydown(function (e) {
    if ($(this).val().length > 0 && e.which != 8) {
      e.preventDefault()
    } else if (e.which < 41 && e.which > 36) {
      e.preventDefault()
    } else if (e.which == 69 || e.which == 190) {
      e.preventDefault()
    } else if (e.which == 9) {
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

function fill(arr) {
  arr.forEach(function (row, i) {
    row.forEach(function (s, j) {
      if (s !== 0) {
        $('.row'+i+'.col'+j).val(s)
      }
    })
  })
}

function setBoard(board, stat, dif) {
  $('.diff').text(dif)
  $('.status').text(stat)
  fill(board)
}
