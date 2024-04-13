# suGOku

![](https://pbs.twimg.com/media/Bo7fvKJIMAA0moL.png)

[https://sugoku.onrender.com/](https://sugoku.onrender.com/)

[Go Challenge 8](http://golang-challenge.com/go-challenge8/) solution for the month of November 2015.

*The API is hosted on a [Render](https://render.com/) free tier instance. The server gets spun down with inactivity, expect a delay of 50 secs or more for the initial request. Then it's ⚡

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

- [Go](https://golang.org)
- [Gin](https://github.com/gin-gonic/gin)
- [Semantic UI](http://semantic-ui.com/)
- [jQuery](http://jquery.com/)
- [Render](https://render.com/)
- [Godep](https://github.com/tools/godep)

## API

### Get

Board - returns a puzzle board

`https://sugoku.onrender.com/board`

Arguments -

- Difficulty:
  - easy
  - medium
  - hard
  - random

Example:

    https://sugoku.onrender.com/board?difficulty=easy

### Post

Solve - returns the solved puzzle, along with difficulty and status

`https://sugoku.onrender.com/solve`

Grade - returns the difficulty of the puzzle

`https://sugoku.onrender.com/grade`

Validate - returns the status of the puzzle

`https://sugoku.onrender.com/validate`

### NOTE:

The request does not support content-type of `application/json`. It must be `application/x-www-form-urlencoded`

To help, here is a quick and dirty encoding functions for a board:

```
const encodeBoard = (board) => board.reduce((result, row, i) => result + `%5B${encodeURIComponent(row)}%5D${i === board.length -1 ? '' : '%2C'}`, '')

const encodeParams = (params) =>
  Object.keys(params)
  .map(key => key + '=' + `%5B${encodeBoard(params[key])}%5D`)
  .join('&');
```

Here is an example sending a board:

```
const data = {board:[[0,0,0,0,0,0,8,0,0],[0,0,4,0,0,8,0,0,9],[0,7,0,0,0,0,0,0,5],[0,1,0,0,7,5,0,0,8],[0,5,6,0,9,1,3,0,0],[7,8,0,0,0,0,0,0,0],[0,2,0,0,0,0,0,0,0],[0,0,0,9,3,0,0,1,0],[0,0,5,7,0,0,4,0,3]]}

fetch('https://sugoku.onrender.com/solve', {
  method: 'POST',
  body: encodeParams(data),
  headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
})
  .then(response => response.json())
  .then(response => console.log(response.solution))
  .catch(console.warn)
```

##### jQuery Example:

````
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
$.post('https://sugoku.onrender.com/solve', data)
  .done(function (response) {

    <% response = {
      difficulty: "hard",
      solution: Array[9],
      status: "solved"
    } &>

  });```

For more, check out the [api.js](https://github.com/bertoort/sugoku/blob/master/public/api.js) file
````
