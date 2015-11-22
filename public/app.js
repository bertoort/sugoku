$(function () {
  setup()
  random()
  $('.difficulty').click(function () {
    reset()
    difficulty[$(this).text().toLowerCase()]()
  })
  $('.clear').click(function () {
    reset()
  })
  $('.solve').click(function () {
    solve()
  })
  $('.validate').click(function () {
    validate()
  })
  $('.grade').click(function () {
    grade()
  })
})
