function random() {
  $.get('/board?difficulty=random')
    .done(function (data) {
      setBoard(data.board,"unsolved")
      grade(data.board)
    })
};

function hard() {
  $.get('/board?difficulty=hard')
    .done(function (data) {
      setBoard(data.board,"unsolved","hard")
    })
};

function medium() {
  $.get('/board?difficulty=medium')
    .done(function (data) {
      setBoard(data.board,"unsolved","medium")
    })
};

function easy() {
  $.get('/board?difficulty=easy')
    .done(function (data) {
      setBoard(data.board,"unsolved","easy")
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
  $.post('/solve', {board: JSON.stringify(board)})
    .done(function (data) {
      setBoard(data.solution,data.status,data.difficulty)
    })
}

function grade(board) {
  board = board || getBoard()
  $.post('/grade', {board: JSON.stringify(board)})
    .done(function (data) {
      $('.diff').text(data.difficulty)
    })
}

function validate(board) {
  board = board || getBoard()
  $.post('/validate', {board: JSON.stringify(board)})
    .done(function (data) {
      $('.status').text(data.status)
    })
}
