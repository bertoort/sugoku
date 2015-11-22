function random() {
  $.get('/board?difficulty=random')
    .done(function (data) {
      data.board.forEach(function (row, i) {
        row.forEach(function (s, j) {
          if (s !== 0) {
            $('.row'+i+'.col'+j).val(s)
          }
        })
      })
    })
};

function hard() {
  $.get('/board?difficulty=hard')
    .done(function (data) {
      data.board.forEach(function (row, i) {
        row.forEach(function (s, j) {
          if (s !== 0) {
            $('.row'+i+'.col'+j).val(s)
          }
        })
      })
    })
};

function medium() {
  $.get('/board?difficulty=medium')
    .done(function (data) {
      data.board.forEach(function (row, i) {
        row.forEach(function (s, j) {
          if (s !== 0) {
            $('.row'+i+'.col'+j).val(s)
          }
        })
      })
    })
};

function easy() {
  $.get('/board?difficulty=easy')
    .done(function (data) {
      data.board.forEach(function (row, i) {
        row.forEach(function (s, j) {
          if (s !== 0) {
            $('.row'+i+'.col'+j).val(s)
          }
        })
      })
    })
}

var difficulty = {
  "easy": easy,
  "medium": medium,
  "hard": hard,
  "random": random,
}

function solve() {
  var board = getBoard()
  console.log(JSON.stringify(board));
  $.post('/solve', {board: JSON.stringify(board)})
    .done(function (data) {
      console.log(data);
    })
}
