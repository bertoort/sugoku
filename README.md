# suGOku

![](https://pbs.twimg.com/media/Bo7fvKJIMAA0moL.png)

[https://sugoku.herokuapp.com/](https://sugoku.herokuapp.com/)

[Go Challenge 8](http://golang-challenge.com/go-challenge8/) solution for the month of November 2015.
## Overview

Sudoku web app: solves, generates, grades, and validates sudoku puzzles.

The algorithm implements two solving functions:

  1. QuickFill - Called so because it quickly checks horizontally, vertically and in the nine grid box for possible options. If only one possible solution remains, it adds the value to the square.

  2. Guess - It fills an empty square with a possible, non-conflicting value. If the puzzle is solved (completely filled), it validates to see if it correct. If not, it goes back and fills the square with another possible value. It recursively fills in and backtracks until the puzzle is complete.

A very interesting finding was that implementing the QuickFill function before each guess did not improve the speed of the algorithm, in fact, it slowed it down. I was not expecting it. I was so surprised that I decided to keep the original name I had given the function without Quickfill: SlowSolve().

SlowSolve is the faster function and therefore the one used for the API.

## Installation

`go get github.com/bertoort/sugoku`

## Technologies

* [Go](https://golang.org)
* [Gin](https://github.com/gin-gonic/gin)
* [Semantic UI](http://semantic-ui.com/)
* [jQuery](http://jquery.com/)
* [Heroku](https://dashboard.heroku.com/)
* [Godep](https://github.com/tools/godep)

## API

### Get

Board - returns a puzzle board

`https://sugoku.herokuapp.com/board`

Arguments -

  - Difficulty:
    * easy
    * medium
    * hard
    * random

Example:

    https://sugoku.herokuapp.com/board?difficulty=easy

### Post

Solve - returns the solved puzzle, along with difficulty and status

`https://sugoku.herokuapp.com/solve`

Grade - returns the difficulty of the puzzle

`https://sugoku.herokuapp.com/grade`

Validate - returns the status of the puzzle

`https://sugoku.herokuapp.com/validate`

##### jQuery Example:

```
var data = {
  board: "[[0,0,1,0,0,0,0,0,0],
          [2,0,0,0,0,0,0,7,0],
          [0,7,0,0,0,0,0,0,0],
          [1,0,0,4,0,6,0,0,7],
          [0,0,0,0,0,0,0,0,0],
          [0,0,0,0,1,2,5,4,6],
          [3,0,2,7,6,0,9,8,0],
          [0,6,4,9,0,3,0,0,1],
          [9,8,0,5,2,1,0,6,0]]"
}
$.post('https://sugoku.herokuapp.com/solve', data)
  .done(function (response) {

    <% response = {
      difficulty: "hard",
      solution: Array[9],
      status: "solved"
    } &>    

  });```

For more, check out the [api.js](https://github.com/BertoOrt/sugoku/blob/master/public/api.js) file
