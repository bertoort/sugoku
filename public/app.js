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
})
