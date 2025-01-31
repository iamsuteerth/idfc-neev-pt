# Tic-Tac-Toe React App

Simple Tic-Tac-Toe game built using React. Utilizes CSS Modules for styling and React Spring for a smooth fade-in animation.

[Website URL](https://tictactoe-ss.vercel.app)

## Features

- Two players can play against each other.
- Detects the winner and displays the result (or a draw).
- Tracks the game history, allowing players to review past moves.
- Reset button to start a new game.
- Smooth fade-in animation on game load.

## Technologies Used

- React
- CSS Modules (for styling)
- React Spring (for animations)

## Project Structure
```
tictactoe/
├── src/
│   ├── components/
│   │   ├── Board.jsx          // Main game board component
│   │   ├── Board.module.css  // Styles for the Board component
│   │   ├── Square.jsx         // Individual square component
│   │   ├── Square.module.css // Styles for the Square component
│   ├── utils/
│   │   └── GameLogic.js      // Game logic (winner calculation)
│   ├── App.jsx              // Main app component
│   ├── App.module.css      // Styles for the App component
│   ├── main.jsx             // Entry point of the application
│   └── ...
└── ...
```
## How to Run

1. Clone the repository: `git clone https://github.com/iamsuteerth/idfc-neev-pt.git`
2. Navigate to the project directory: `cd tictactoe`
3. Install dependencies: `npm install` 
4. Start: `npm run dev` 

## Game Logic

- The game logic, including the `calculateWinner` function, is located in [here](/tictactoe/src/utils/GameLogic.js) 
- This function takes an array representing the board squares as input and checks all possible winning combinations. It returns the winner ('X' or 'O') if there is one, or `null` if the game is still ongoing or a draw.

## Components

- **Board:** The main game board component. It manages the rendering of the squares and handles user clicks. It uses nested `map` functions to dynamically render the 3x3 board.  It receives the `squares` array (representing the current board state) and the `onClick` function (to handle clicks) as props from the `App` component. The component can be found [here](/tictactoe/src/components/Board.jsx)

* **Square:** Represents a single square on the board. It displays the value (X or O) and triggers the `onClick` event.  This component uses React Spring to animate the X or O when it appears on the square, creating a "wobbly" effect. It receives the `value` (X or O) and the `onClick` function as props. The component can be found [here](/tictactoe/src/components/Square.jsx)

## App Component (`App.jsx`)

The `App` component manages the overall game state, including the move history, current step, and whose turn it is.  It handles clicks on the board, updates the game state, and determines the winner by calling the `calculateWinner` function. It also provides the "Jump to" functionality to review previous moves and the "Reset Game" button.  The [App](/tictactoe/src/App.jsx) component uses React Spring to animate the game board on load. 

## Styling

CSS Modules are used for styling to prevent style conflicts. Each component has its own CSS module file.

## Animation

React Spring is used to create a smooth fade-in animation when the game board loads and a "wobbly" animation for the X's and O's when they appear on the squares.